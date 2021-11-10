package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/ozonmp/bss-equipment-request-api/internal/app/repo"
	"github.com/ozonmp/bss-equipment-request-api/internal/app/retranslator"
	"github.com/ozonmp/bss-equipment-request-api/internal/config"
	"github.com/ozonmp/bss-equipment-request-api/internal/database"
	"github.com/ozonmp/bss-equipment-request-api/internal/tracer"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}
	cfg := config.GetConfigInstance()

	migration := flag.Bool("migration", true, "Defines the migration start option")
	flag.Parse()

	log.Info().
		Str("version", cfg.Project.Version).
		Str("commitHash", cfg.Project.CommitHash).
		Bool("debug", cfg.Project.Debug).
		Str("environment", cfg.Project.Environment).
		Msgf("Starting service: %s", cfg.Project.Name)

	// default: zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if cfg.Project.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SslMode,
	)

	db, err := database.NewPostgres(dsn, cfg.Database.Driver)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed init postgres")
	}
	defer db.Close()

	if *migration {
		if err = goose.Up(db.DB, cfg.Database.Migrations); err != nil {
			log.Error().Err(err).Msg("Migration failed")

			return
		}
	}

	tracing, err := tracer.NewTracer(&cfg)
	if err != nil {
		log.Error().Err(err).Msg("Failed init tracing")

		return
	}
	defer tracing.Close()

	sigs := make(chan os.Signal, 1)

	parentCtx := context.Background()
	ctx, cancel := context.WithCancel(parentCtx)

	eventRepository := repo.NewEventRepo(db)

	rcfg := retranslator.Config{
		DB:              db,
		ChannelSize:     512,
		ConsumerCount:   2,
		ConsumeTimeout:  10 * time.Second,
		ProducerTimeout: 10 * time.Second,
		ConsumeSize:     10,
		ProducerCount:   28,
		WorkerCount:     2,
		Ctx:             ctx,
		CancelCtxFunc:   cancel,
		Repo:            eventRepository,
		BatchSize:       20,
	}

	retranslator := retranslator.NewRetranslator(rcfg)
	retranslator.Start()

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
}
