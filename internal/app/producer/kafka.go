package producer

import (
	"context"
	"github.com/ozonmp/bss-equipment-request-api/internal/app/repo"
	"log"
	"sync"
	"time"

	"github.com/ozonmp/bss-equipment-request-api/internal/app/sender"
	"github.com/ozonmp/bss-equipment-request-api/internal/model"

	"github.com/gammazero/workerpool"
)

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

type Config struct {
	n          uint64
	sender     sender.EventSender
	eventRepo  repo.EventRepo
	timeout    time.Duration
	events     <-chan model.EquipmentRequestEvent
	ctx        context.Context
	batchSize  uint64
	workerPool *workerpool.WorkerPool
}

func NewKafkaProducer(
	n uint64,
	sender sender.EventSender,
	eventRepo repo.EventRepo,
	timeout time.Duration,
	events <-chan model.EquipmentRequestEvent,
	ctx context.Context,
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
						log.Fatal("unable to read from the channel")
						return
					}
					if err := p.sender.Send(&event); err != nil {
						p.addToUnlockBatch(&toUnlockBatch, event.ID)
					} else {
						p.addToRemoveBatch(&toRemoveBatch, event.ID)
					}
				case <-ticker.C:
					p.sendToUnlockBatch(&toUnlockBatch)
					p.sendToRemoveBatch(&toRemoveBatch)
				case <-p.ctx.Done():
					return
				}
			}
		}()
	}
}

func (p *producer) addToUnlockBatch(toUnlockBatch *[]uint64, eventId uint64) {
	if uint64(len(*toUnlockBatch)) < p.batchSize {
		*toUnlockBatch = append(*toUnlockBatch, eventId)

		if uint64(len(*toUnlockBatch)) == p.batchSize {
			p.sendToUnlockBatch(toUnlockBatch)
		}
	} else {
		p.sendToUnlockBatch(toUnlockBatch)
		*toUnlockBatch = append(*toUnlockBatch, eventId)
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

func (p *producer) addToRemoveBatch(toRemoveBatch *[]uint64, eventId uint64) {

	if uint64(len(*toRemoveBatch)) < p.batchSize {
		*toRemoveBatch = append(*toRemoveBatch, eventId)
		if uint64(len(*toRemoveBatch)) == p.batchSize {
			p.sendToRemoveBatch(toRemoveBatch)
		}
	} else {
		p.sendToRemoveBatch(toRemoveBatch)
		*toRemoveBatch = append(*toRemoveBatch, eventId)
	}
}

func (p *producer) unlockBatch(toUnlockBatch []uint64) {
	err := p.repo.Unlock(toUnlockBatch)
	if err != nil {
		log.Fatalf("unable to Unlock %v, %v", toUnlockBatch, err)
	}
}

func (p *producer) removeBatch(toRemoveBatch []uint64) {
	err := p.repo.Remove(toRemoveBatch)
	if err != nil {
		log.Fatalf("unable to Remove %v, %v", toRemoveBatch, err)
	}
}

func (p *producer) cleanBatch(batch *[]uint64) {
	*batch = (*batch)[:0]
}

func (p *producer) Close() {
	p.wg.Wait()
}
