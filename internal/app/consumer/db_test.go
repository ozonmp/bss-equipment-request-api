package consumer

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/ozonmp/bss-equipment-request-api/internal/mocks"
	"github.com/ozonmp/bss-equipment-request-api/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type DbTestSuite struct {
	suite.Suite
	consumers   uint64
	ctrl        *gomock.Controller
	repo        *mocks.MockEventRepo
	ctx         context.Context
	ctxCancel   context.CancelFunc
	events      chan model.EquipmentRequestEvent
	db          Consumer
	timeout     time.Duration
	channelSize int
	batchSize   uint64
}

func (suite *DbTestSuite) SetupTest() {
	parentCtx := context.Background()
	suite.ctrl = gomock.NewController(suite.T())
	suite.repo = mocks.NewMockEventRepo(suite.ctrl)
	suite.ctx, suite.ctxCancel = context.WithCancel(parentCtx)
}

func (suite *DbTestSuite) setDbItem(consumers uint64, channelSize int, timeout time.Duration, batchSize uint64) {
	suite.consumers = consumers
	suite.channelSize = channelSize
	suite.timeout = timeout
	suite.batchSize = batchSize

	suite.events = make(chan model.EquipmentRequestEvent, suite.channelSize)

	cfg := Config{
		n:         suite.consumers,
		events:    suite.events,
		repo:      suite.repo,
		batchSize: suite.batchSize,
		timeout:   suite.timeout,
		ctx:       suite.ctx,
	}

	suite.db = NewDbConsumer(
		cfg.n,
		cfg.batchSize,
		cfg.timeout,
		cfg.repo,
		cfg.ctx,
		suite.events)
}

func (suite *DbTestSuite) TestStart() {
	suite.setDbItem(2, 1, time.Millisecond, 1)
	suite.repo.EXPECT().Lock(suite.batchSize).
		Return([]model.EquipmentRequestEvent{}, nil).AnyTimes()

	suite.db.Start()
	defer func() {
		suite.ctxCancel()
		suite.db.Close()
	}()
}

func (suite *DbTestSuite) TestStartAndGetOneEvent() {
	suite.setDbItem(2, 2, time.Millisecond, 1)

	event := model.EquipmentRequestEvent{
		ID:     12,
		Type:   model.Created,
		Status: model.Deferred,
		Entity: &model.EquipmentRequest{Id: 1, EmployeeId: 1, EquipmentType: "Laptop", EquipmentId: 1, CreatedAt: "2020-01-19T10:00:00", DoneAt: "2020-01-19T10:00:00", Status: true},
	}

	suite.repo.EXPECT().Lock(suite.batchSize).
		Return([]model.EquipmentRequestEvent{event}, nil).Times(1)
	suite.repo.EXPECT().Lock(gomock.Any()).Return(nil, errors.New("empty result error")).AnyTimes()

	suite.db.Start()
	ev := <-suite.events

	assert.Equal(suite.T(), event.ID, ev.ID, "The two events should be the same.")

	defer func() {
		suite.ctxCancel()
		suite.db.Close()
	}()
}

func (suite *DbTestSuite) TestStartAndGetErrors() {
	suite.setDbItem(2, 2, time.Millisecond, 1)

	suite.repo.EXPECT().Lock(suite.batchSize).Return(nil, errors.New("empty result error")).AnyTimes()

	suite.db.Start()

	defer func() {
		suite.ctxCancel()
		suite.db.Close()
	}()
}

func (suite *DbTestSuite) TestStartAndGetSeveralEvent() {
	suite.setDbItem(1, 2, 5*time.Millisecond, 2)

	events := []model.EquipmentRequestEvent{
		{
			ID:     12,
			Type:   model.Created,
			Status: model.Deferred,
			Entity: &model.EquipmentRequest{Id: 1, EmployeeId: 1, EquipmentType: "Laptop", EquipmentId: 1, CreatedAt: "2020-01-19T10:00:00", DoneAt: "2020-01-19T10:00:00", Status: true},
		},
		{
			ID:     14,
			Type:   model.Updated,
			Status: model.Processed,
			Entity: &model.EquipmentRequest{Id: 12, EmployeeId: 3, EquipmentType: "Laptop2", EquipmentId: 1, CreatedAt: "2020-01-19T10:00:00", DoneAt: "2020-01-19T10:00:00", Status: true},
		},
	}

	suite.repo.EXPECT().Lock(suite.batchSize).
		Return(events, nil).Times(1)
	suite.repo.EXPECT().Lock(gomock.Any()).Return(nil, errors.New("empty result error")).AnyTimes()

	suite.db.Start()

	ev := <-suite.events
	assert.Equal(suite.T(), events[0].ID, ev.ID, "The two events should be the same.")

	ev1 := <-suite.events
	assert.Equal(suite.T(), events[1].ID, ev1.ID, "The two events should be the same.")

	defer func() {
		suite.ctxCancel()
		suite.db.Close()
	}()
}

func TestDbTestSuite(t *testing.T) {
	suite.Run(t, new(DbTestSuite))
}
