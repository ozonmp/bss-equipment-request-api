package retranslator

import (
	"context"
	"errors"
	"github.com/ozonmp/bss-equipment-request-api/internal/model"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/ozonmp/bss-equipment-request-api/internal/mocks"
)

type RetranslatorTestSuite struct {
	suite.Suite
	ctrl         *gomock.Controller
	repo         *mocks.MockEventRepo
	sender       *mocks.MockEventSender
	ctx          context.Context
	ctxCancel    context.CancelFunc
	retranslator Retranslator
}

func (suite *RetranslatorTestSuite) SetupTest() {
	parentCtx := context.Background()
	suite.ctrl = gomock.NewController(suite.T())
	suite.repo = mocks.NewMockEventRepo(suite.ctrl)
	suite.sender = mocks.NewMockEventSender(suite.ctrl)
	suite.ctx, suite.ctxCancel = context.WithCancel(parentCtx)
}

func (suite *RetranslatorTestSuite) setRetranslatorItem() {
	cfg := Config{
		ChannelSize:     512,
		ConsumerCount:   2,
		ConsumeSize:     10,
		ConsumeTimeout:  50 * time.Millisecond,
		ProducerCount:   2,
		ProducerTimeout: 50 * time.Millisecond,
		BatchSize:       5,
		WorkerCount:     2,
		Repo:            suite.repo,
		Sender:          suite.sender,
		Ctx:             suite.ctx,
		CancelCtxFunc:   suite.ctxCancel,
	}

	suite.retranslator = NewRetranslator(cfg)
}

func (suite *RetranslatorTestSuite) TestStart() {
	suite.setRetranslatorItem()

	suite.retranslator.Start()
	defer suite.retranslator.Close()
}

func (suite *RetranslatorTestSuite) TestStartAndWork() {
	suite.setRetranslatorItem()

	event := model.EquipmentRequestEvent{
		ID:     12,
		Type:   model.Created,
		Status: model.Deferred,
		Entity: &model.EquipmentRequest{Id: 1, EmployeeId: 1, EquipmentType: "Laptop", EquipmentId: 1, CreatedAt: "2020-01-19T10:00:00", DoneAt: "2020-01-19T10:00:00", Status: true},
	}

	suite.repo.EXPECT().Lock(uint64(10)).
		Return([]model.EquipmentRequestEvent{event}, nil).Times(1)
	suite.repo.EXPECT().Lock(gomock.Any()).Return(nil, errors.New("empty result error")).AnyTimes()

	suite.sender.EXPECT().Send(&event).Return(nil).Times(1)
	suite.repo.EXPECT().Remove([]uint64{event.ID}).Times(1)

	suite.retranslator.Start()
	time.Sleep(time.Second)
	defer suite.retranslator.Close()
}

func TestRetranslatorTestSuite(t *testing.T) {
	suite.Run(t, new(RetranslatorTestSuite))
}
