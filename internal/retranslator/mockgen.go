package internal

//go:generate mockgen -destination=./mocks/event_repo_mock.go -package=mocks github.com/ozonmp/bss-equipment-request-api/retranslator/internal/repo EventRepo
//go:generate mockgen -destination=./mocks/sender_mock.go -package=mocks github.com/ozonmp/bss-equipment-request-api/retranslator/internal/sender EventSender
