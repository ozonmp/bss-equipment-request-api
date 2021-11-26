package retranslator

import (
	"context"
	"github.com/gammazero/workerpool"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/bss-equipment-request-api/internal/model"
	"github.com/ozonmp/bss-equipment-request-api/internal/retranslator/config"
	"github.com/ozonmp/bss-equipment-request-api/internal/retranslator/consumer"
	"github.com/ozonmp/bss-equipment-request-api/internal/retranslator/producer"
	"github.com/ozonmp/bss-equipment-request-api/internal/retranslator/repo"
	"github.com/ozonmp/bss-equipment-request-api/internal/retranslator/sender"
)

// Retranslator is a public interface for events translators
type Retranslator interface {
	Start()
	Close()
}

type retranslator struct {
	events        chan model.EquipmentRequestEvent
	consumer      consumer.Consumer
	producer      producer.Producer
	workerPool    *workerpool.WorkerPool
	cancelCtxFunc context.CancelFunc
}

// NewRetranslator used to create a new retranslator
func NewRetranslator(ctx context.Context, db *sqlx.DB, cfg config.Retranslator, repo repo.EventRepo, sender sender.EventSender) Retranslator {
	ctx, cancelCtxFunc := context.WithCancel(ctx)

	events := make(chan model.EquipmentRequestEvent, cfg.ChannelSize)
	workerPool := workerpool.New(cfg.WorkerCount)

	consumer := consumer.NewDbConsumer(
		ctx,
		cfg.ConsumerCount,
		cfg.ConsumeSize,
		cfg.ConsumeTimeout,
		repo,
		events,
		db)
	producer := producer.NewKafkaProducer(
		ctx,
		cfg.ProducerCount,
		sender,
		repo,
		cfg.ProducerTimeout,
		events,
		cfg.BatchSize,
		workerPool)

	return &retranslator{
		events:        events,
		consumer:      consumer,
		producer:      producer,
		workerPool:    workerPool,
		cancelCtxFunc: cancelCtxFunc,
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
