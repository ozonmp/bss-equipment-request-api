package sender

import (
	"context"
	"github.com/ozonmp/bss-equipment-request-api/internal/retranslator/logger"
	"github.com/ozonmp/bss-equipment-request-api/internal/retranslator/model"
)

// EventSender is sender of events
type EventSender interface {
	Send(subdomain *model.EquipmentRequestEvent) error
}

type eventSender struct {
	ctx context.Context
}

// NewEventSender returns EventSender interface
func NewEventSender(ctx context.Context) EventSender {
	return &eventSender{
		ctx: ctx,
	}
}

func (e eventSender) Send(subdomain *model.EquipmentRequestEvent) error {
	logger.InfoKV(e.ctx, "EventSender: send",
		"subdomain", subdomain,
	)

	return nil
}
