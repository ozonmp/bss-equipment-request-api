package main

import (
	"context"
	"fmt"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"

	"github.com/ozonmp/bss-equipment-request-api/internal/retranslator/config"
	"github.com/ozonmp/bss-equipment-request-api/internal/retranslator/database"
	"github.com/ozonmp/bss-equipment-request-api/internal/retranslator/logger"
	"github.com/ozonmp/bss-equipment-request-api/internal/retranslator/metrics"
	"github.com/ozonmp/bss-equipment-request-api/internal/retranslator/repo"
	"github.com/ozonmp/bss-equipment-request-api/internal/retranslator/retranslator"
	"github.com/ozonmp/bss-equipment-request-api/internal/retranslator/sender"
	"github.com/ozonmp/bss-equipment-request-api/internal/retranslator/server"
	"github.com/ozonmp/bss-equipment-request-api/internal/retranslator/tracer"
	"time"
)

func main() {
	ctx := context.Background()

	if err := config.ReadConfigYML("config.yml"); err != nil {
		logger.FatalKV(ctx, "Failed init configuration", "err", err)
	}
	cfg := config.GetConfigInstance()

	syncLogger := logger.NewLogger(ctx, cfg)
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

	db, err := database.NewPostgres(initCtx, dsn, cfg.Database.Driver)
	if err != nil {
		logger.ErrorKV(ctx, "failed init postgres", "err", err)

		return
	}
	defer db.Close()

	tracing, err := tracer.NewTracer(ctx, &cfg)

	if err != nil {
		logger.ErrorKV(ctx, "failed init tracing", "err", err)

		return
	}
	defer tracing.Close()

	eventRepository := repo.NewEventRepo(db)
	eventSender := sender.NewEventSender(ctx)
	retranslator := retranslator.NewRetranslator(ctx, db, cfg.Retranslator, eventRepository, eventSender)

	if err := server.NewRetranslatorServer(retranslator).Start(ctx, &cfg); err != nil {
		logger.ErrorKV(ctx, "failed creating gRPC server", "err", err)

		return
	}
}
