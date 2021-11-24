package retranslator

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/bss-equipment-request-api/internal/retranslator/config"
	"github.com/ozonmp/bss-equipment-request-api/internal/retranslator/model"
	"sync"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/ozonmp/bss-equipment-request-api/internal/retranslator/mocks"
)

type retranslatorSetting struct {
	ctx    context.Context
	config config.Retranslator
	repo   *mocks.MockEventRepo
	sender *mocks.MockEventSender
	db     *sqlx.DB
}

func setUp(t *testing.T) retranslatorSetting {
	ctx := context.Background()
	ctrl := gomock.NewController(t)

	repo := mocks.NewMockEventRepo(ctrl)
	sender := mocks.NewMockEventSender(ctrl)

	mockDB, _, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")

	config := config.Retranslator{
		ChannelSize:     512,
		ConsumerCount:   2,
		ConsumeSize:     10,
		ConsumeTimeout:  50 * time.Millisecond,
		ProducerCount:   2,
		ProducerTimeout: 50 * time.Millisecond,
		BatchSize:       1,
		WorkerCount:     2,
	}

	return retranslatorSetting{
		ctx:    ctx,
		config: config,
		repo:   repo,
		sender: sender,
		db:     sqlxDB,
	}
}

func tearDown(retranslator Retranslator) {
	retranslator.Close()
}

func TestStartAndGetOneEvent(t *testing.T) {
	t.Parallel()
	setting := setUp(t)
	retranslator := NewRetranslator(setting.ctx, setting.db, setting.config, setting.repo, setting.sender)
	defer tearDown(retranslator)

	event := model.EquipmentRequestEvent{
		ID:                 12,
		Type:               model.Created,
		Status:             model.Unlocked,
		CreatedAt:          time.Now(),
		EquipmentRequestID: 17,
		Payload:            &model.EquipmentRequest{},
	}

	eventLockCount := int(setting.config.ConsumerCount * setting.config.BatchSize)
	var wg sync.WaitGroup
	wg.Add(eventLockCount)
	defer wg.Wait()

	setting.repo.EXPECT().Lock(gomock.Any(), setting.db, setting.config.ConsumeSize).DoAndReturn(
		func(ctx context.Context, db *sqlx.DB, batchSize uint64) ([]model.EquipmentRequestEvent, error) {
			wg.Done()
			return []model.EquipmentRequestEvent{event}, nil
		}).Times(eventLockCount)

	eventSendCount := int(setting.config.ProducerCount * setting.config.BatchSize)
	removeCount := int(setting.config.ProducerCount * setting.config.BatchSize)

	var wgSender sync.WaitGroup
	wgSender.Add(eventSendCount)
	defer wgSender.Wait()

	var wgRepo sync.WaitGroup
	wgRepo.Add(removeCount)
	defer wgRepo.Wait()

	setting.sender.EXPECT().Send(&event).DoAndReturn(
		func(*model.EquipmentRequestEvent) error {
			wgSender.Done()
			return nil
		}).Times(eventSendCount)

	setting.repo.EXPECT().Remove(gomock.Any(), []uint64{event.ID}).DoAndReturn(
		func(context.Context, []uint64) error {
			wgRepo.Done()
			return nil
		}).Times(removeCount)

	retranslator.Start()
}
