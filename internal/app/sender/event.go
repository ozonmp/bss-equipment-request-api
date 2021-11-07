package sender

import (
	"github.com/ozonmp/bss-equipment-request-api/internal/model"
)

// EventSender is sender of events
type EventSender interface {
	Send(subdomain *model.EquipmentRequestEvent) error
}
