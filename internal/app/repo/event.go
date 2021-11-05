package repo

import (
	"github.com/ozonmp/bss-equipment-request-api/internal/model"
)

// EventRepo is repository for events
type EventRepo interface {

	/**
	If we want to support multiple types of events given the correct order
	we can use this query inside Lock function to get each time only the first and not locked events of entity by its type
	WITH added_row_number AS (
	  SELECT
		*,
		ROW_NUMBER() OVER(PARTITION BY entity_id order by type, id) AS row_number
	  FROM events
	)
	SELECT
	  *
	FROM added_row_number
	WHERE row_number = 1;
	*/

	Lock(n uint64) ([]model.EquipmentRequestEvent, error)
	Unlock(eventIDs []uint64) error

	Add(event []model.EquipmentRequestEvent) error
	Remove(eventIDs []uint64) error
}
