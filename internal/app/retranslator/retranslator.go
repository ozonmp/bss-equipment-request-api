package retranslator

import (
	"context"
	"github.com/ozonmp/bss-equipment-request-api/internal/app/consumer"
	"github.com/ozonmp/bss-equipment-request-api/internal/app/producer"
	"github.com/ozonmp/bss-equipment-request-api/internal/app/repo"
	"github.com/ozonmp/bss-equipment-request-api/internal/app/sender"
	"github.com/ozonmp/bss-equipment-request-api/internal/model"
	"time"

	"github.com/gammazero/workerpool"
)

type Retranslator interface {
	Start()
	Close()
}

type Config struct {
	ChannelSize uint64

	ConsumerCount  uint64
	ConsumeSize    uint64
	ConsumeTimeout time.Duration

	ProducerCount   uint64
	ProducerTimeout time.Duration
	WorkerCount     int
	BatchSize       uint64

	Repo   repo.EventRepo
	Sender sender.EventSender

	Ctx           context.Context
	CancelCtxFunc context.CancelFunc
}

type retranslator struct {
	events        chan model.EquipmentRequestEvent
	consumer      consumer.Consumer
	producer      producer.Producer
	workerPool    *workerpool.WorkerPool
	cancelCtxFunc context.CancelFunc
}

func NewRetranslator(cfg Config) Retranslator {
	events := make(chan model.EquipmentRequestEvent, cfg.ChannelSize)
	workerPool := workerpool.New(cfg.WorkerCount)

	consumer := consumer.NewDbConsumer(
		cfg.ConsumerCount,
		cfg.ConsumeSize,
		cfg.ConsumeTimeout,
		cfg.Repo,
		cfg.Ctx,
		events)
	producer := producer.NewKafkaProducer(
		cfg.ProducerCount,
		cfg.Sender,
		cfg.Repo,
		cfg.ProducerTimeout,
		events,
		cfg.Ctx,
		cfg.BatchSize,
		workerPool)

	return &retranslator{
		events:        events,
		consumer:      consumer,
		producer:      producer,
		workerPool:    workerPool,
		cancelCtxFunc: cfg.CancelCtxFunc,
	}
}

func (r *retranslator) Start() {
	r.consumer.Start()
	r.producer.Start()
}

func (r *retranslator) Close() {
	r.cancelCtxFunc()
	r.consumer.Close()
	r.producer.Close()
	r.workerPool.StopWait()
	close(r.events)
}
