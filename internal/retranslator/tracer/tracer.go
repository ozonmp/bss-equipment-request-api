package tracer

import (
	"context"
	"github.com/ozonmp/bss-equipment-request-api/internal/retranslator/logger"
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"

	"github.com/ozonmp/bss-equipment-request-api/internal/retranslator/config"

	jaegercfg "github.com/uber/jaeger-client-go/config"
)

// NewTracer - returns new tracer.
func NewTracer(ctx context.Context, cfg *config.Config) (io.Closer, error) {
	cfgTracer := &jaegercfg.Configuration{
		ServiceName: cfg.Jaeger.Service,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: cfg.Jaeger.Host + cfg.Jaeger.Port,
		},
	}
	tracer, closer, err := cfgTracer.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		logger.ErrorKV(ctx, "cfg.NewTracer()", "err", err)

		return nil, err
	}
	opentracing.SetGlobalTracer(tracer)

	logger.InfoKV(ctx, "cfg.NewTracer()", ": Traces started")

	return closer, nil
}