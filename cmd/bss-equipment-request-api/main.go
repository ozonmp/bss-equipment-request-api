package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ozonmp/bss-equipment-request-api/internal/app/retranslator"
)

func main() {

	sigs := make(chan os.Signal, 1)

	parentCtx := context.Background()
	ctx, cancel := context.WithCancel(parentCtx)

	cfg := retranslator.Config{
		ChannelSize:    512,
		ConsumerCount:  2,
		ConsumeTimeout: 10 * time.Second,
		ConsumeSize:    10,
		ProducerCount:  28,
		WorkerCount:    2,
		Ctx:            ctx,
		CancelCtxFunc:  cancel,
	}

	retranslator := retranslator.NewRetranslator(cfg)
	retranslator.Start()

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
}
