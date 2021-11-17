package internal

//go:generate mockgen -destination=./mocks/repo_mock.go -package=mocks github.com/ozonmp/bss-equipment-request-api/internal/app/repo EventRepo
//go:generate mockgen -destination=./mocks/sender_mock.go -package=mocks github.com/ozonmp/bss-equipment-request-api/internal/app/sender EventSender
//go:generate mockgen -destination=./mocks/server/repo_mock.go -package=mocks github.com/ozonmp/bss-equipment-request-api/internal/repo EquipmentRequestRepo
//go:generate mockgen -destination=./mocks/server/event_repo_mock.go -package=mocks github.com/ozonmp/bss-equipment-request-api/internal/repo EventRepo
