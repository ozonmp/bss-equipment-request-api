package producer

import (
	"context"
	"errors"
	"github.com/gammazero/workerpool"
	"github.com/golang/mock/gomock"
	"github.com/ozonmp/bss-equipment-request-api/internal/mocks"
	"github.com/ozonmp/bss-equipment-request-api/internal/model"
	"sync"
	"testing"
	"time"
)

func setUp(t *testing.T, producers uint64, channelSize int, timeout time.Duration, batchSize uint64, workers int) (Config, chan model.EquipmentRequestEvent, *mocks.MockEventRepo, *mocks.MockEventSender, context.CancelFunc) {
	parentCtx := context.Background()
	ctrl := gomock.NewController(t)

	repo := mocks.NewMockEventRepo(ctrl)
	sender := mocks.NewMockEventSender(ctrl)

	ctx, ctxCancel := context.WithCancel(parentCtx)

	events := make(chan model.EquipmentRequestEvent, channelSize)
	workerPool := workerpool.New(workers)

	config := Config{
		N:          producers,
		Sender:     sender,
		EventRepo:  repo,
		Timeout:    timeout,
		Events:     events,
		Ctx:        ctx,
		BatchSize:  batchSize,
		WorkerPool: workerPool,
	}

	return config, events, repo, sender, ctxCancel
}

func tearDown(kafka Producer, ctxFunc context.CancelFunc) {
	ctxFunc()
	kafka.Close()
}

func TestStartAndGetOneEvent(t *testing.T) {
	t.Parallel()
	config, events, repo, sender, ctxFunc := setUp(t, 2, 2, time.Millisecond, 1, 2)
	kafka := NewKafkaProducer(
		config.N,
		config.Sender,
		config.EventRepo,
		config.Timeout,
		config.Events,
		config.Ctx,
		config.BatchSize,
		config.WorkerPool)
	defer tearDown(kafka, ctxFunc)

	event := model.EquipmentRequestEvent{
		ID:     12,
		Type:   model.Created,
		Status: model.Deferred,
		Entity: &model.EquipmentRequest{Id: 1, EmployeeId: 1, EquipmentType: "Laptop", EquipmentId: 1, CreatedAt: "2020-01-19T10:00:00", DoneAt: "2020-01-19T10:00:00", Status: true},
	}

	evensCount := int(config.N * config.BatchSize)
	removeCount := int(config.N)

	var wgSender sync.WaitGroup
	wgSender.Add(evensCount)
	defer wgSender.Wait()

	var wgRepo sync.WaitGroup
	wgRepo.Add(removeCount)
	defer wgRepo.Wait()

	sender.EXPECT().Send(&event).DoAndReturn(
		func(*model.EquipmentRequestEvent) error {
			wgSender.Done()
			return nil
		}).Times(evensCount)

	repo.EXPECT().Remove([]uint64{event.ID}).DoAndReturn(
		func([]uint64) error {
			wgRepo.Done()
			return nil
		}).Times(removeCount)

	kafka.Start()
	for i := 0; i < evensCount; i++ {
		events <- event
	}
}

func TestStartAndRemoveByTicker(t *testing.T) {
	t.Parallel()
	config, events, repo, sender, ctxFunc := setUp(t, 2, 2, time.Millisecond, 1, 2)
	kafka := NewKafkaProducer(
		config.N,
		config.Sender,
		config.EventRepo,
		config.Timeout,
		config.Events,
		config.Ctx,
		config.BatchSize,
		config.WorkerPool)
	defer tearDown(kafka, ctxFunc)

	event := model.EquipmentRequestEvent{
		ID:     12,
		Type:   model.Created,
		Status: model.Deferred,
		Entity: &model.EquipmentRequest{Id: 1, EmployeeId: 1, EquipmentType: "Laptop", EquipmentId: 1, CreatedAt: "2020-01-19T10:00:00", DoneAt: "2020-01-19T10:00:00", Status: true},
	}

	evensCount := int(config.N * config.BatchSize)
	removeCount := int(config.N)

	var wgSender sync.WaitGroup
	wgSender.Add(evensCount)
	defer wgSender.Wait()

	var wgRepo sync.WaitGroup
	wgRepo.Add(removeCount)
	defer wgRepo.Wait()

	sender.EXPECT().Send(&event).DoAndReturn(
		func(*model.EquipmentRequestEvent) error {
			wgSender.Done()
			time.Sleep(time.Second)
			return nil
		}).Times(evensCount)

	repo.EXPECT().Remove([]uint64{event.ID}).DoAndReturn(
		func([]uint64) error {
			wgRepo.Done()
			return nil
		}).Times(removeCount)

	kafka.Start()
	for i := 0; i < evensCount; i++ {
		events <- event
	}
}

func TestStartAndRemoveByDefer(t *testing.T) {
	t.Parallel()
	config, events, repo, sender, ctxFunc := setUp(t, 2, 5, time.Millisecond, 1, 2)
	kafka := NewKafkaProducer(
		config.N,
		config.Sender,
		config.EventRepo,
		config.Timeout,
		config.Events,
		config.Ctx,
		config.BatchSize,
		config.WorkerPool)
	defer kafka.Close()

	event := model.EquipmentRequestEvent{
		ID:     1,
		Type:   model.Created,
		Status: model.Deferred,
		Entity: &model.EquipmentRequest{Id: 1, EmployeeId: 1, EquipmentType: "Laptop", EquipmentId: 1, CreatedAt: "2020-01-19T10:00:00", DoneAt: "2020-01-19T10:00:00", Status: true},
	}

	evensCount := int(config.N * config.BatchSize)
	removeCount := int(config.N)

	var wgSender sync.WaitGroup
	wgSender.Add(evensCount)
	defer wgSender.Wait()

	var wgRepo sync.WaitGroup
	wgRepo.Add(removeCount)
	defer wgRepo.Wait()

	sender.EXPECT().Send(&event).DoAndReturn(
		func(*model.EquipmentRequestEvent) error {
			wgSender.Done()
			return nil
		}).Times(evensCount)

	repo.EXPECT().Remove([]uint64{event.ID}).DoAndReturn(
		func([]uint64) error {
			wgRepo.Done()
			return nil
		}).Times(removeCount)

	kafka.Start()
	for i := 0; i < evensCount; i++ {
		events <- event

		if i == removeCount/2 {
			ctxFunc()
		}
	}
}

func TestStartAndUnlockByTicker(t *testing.T) {
	t.Parallel()
	config, events, repo, sender, ctxFunc := setUp(t, 2, 2, time.Millisecond, 1, 2)
	kafka := NewKafkaProducer(
		config.N,
		config.Sender,
		config.EventRepo,
		config.Timeout,
		config.Events,
		config.Ctx,
		config.BatchSize,
		config.WorkerPool)
	defer tearDown(kafka, ctxFunc)

	event := model.EquipmentRequestEvent{
		ID:     12,
		Type:   model.Created,
		Status: model.Deferred,
		Entity: &model.EquipmentRequest{Id: 1, EmployeeId: 1, EquipmentType: "Laptop", EquipmentId: 1, CreatedAt: "2020-01-19T10:00:00", DoneAt: "2020-01-19T10:00:00", Status: true},
	}

	evensCount := int(config.N * config.BatchSize)
	unlockCount := int(config.N)

	var wgSender sync.WaitGroup
	wgSender.Add(evensCount)
	defer wgSender.Wait()

	var wgRepo sync.WaitGroup
	wgRepo.Add(unlockCount)
	defer wgRepo.Wait()

	sender.EXPECT().Send(&event).DoAndReturn(
		func(*model.EquipmentRequestEvent) error {
			wgSender.Done()
			time.Sleep(time.Second)
			return errors.New("error during send")
		}).Times(evensCount)

	repo.EXPECT().Unlock([]uint64{event.ID}).DoAndReturn(
		func([]uint64) error {
			wgRepo.Done()
			return nil
		}).Times(unlockCount)

	kafka.Start()
	for i := 0; i < evensCount; i++ {
		events <- event
	}
}

func TestStartAndUnlockByDefer(t *testing.T) {
	t.Parallel()
	config, events, repo, sender, ctxFunc := setUp(t, 4, 5, time.Millisecond, 1, 2)
	kafka := NewKafkaProducer(
		config.N,
		config.Sender,
		config.EventRepo,
		config.Timeout,
		config.Events,
		config.Ctx,
		config.BatchSize,
		config.WorkerPool)
	defer kafka.Close()

	event := model.EquipmentRequestEvent{
		ID:     1,
		Type:   model.Created,
		Status: model.Deferred,
		Entity: &model.EquipmentRequest{Id: 1, EmployeeId: 1, EquipmentType: "Laptop", EquipmentId: 1, CreatedAt: "2020-01-19T10:00:00", DoneAt: "2020-01-19T10:00:00", Status: true},
	}

	evensCount := int(config.N * config.BatchSize)
	unlockCount := int(config.N)

	var wgSender sync.WaitGroup
	wgSender.Add(evensCount)
	defer wgSender.Wait()

	var wgRepo sync.WaitGroup
	wgRepo.Add(unlockCount)
	defer wgRepo.Wait()

	sender.EXPECT().Send(&event).DoAndReturn(
		func(*model.EquipmentRequestEvent) error {
			wgSender.Done()
			return errors.New("error during send")
		}).Times(evensCount)

	repo.EXPECT().Unlock([]uint64{event.ID}).DoAndReturn(
		func([]uint64) error {
			wgRepo.Done()
			return nil
		}).Times(unlockCount)

	kafka.Start()
	for i := 0; i < evensCount; i++ {
		events <- event

		if i == unlockCount/2 {
			ctxFunc()
		}
	}
}
