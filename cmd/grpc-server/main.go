package main

import (
	"context"
	"fmt"
	"github.com/ozonmp/bss-equipment-request-api/internal/logger"
	"github.com/ozonmp/bss-equipment-request-api/internal/metrics"
	"github.com/ozonmp/bss-equipment-request-api/internal/repo"
	"github.com/ozonmp/bss-equipment-request-api/internal/service/equipment_request"
	"time"

	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"

	"github.com/ozonmp/bss-equipment-request-api/internal/config"
	"github.com/ozonmp/bss-equipment-request-api/internal/database"
	"github.com/ozonmp/bss-equipment-request-api/internal/server"
	"github.com/ozonmp/bss-equipment-request-api/internal/tracer"
)

const grpsServerMainLogTag = "GrpsServerMain"

func main() {
	ctx := context.Background()

	if err := config.ReadConfigYML("config.yml"); err != nil {
		logger.FatalKV(ctx, fmt.Sprintf("%s: failed init configuration", grpsServerMainLogTag),
			"err", err,
		)
	}
	cfg := config.GetConfigInstance()

	syncLogger := logger.NewLogger(ctx, cfg.Project.Debug, cfg.Telemetry.GraylogPath, cfg.Project.ServiceName)
	defer syncLogger()

	metrics.InitMetrics(cfg)

	logger.InfoKV(ctx, fmt.Sprintf("%s: starting service: %s", grpsServerMainLogTag, cfg.Project.Name),
		"version", cfg.Project.Version,
		"commitHash", cfg.Project.CommitHash,
		"debug", cfg.Project.Debug,
		"environment", cfg.Project.Environment,
	)

	initCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SslMode,
	)

	db, err := database.NewPostgres(initCtx, dsn, cfg.Database.Driver, cfg.Database.ConnectAttempts)
	if err != nil {
		return
	}
	defer db.Close()

	tracing, err := tracer.NewTracer(ctx, cfg.Jaeger.Service, cfg.Jaeger.Host, cfg.Jaeger.Port)

	if err != nil {
		return
	}
	defer tracing.Close()

	requestRepository := repo.NewEquipmentRequestRepo(db)
	eventRepository := repo.NewEventRepo(db)

	equipmentRequestService := equipment_request.New(db, requestRepository, eventRepository)

	if err := server.NewGrpcServer(equipmentRequestService).Start(ctx, &cfg); err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: failed creating gRPC server", grpsServerMainLogTag), "err", err)

		return
	}
}
