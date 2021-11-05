package consumer

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/ozonmp/bss-equipment-request-api/internal/mocks"
	"github.com/ozonmp/bss-equipment-request-api/internal/model"
	"sync"
	"testing"
	"time"
)

func setUp(t *testing.T, consumers uint64, channelSize int, timeout time.Duration, batchSize uint64) (Config, chan model.EquipmentRequestEvent, *mocks.MockEventRepo, context.CancelFunc) {
	parentCtx := context.Background()
	ctrl := gomock.NewController(t)
	repo := mocks.NewMockEventRepo(ctrl)
	ctx, ctxCancel := context.WithCancel(parentCtx)

	events := make(chan model.EquipmentRequestEvent, channelSize)

	config := Config{
		N:         consumers,
		Events:    events,
		Repo:      repo,
		BatchSize: batchSize,
		Timeout:   timeout,
		Ctx:       ctx,
	}

	return config, events, repo, ctxCancel
}

func tearDown(db Consumer, ctxFunc context.CancelFunc) {
	ctxFunc()
	db.Close()
}

func TestStartAndGetOneEvent(t *testing.T) {
	t.Parallel()
	config, events, repo, ctxFunc := setUp(t, 5, 1, time.Millisecond, 1)
	db := NewDbConsumer(
		config.Ctx,
		config.N,
		config.BatchSize,
		config.Timeout,
		config.Repo,
		config.Events)
	defer tearDown(db, ctxFunc)

	event := model.EquipmentRequestEvent{
		ID:     12,
		Type:   model.Created,
		Status: model.Deferred,
		Entity: &model.EquipmentRequest{ID: 1, EmployeeID: 1, EquipmentID: 1, CreatedAt: nil, DoneAt: nil, EquipmentRequestStatusID: model.Done},
	}

	eventCount := int(config.N)
	var wg sync.WaitGroup
	wg.Add(eventCount)
	defer wg.Wait()

	repo.EXPECT().Lock(config.BatchSize).DoAndReturn(
		func(uint642 uint64) ([]model.EquipmentRequestEvent, error) {
			wg.Done()
			return []model.EquipmentRequestEvent{event}, nil
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
		config.Events)
	defer tearDown(db, ctxFunc)

	eventCount := int(config.N)
	var wg sync.WaitGroup
	wg.Add(eventCount)
	defer wg.Wait()

	repo.EXPECT().Lock(config.BatchSize).DoAndReturn(
		func(uint642 uint64) ([]model.EquipmentRequestEvent, error) {
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
		config.Events)
	defer tearDown(db, ctxFunc)

	events := []model.EquipmentRequestEvent{
		{
			ID:     12,
			Type:   model.Created,
			Status: model.Deferred,
			Entity: &model.EquipmentRequest{ID: 1, EmployeeID: 1, EquipmentID: 1, CreatedAt: nil, DoneAt: nil, EquipmentRequestStatusID: model.Done},
		},
		{
			ID:     14,
			Type:   model.Updated,
			Status: model.Processed,
			Entity: &model.EquipmentRequest{ID: 1, EmployeeID: 1, EquipmentID: 1, CreatedAt: nil, DoneAt: nil, EquipmentRequestStatusID: model.Done},
		},
	}

	eventCount := int(config.N)

	var wg sync.WaitGroup
	wg.Add(eventCount)
	defer wg.Wait()

	repo.EXPECT().Lock(config.BatchSize).DoAndReturn(
		func(uint642 uint64) ([]model.EquipmentRequestEvent, error) {
			wg.Done()
			return events, nil
		}).Times(1)

	repo.EXPECT().Lock(config.BatchSize).DoAndReturn(
		func(uint642 uint64) ([]model.EquipmentRequestEvent, error) {
			wg.Done()
			return nil, errors.New("empty result error")
		}).Times(eventCount - 1)

	db.Start()
}
