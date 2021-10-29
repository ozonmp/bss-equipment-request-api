package retranslator

import (
	"context"
	"github.com/ozonmp/bss-equipment-request-api/internal/model"
	"sync"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/ozonmp/bss-equipment-request-api/internal/mocks"
)

func setUp(t *testing.T) (Config, *mocks.MockEventRepo, *mocks.MockEventSender, context.CancelFunc) {
	parentCtx := context.Background()
	ctrl := gomock.NewController(t)

	repo := mocks.NewMockEventRepo(ctrl)
	sender := mocks.NewMockEventSender(ctrl)

	ctx, ctxCancel := context.WithCancel(parentCtx)

	config := Config{
		ChannelSize:     512,
		ConsumerCount:   2,
		ConsumeSize:     10,
		ConsumeTimeout:  50 * time.Millisecond,
		ProducerCount:   2,
		ProducerTimeout: 50 * time.Millisecond,
		BatchSize:       1,
		WorkerCount:     2,
		Repo:            repo,
		Sender:          sender,
		Ctx:             ctx,
		CancelCtxFunc:   ctxCancel,
	}

	return config, repo, sender, ctxCancel
}

func tearDown(retranslator Retranslator, ctxFunc context.CancelFunc) {
	ctxFunc()
	retranslator.Close()
}

func TestStartAndGetOneEvent(t *testing.T) {
	t.Parallel()
	config, repo, sender, ctxFunc := setUp(t)
	retranslator := NewRetranslator(config)
	defer tearDown(retranslator, ctxFunc)

	event := model.EquipmentRequestEvent{
		ID:     12,
		Type:   model.Created,
		Status: model.Deferred,
		Entity: &model.EquipmentRequest{Id: 1, EmployeeId: 1, EquipmentType: "Laptop", EquipmentId: 1, CreatedAt: "2020-01-19T10:00:00", DoneAt: "2020-01-19T10:00:00", Status: true},
	}

	eventLockCount := int(config.ConsumerCount * config.BatchSize)
	var wg sync.WaitGroup
	wg.Add(eventLockCount)
	defer wg.Wait()

	repo.EXPECT().Lock(config.ConsumeSize).DoAndReturn(
		func(uint642 uint64) ([]model.EquipmentRequestEvent, error) {
			wg.Done()
			return []model.EquipmentRequestEvent{event}, nil
		}).Times(eventLockCount)

	eventSendCount := int(config.ProducerCount * config.BatchSize)
	removeCount := int(config.ProducerCount * config.BatchSize)

	var wgSender sync.WaitGroup
	wgSender.Add(eventSendCount)
	defer wgSender.Wait()

	var wgRepo sync.WaitGroup
	wgRepo.Add(removeCount)
	defer wgRepo.Wait()

	sender.EXPECT().Send(&event).DoAndReturn(
		func(*model.EquipmentRequestEvent) error {
			wgSender.Done()
			return nil
		}).Times(eventSendCount)

	repo.EXPECT().Remove([]uint64{event.ID}).DoAndReturn(
		func([]uint64) error {
			wgRepo.Done()
			return nil
		}).Times(removeCount)

	retranslator.Start()
}
