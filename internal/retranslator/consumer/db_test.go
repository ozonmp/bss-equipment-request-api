package consumer

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/bss-equipment-request-api/internal/retranslator/mocks"
	"github.com/ozonmp/bss-equipment-request-api/internal/retranslator/model"
	"sync"
	"testing"
	"time"
)

func setUp(t *testing.T, consumers uint64, channelSize int, timeout time.Duration, batchSize uint64) (Config, chan model.EquipmentRequestEvent, *mocks.MockEventRepo, context.CancelFunc) {
	parentCtx := context.Background()
	ctrl := gomock.NewController(t)
	repo := mocks.NewMockEventRepo(ctrl)
	ctx, ctxCancel := context.WithCancel(parentCtx)

	mockDB, _, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	events := make(chan model.EquipmentRequestEvent, channelSize)

	config := Config{
		N:         consumers,
		Events:    events,
		Repo:      repo,
		BatchSize: batchSize,
		Timeout:   timeout,
		Ctx:       ctx,
		DB:        sqlxDB,
	}

	return config, events, repo, ctxCancel
}

func tearDown(db Consumer, ctxFunc context.CancelFunc) {
	ctxFunc()
	db.Close()
}

func TestStartAndGetOneEvent(t *testing.T) {
	t.Parallel()
	config, events, repo, ctxFunc := setUp(t, 4, 1, time.Millisecond, 1)

	db := NewDbConsumer(
		config.Ctx,
		config.N,
		config.BatchSize,
		config.Timeout,
		config.Repo,
		config.Events,
		config.DB)
	defer tearDown(db, ctxFunc)

	eventCount := int(config.N)
	var wg sync.WaitGroup
	wg.Add(eventCount)
	defer wg.Wait()

	var eventList = []model.EquipmentRequestEvent{
		{
			ID:                 12,
			Type:               model.Created,
			Status:             model.Unlocked,
			CreatedAt:          time.Now(),
			EquipmentRequestID: 17,
			Payload:            &model.EquipmentRequest{},
		},
	}

	repo.EXPECT().Lock(config.Ctx, config.DB, config.BatchSize).DoAndReturn(
		func(ctx context.Context, db *sqlx.DB, batchSize uint64) ([]model.EquipmentRequestEvent, error) {
			wg.Done()
			return eventList, nil
		}).Times(eventCount)

	db.Start()
	for i := 0; i < eventCount; i++ {
		<-events
	}
}

func TestStartAndGetErrors(t *testing.T) {
	t.Parallel()
	config, _, repo, ctxFunc := setUp(t, 5, 10, time.Millisecond, 3)

	db := NewDbConsumer(
		config.Ctx,
		config.N,
		config.BatchSize,
		config.Timeout,
		config.Repo,
		config.Events,
		config.DB)
	defer tearDown(db, ctxFunc)

	eventCount := int(config.N)
	var wg sync.WaitGroup
	wg.Add(eventCount)
	defer wg.Wait()

	repo.EXPECT().Lock(config.Ctx, config.DB, config.BatchSize).DoAndReturn(
		func(ctx context.Context, db *sqlx.DB, batchSize uint64) ([]model.EquipmentRequestEvent, error) {
			wg.Done()
			return nil, errors.New("empty result error")
		}).Times(eventCount)

	db.Start()
}

func TestStartAndGetSeveralEvent(t *testing.T) {
	t.Parallel()
	config, _, repo, ctxFunc := setUp(t, 3, 10, time.Millisecond, 2)

	db := NewDbConsumer(
		config.Ctx,
		config.N,
		config.BatchSize,
		config.Timeout,
		config.Repo,
		config.Events,
		config.DB)
	defer tearDown(db, ctxFunc)

	eventCount := int(config.N)

	var wg sync.WaitGroup
	wg.Add(eventCount)
	defer wg.Wait()

	eventList := []model.EquipmentRequestEvent{
		{
			ID:                 12,
			Type:               model.Created,
			Status:             model.Unlocked,
			CreatedAt:          time.Now(),
			EquipmentRequestID: 17,
			Payload:            &model.EquipmentRequest{},
		},
		{
			ID:                 14,
			Type:               model.UpdatedStatus,
			Status:             model.Unlocked,
			CreatedAt:          time.Now(),
			EquipmentRequestID: 12,
			Payload:            &model.EquipmentRequest{},
		},
	}

	repo.EXPECT().Lock(config.Ctx, config.DB, config.BatchSize).DoAndReturn(
		func(ctx context.Context, db *sqlx.DB, batchSize uint64) ([]model.EquipmentRequestEvent, error) {
			wg.Done()
			return eventList, nil
		}).Times(1)

	repo.EXPECT().Lock(config.Ctx, config.DB, config.BatchSize).DoAndReturn(
		func(ctx context.Context, db *sqlx.DB, batchSize uint64) ([]model.EquipmentRequestEvent, error) {
			wg.Done()
			return nil, errors.New("empty result error")
		}).Times(eventCount - 1)

	db.Start()
}
