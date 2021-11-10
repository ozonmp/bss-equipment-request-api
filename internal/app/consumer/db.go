package consumer

import (
	"context"
	"github.com/jmoiron/sqlx"
	"log"
	"sync"
	"time"

	"github.com/ozonmp/bss-equipment-request-api/internal/app/repo"
	"github.com/ozonmp/bss-equipment-request-api/internal/model"
)

// Consumer is a public interface for events consumers
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
	db        *sqlx.DB
}

// Config is a config for events consumer
type Config struct {
	N         uint64
	Events    chan<- model.EquipmentRequestEvent
	Repo      repo.EventRepo
	BatchSize uint64
	Timeout   time.Duration
	Ctx       context.Context
	DB        *sqlx.DB
}

// NewDbConsumer used to create a new db consumer
func NewDbConsumer(
	ctx context.Context,
	n uint64,
	batchSize uint64,
	consumeTimeout time.Duration,
	repo repo.EventRepo,
	events chan<- model.EquipmentRequestEvent,
	db *sqlx.DB) Consumer {

	wg := &sync.WaitGroup{}

	return &consumer{
		n:         n,
		batchSize: batchSize,
		timeout:   consumeTimeout,
		repo:      repo,
		ctx:       ctx,
		events:    events,
		wg:        wg,
		db:        db,
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
					events, txErr := c.repo.Lock(c.ctx, c.db, c.batchSize)

					if txErr != nil {
						log.Printf("Unable to get and lock data from database: %v", txErr)
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
						err := c.repo.Unlock(c.ctx, eventIds)
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
