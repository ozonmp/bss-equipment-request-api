package sender

import (
	"github.com/ozonmp/bss-equipment-request-api/internal/model"
)

type EventSender interface {
	Send(subdomain *model.EquipmentRequestEvent) error
}
