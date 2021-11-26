package main

import (
	"context"
	"fmt"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"

	"github.com/ozonmp/bss-equipment-request-api/internal/database"
	"github.com/ozonmp/bss-equipment-request-api/internal/logger"
	"github.com/ozonmp/bss-equipment-request-api/internal/retranslator/config"
	"github.com/ozonmp/bss-equipment-request-api/internal/retranslator/metrics"
	"github.com/ozonmp/bss-equipment-request-api/internal/retranslator/repo"
	"github.com/ozonmp/bss-equipment-request-api/internal/retranslator/retranslator"
	"github.com/ozonmp/bss-equipment-request-api/internal/retranslator/sender"
	"github.com/ozonmp/bss-equipment-request-api/internal/server"
	"github.com/ozonmp/bss-equipment-request-api/internal/tracer"
	"time"
)

const retranslatorMainLogTag = "RetranslatorMain"

func main() {
	ctx := context.Background()

	if err := config.ReadConfigYML("config.yml"); err != nil {
		logger.FatalKV(ctx, fmt.Sprintf("%s: failed init configuration", retranslatorMainLogTag),
			"err", err,
		)
	}
	cfg := config.GetConfigInstance()

	syncLogger := logger.NewLogger(ctx, cfg.Project.Debug, cfg.Telemetry.GraylogPath, cfg.Project.ServiceName)
	defer syncLogger()

	metrics.InitMetrics(cfg)

	logger.InfoKV(ctx, fmt.Sprintf("Starting service: %s", cfg.Project.Name),
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

	eventRepository := repo.NewEventRepo(db)
	eventSender, err := sender.NewEventSender(ctx, cfg.Kafka.Brokers, cfg.Kafka.RetryMax, cfg.Kafka.RetryBackoff)
	if err != nil {
		logger.FatalKV(ctx, fmt.Sprintf("%s: sender.NewEventSender failed", retranslatorMainLogTag),
			"err", err,
		)
	}

	retranslator := retranslator.NewRetranslator(ctx, db, cfg.Retranslator, eventRepository, eventSender)

	if err := server.NewRetranslatorServer(retranslator).Start(ctx, &cfg); err != nil {
		logger.ErrorKV(ctx, fmt.Sprintf("%s: failed creating retranslator server", retranslatorMainLogTag),
			"err", err,
		)

		return
	}
}
