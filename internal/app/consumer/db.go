package consumer

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/ozonmp/bss-equipment-request-api/internal/app/repo"
	"github.com/ozonmp/bss-equipment-request-api/internal/model"
)

type Consumer interface {
	Start()
	Close()
}

type consumer struct {
	n         uint64
	batchSize uint64
	timeout   time.Duration
	repo      repo.EventRepo
	ctx       context.Context
	events    chan<- model.EquipmentRequestEvent
	wg        *sync.WaitGroup
}

type Config struct {
	N         uint64
	Events    chan<- model.EquipmentRequestEvent
	Repo      repo.EventRepo
	BatchSize uint64
	Timeout   time.Duration
	Ctx       context.Context
}

func NewDbConsumer(
	n uint64,
	batchSize uint64,
	consumeTimeout time.Duration,
	repo repo.EventRepo,
	ctx context.Context,
	events chan<- model.EquipmentRequestEvent) Consumer {

	wg := &sync.WaitGroup{}

	return &consumer{
		n:         n,
		batchSize: batchSize,
		timeout:   consumeTimeout,
		repo:      repo,
		ctx:       ctx,
		events:    events,
		wg:        wg,
	}
}

func (c *consumer) Start() {
	for i := uint64(0); i < c.n; i++ {
		c.wg.Add(1)

		go func() {
			defer c.wg.Done()
			ticker := time.NewTicker(c.timeout)
			currentEvents := make([]model.EquipmentRequestEvent, 0, c.batchSize)
			for {
				select {
				case <-ticker.C:
					events, err := c.repo.Lock(c.batchSize)

					if err != nil {
						log.Printf("Unable to get and lock data from database: %v", err)
						continue
					}

					currentEvents = append(currentEvents, events...)

					for i, event := range currentEvents {
						c.events <- event
						if len(currentEvents) == 1 {
							currentEvents = currentEvents[:0]
						} else {
							currentEvents = append(currentEvents[:i], currentEvents[i+1:]...)
						}
					}
				case <-c.ctx.Done():
					if (len(currentEvents)) > 0 {
						eventIds := make([]uint64, 0, len(currentEvents))
						for _, event := range currentEvents {
							eventIds = append(eventIds, event.ID)
						}
						err := c.repo.Unlock(eventIds)
						if err != nil {
							log.Printf("Unable to unlock data from database: %v", err)
							return
						}
					}

					return
				}
			}
		}()
	}
}

func (c *consumer) Close() {
	c.wg.Wait()
}
