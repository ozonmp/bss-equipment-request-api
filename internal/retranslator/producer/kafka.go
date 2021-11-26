package producer

import (
	"context"
	"fmt"
	"github.com/ozonmp/bss-equipment-request-api/internal/logger"
	"github.com/ozonmp/bss-equipment-request-api/internal/retranslator/repo"
	"sync"
	"time"

	"github.com/ozonmp/bss-equipment-request-api/internal/model"
	"github.com/ozonmp/bss-equipment-request-api/internal/retranslator/sender"

	"github.com/gammazero/workerpool"
)

const producerLogTag = "Producer"

// Producer is a public interface for events producers
type Producer interface {
	Start()
	Close()
}

type producer struct {
	n          uint64
	sender     sender.EventSender
	repo       repo.EventRepo
	timeout    time.Duration
	events     <-chan model.EquipmentRequestEvent
	ctx        context.Context
	batchSize  uint64
	workerPool *workerpool.WorkerPool
	wg         *sync.WaitGroup
}

// Config is a config for events producer
type Config struct {
	N          uint64
	Sender     sender.EventSender
	EventRepo  repo.EventRepo
	Timeout    time.Duration
	Events     <-chan model.EquipmentRequestEvent
	Ctx        context.Context
	BatchSize  uint64
	WorkerPool *workerpool.WorkerPool
}

// NewKafkaProducer used to create a new kafka producer
func NewKafkaProducer(
	ctx context.Context,
	n uint64,
	sender sender.EventSender,
	eventRepo repo.EventRepo,
	timeout time.Duration,
	events <-chan model.EquipmentRequestEvent,
	batchSize uint64,
	workerPool *workerpool.WorkerPool,
) Producer {

	wg := &sync.WaitGroup{}

	return &producer{
		n:          n,
		sender:     sender,
		repo:       eventRepo,
		timeout:    timeout,
		events:     events,
		ctx:        ctx,
		batchSize:  batchSize,
		workerPool: workerPool,
		wg:         wg,
	}
}

func (p *producer) Start() {
	for i := uint64(0); i < p.n; i++ {
		p.wg.Add(1)
		go func() {
			ticker := time.NewTicker(p.timeout)

			var toUnlockBatch = make([]uint64, 0, p.batchSize)
			var toRemoveBatch = make([]uint64, 0, p.batchSize)

			defer func() {
				p.wg.Done()
				p.sendToUnlockBatch(&toUnlockBatch)
				p.sendToRemoveBatch(&toRemoveBatch)
			}()

			for {
				select {
				case event, ok := <-p.events:
					if !ok {
						logger.FatalKV(p.ctx, fmt.Sprintf("%s: read from p.events failed", producerLogTag),
							"err", "unable to read from the channel",
						)
					}
					if err := p.sender.Send(&event); err != nil {
						p.addToUnlockBatch(&toUnlockBatch, event.ID)
						logger.ErrorKV(p.ctx, fmt.Sprintf("%s: sender.Send failed", producerLogTag),
							"err", err,
						)
					} else {
						p.addToRemoveBatch(&toRemoveBatch, event.ID)
					}
				case <-ticker.C:
					p.sendToUnlockBatch(&toUnlockBatch)
					p.sendToRemoveBatch(&toRemoveBatch)
				case <-p.ctx.Done():
					for len(p.events) > 0 {
						event, ok := <-p.events
						if !ok {
							logger.FatalKV(p.ctx, fmt.Sprintf("%s: read from p.events failed", producerLogTag),
								"err", "unable to read from the channel",
							)
							return
						}
						if err := p.sender.Send(&event); err != nil {
							logger.ErrorKV(p.ctx, fmt.Sprintf("%s: sender.Send failed", producerLogTag),
								"err", err,
							)
							p.addToUnlockBatch(&toUnlockBatch, event.ID)
						} else {
							p.addToRemoveBatch(&toRemoveBatch, event.ID)
						}
					}

					return
				}
			}
		}()
	}
}

func (p *producer) addToUnlockBatch(toUnlockBatch *[]uint64, eventID uint64) {
	if uint64(len(*toUnlockBatch)) < p.batchSize {
		*toUnlockBatch = append(*toUnlockBatch, eventID)

		if uint64(len(*toUnlockBatch)) == p.batchSize {
			p.sendToUnlockBatch(toUnlockBatch)
		}
	} else {
		p.sendToUnlockBatch(toUnlockBatch)
		*toUnlockBatch = append(*toUnlockBatch, eventID)
	}
}

func (p *producer) sendToUnlockBatch(toUnlockBatch *[]uint64) {
	if uint64(len(*toUnlockBatch)) > 0 {
		var tmp = make([]uint64, len(*toUnlockBatch))
		copy(tmp, *toUnlockBatch)
		p.workerPool.Submit(func() {
			p.unlockBatch(tmp)
		})

		p.cleanBatch(toUnlockBatch)
	}
}

func (p *producer) sendToRemoveBatch(toRemoveBatch *[]uint64) {
	if uint64(len(*toRemoveBatch)) > 0 {
		var tmp = make([]uint64, len(*toRemoveBatch))
		copy(tmp, *toRemoveBatch)
		p.workerPool.Submit(func() {
			p.removeBatch(tmp)
		})

		p.cleanBatch(toRemoveBatch)
	}
}

func (p *producer) addToRemoveBatch(toRemoveBatch *[]uint64, eventID uint64) {
	if uint64(len(*toRemoveBatch)) < p.batchSize {
		*toRemoveBatch = append(*toRemoveBatch, eventID)
		if uint64(len(*toRemoveBatch)) == p.batchSize {
			p.sendToRemoveBatch(toRemoveBatch)
		}
	} else {
		p.sendToRemoveBatch(toRemoveBatch)
		*toRemoveBatch = append(*toRemoveBatch, eventID)
	}
}

func (p *producer) unlockBatch(toUnlockBatch []uint64) {
	err := p.repo.Unlock(p.ctx, toUnlockBatch)
	if err != nil {
		logger.FatalKV(p.ctx, fmt.Sprintf("%s: repo.Unlock failed", producerLogTag),
			"err", err,
			"toUnlockBatch", toUnlockBatch,
		)
	}
}

func (p *producer) removeBatch(toRemoveBatch []uint64) {
	err := p.repo.Remove(p.ctx, toRemoveBatch)
	if err != nil {
		logger.FatalKV(p.ctx, fmt.Sprintf("%s: repo.Remove failed", producerLogTag),
			"err", err,
			"toRemoveBatch", toRemoveBatch,
		)
	}
}

func (p *producer) cleanBatch(batch *[]uint64) {
	*batch = (*batch)[:0]
}

func (p *producer) Close() {
	p.wg.Wait()
}
