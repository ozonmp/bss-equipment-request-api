package internal

//go:generate mockgen -destination=./mocks/repo_mock.go -package=mocks github.com/ozonmp/bss-equipment-request-api/internal/repo EquipmentRequestRepo
//go:generate mockgen -destination=./mocks/event_repo_mock.go -package=mocks github.com/ozonmp/bss-equipment-request-api/internal/repo EventRepo
