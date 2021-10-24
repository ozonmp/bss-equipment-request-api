package repo

import (
	"github.com/ozonmp/bss-equipment-request-api/internal/model"
)

type EventRepo interface {
	Lock(n uint64) ([]model.EquipmentRequestEvent, error)
	Unlock(eventIDs []uint64) error

	Add(event []model.EquipmentRequestEvent) error
	Remove(eventIDs []uint64) error
}
