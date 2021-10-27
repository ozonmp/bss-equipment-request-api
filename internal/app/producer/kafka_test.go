package producer

import (
	"context"
	"errors"
	"github.com/gammazero/workerpool"
	"github.com/golang/mock/gomock"
	"github.com/ozonmp/bss-equipment-request-api/internal/mocks"
	"github.com/ozonmp/bss-equipment-request-api/internal/model"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type KafkaTestSuite struct {
	suite.Suite
	producers   uint64
	ctrl        *gomock.Controller
	repo        *mocks.MockEventRepo
	sender      *mocks.MockEventSender
	ctx         context.Context
	ctxCancel   context.CancelFunc
	events      chan model.EquipmentRequestEvent
	kafka       Producer
	timeout     time.Duration
	channelSize int
	batchSize   uint64
	workers     int
}

func (suite *KafkaTestSuite) SetupTest() {
	parentCtx := context.Background()
	suite.ctrl = gomock.NewController(suite.T())
	suite.repo = mocks.NewMockEventRepo(suite.ctrl)
	suite.sender = mocks.NewMockEventSender(suite.ctrl)
	suite.ctx, suite.ctxCancel = context.WithCancel(parentCtx)
}

func (suite *KafkaTestSuite) setKafkaItem(producers uint64, channelSize int, timeout time.Duration, batchSize uint64, workers int) {
	suite.producers = producers
	suite.channelSize = channelSize
	suite.timeout = timeout
	suite.batchSize = batchSize
	suite.workers = workers

	suite.events = make(chan model.EquipmentRequestEvent, suite.channelSize)
	workerPool := workerpool.New(suite.workers)

	cfg := Config{
		n:          suite.producers,
		sender:     suite.sender,
		eventRepo:  suite.repo,
		timeout:    suite.timeout,
		events:     suite.events,
		ctx:        suite.ctx,
		batchSize:  suite.batchSize,
		workerPool: workerPool,
	}

	suite.kafka = NewKafkaProducer(
		cfg.n,
		cfg.sender,
		cfg.eventRepo,
		cfg.timeout,
		cfg.events,
		cfg.ctx,
		cfg.batchSize,
		cfg.workerPool)
}

func (suite *KafkaTestSuite) TestStart() {
	suite.setKafkaItem(2, 2, 2*time.Millisecond, 5, 2)

	suite.kafka.Start()
	defer func() {
		suite.ctxCancel()
		suite.kafka.Close()
	}()
}

func (suite *KafkaTestSuite) TestStartAndRemoveByDefer() {
	suite.setKafkaItem(2, 2, time.Second, 2, 2)

	event := model.EquipmentRequestEvent{
		ID:     12,
		Type:   model.Created,
		Status: model.Deferred,
		Entity: &model.EquipmentRequest{Id: 1, EmployeeId: 1, EquipmentType: "Laptop", EquipmentId: 1, CreatedAt: "2020-01-19T10:00:00", DoneAt: "2020-01-19T10:00:00", Status: true},
	}

	suite.sender.EXPECT().Send(&event).Return(nil).Times(1)
	suite.repo.EXPECT().Remove([]uint64{event.ID}).Times(1)

	suite.kafka.Start()

	suite.events <- event
	time.Sleep(time.Millisecond)

	defer func() {
		suite.ctxCancel()
		suite.kafka.Close()
	}()
}

func (suite *KafkaTestSuite) TestStartAndRemoveByTicker() {
	suite.setKafkaItem(2, 2, time.Second, 2, 2)

	event := model.EquipmentRequestEvent{
		ID:     12,
		Type:   model.Created,
		Status: model.Deferred,
		Entity: &model.EquipmentRequest{Id: 1, EmployeeId: 1, EquipmentType: "Laptop", EquipmentId: 1, CreatedAt: "2020-01-19T10:00:00", DoneAt: "2020-01-19T10:00:00", Status: true},
	}

	suite.sender.EXPECT().Send(&event).Return(nil).Times(1)
	suite.repo.EXPECT().Remove([]uint64{event.ID}).Times(1)

	suite.kafka.Start()

	suite.events <- event
	time.Sleep(2 * time.Second)

	defer func() {
		suite.ctxCancel()
		suite.kafka.Close()
	}()
}

func (suite *KafkaTestSuite) TestStartAndUnlockByDefer() {
	suite.setKafkaItem(2, 2, time.Second, 2, 2)

	event := model.EquipmentRequestEvent{
		ID:     12,
		Type:   model.Created,
		Status: model.Deferred,
		Entity: &model.EquipmentRequest{Id: 1, EmployeeId: 1, EquipmentType: "Laptop", EquipmentId: 1, CreatedAt: "2020-01-19T10:00:00", DoneAt: "2020-01-19T10:00:00", Status: true},
	}

	suite.sender.EXPECT().Send(&event).Return(errors.New("error during send")).Times(1)
	suite.repo.EXPECT().Unlock([]uint64{event.ID}).Times(1)

	suite.kafka.Start()

	suite.events <- event
	time.Sleep(time.Millisecond)

	defer func() {
		suite.ctxCancel()
		suite.kafka.Close()
	}()
}

func (suite *KafkaTestSuite) TestStartAndUnlockByTicker() {
	suite.setKafkaItem(2, 2, time.Second, 2, 2)

	event := model.EquipmentRequestEvent{
		ID:     12,
		Type:   model.Created,
		Status: model.Deferred,
		Entity: &model.EquipmentRequest{Id: 1, EmployeeId: 1, EquipmentType: "Laptop", EquipmentId: 1, CreatedAt: "2020-01-19T10:00:00", DoneAt: "2020-01-19T10:00:00", Status: true},
	}

	suite.sender.EXPECT().Send(&event).Return(errors.New("error during send")).Times(1)
	suite.repo.EXPECT().Unlock([]uint64{event.ID}).Times(1)

	suite.kafka.Start()

	suite.events <- event
	time.Sleep(2 * time.Second)

	defer func() {
		suite.ctxCancel()
		suite.kafka.Close()
	}()
}

func (suite *KafkaTestSuite) TestStartAndMultipleEvents() {
	suite.setKafkaItem(1, 2, time.Second, 2, 2)

	events := []model.EquipmentRequestEvent{
		{
			ID:     12,
			Type:   model.Created,
			Status: model.Deferred,
			Entity: &model.EquipmentRequest{Id: 1, EmployeeId: 1, EquipmentType: "Laptop", EquipmentId: 1, CreatedAt: "2020-01-19T10:00:00", DoneAt: "2020-01-19T10:00:00", Status: true},
		},
		{
			ID:     14,
			Type:   model.Created,
			Status: model.Deferred,
			Entity: &model.EquipmentRequest{Id: 1, EmployeeId: 1, EquipmentType: "Laptop", EquipmentId: 1, CreatedAt: "2020-01-19T10:00:00", DoneAt: "2020-01-19T10:00:00", Status: true},
		},
		{
			ID:     20,
			Type:   model.Created,
			Status: model.Deferred,
			Entity: &model.EquipmentRequest{Id: 1, EmployeeId: 1, EquipmentType: "Laptop", EquipmentId: 1, CreatedAt: "2020-01-19T10:00:00", DoneAt: "2020-01-19T10:00:00", Status: true},
		},
	}

	suite.sender.EXPECT().Send(&events[0]).Return(nil).Times(1)
	suite.sender.EXPECT().Send(&events[1]).Return(nil).Times(1)
	suite.sender.EXPECT().Send(&events[2]).Return(nil).Times(1)

	suite.repo.EXPECT().Remove([]uint64{12, 14}).Times(1)
	suite.repo.EXPECT().Remove([]uint64{20}).Times(1)

	suite.kafka.Start()

	for _, v := range events {
		suite.events <- v
		time.Sleep(time.Millisecond)
	}

	defer func() {
		suite.ctxCancel()
		suite.kafka.Close()
	}()
}

func TestKafkaTestSuite(t *testing.T) {
	suite.Run(t, new(KafkaTestSuite))
}
