package grps_logger

import (
	"context"
	"errors"
	"github.com/ozonmp/bss-equipment-request-api/internal/logger"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware/logging"

	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strings"
)

// UnaryServerInterceptor - get a new log level from meta and set it
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			logger.ErrorKV(ctx, "got log level",
				"error", "couldn't parse incoming context metadata",
			)

			return nil, errors.New("couldn't parse incoming context metadata")
		}

		levels := md.Get("log-level")
		logger.InfoKV(ctx, "got log level", "levels", levels)

		if len(levels) > 0 {
			if parsedLevel, ok := parseLevel(levels[0]); ok {
				newLogger := logger.CloneWithLevel(ctx, parsedLevel)
				ctx = logger.AttachLogger(ctx, newLogger)
			}
		}

		h, err := handler(ctx, req)
		return h, err
	}
}

// ServerPayloadLoggingDecider - check should we log response and request data or not
func ServerPayloadLoggingDecider() grpc_middleware.ServerPayloadLoggingDecider {
	return func(ctx context.Context, fullMethodName string, servingObject interface{}) bool {
		return logger.IsEnabled(ctx, zapcore.DebugLevel)
	}
}

func parseLevel(str string) (zapcore.Level, bool) {
	switch strings.ToLower(str) {
	case "debug":
		return zapcore.DebugLevel, true
	case "info":
		return zapcore.InfoLevel, true
	case "warn":
		return zapcore.WarnLevel, true
	case "error":
		return zapcore.ErrorLevel, true
	default:
		return zapcore.DebugLevel, false
	}
}
