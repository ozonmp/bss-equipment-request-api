package retranslator

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/bss-equipment-request-api/internal/app/consumer"
	"github.com/ozonmp/bss-equipment-request-api/internal/app/producer"
	"github.com/ozonmp/bss-equipment-request-api/internal/app/repo"
	"github.com/ozonmp/bss-equipment-request-api/internal/app/sender"
	"github.com/ozonmp/bss-equipment-request-api/internal/model"
	"time"

	"github.com/gammazero/workerpool"
)

// Retranslator is a public interface for events translators
type Retranslator interface {
	Start()
	Close()
}

// Config is a config for events retranslator
type Config struct {
	DB          *sqlx.DB
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

// NewRetranslator used to create a new retranslator
func NewRetranslator(cfg Config) Retranslator {
	events := make(chan model.EquipmentRequestEvent, cfg.ChannelSize)
	workerPool := workerpool.New(cfg.WorkerCount)

	consumer := consumer.NewDbConsumer(
		cfg.Ctx,
		cfg.ConsumerCount,
		cfg.ConsumeSize,
		cfg.ConsumeTimeout,
		cfg.Repo,
		events,
		cfg.DB)
	producer := producer.NewKafkaProducer(
		cfg.Ctx,
		cfg.ProducerCount,
		cfg.Sender,
		cfg.Repo,
		cfg.ProducerTimeout,
		events,
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
