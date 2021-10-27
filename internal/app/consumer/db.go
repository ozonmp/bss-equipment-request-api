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
	n         uint64
	events    chan<- model.EquipmentRequestEvent
	repo      repo.EventRepo
	batchSize uint64
	timeout   time.Duration
	ctx       context.Context
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
			for {
				select {
				case <-ticker.C:
					events, err := c.repo.Lock(c.batchSize)
					if err != nil {
						log.Printf("Unable to get and lock data from database: %v", err)
						continue
					}
					for _, event := range events {
						c.events <- event
					}
				case <-c.ctx.Done():
					return
				}
			}
		}()
	}
}

func (c *consumer) Close() {
	c.wg.Wait()
}
