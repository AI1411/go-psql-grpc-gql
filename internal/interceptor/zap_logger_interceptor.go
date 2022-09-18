package interceptor

import (
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcZap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpcCtxTags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func ZapLoggerInterceptor() grpc.ServerOption {
	zapPrd, _ := zap.NewProduction()
	opt := grpcZap.WithLevels(
		func(c codes.Code) zapcore.Level {
			var l zapcore.Level
			switch c {
			case codes.OK:
				l = zapcore.InfoLevel
			case codes.Internal:
				l = zapcore.ErrorLevel
			default:
				l = zapcore.WarnLevel
			}
			return l
		})
	return grpcMiddleware.WithUnaryServerChain(
		grpcCtxTags.UnaryServerInterceptor(),

		grpcZap.UnaryServerInterceptor(zapPrd, opt),
	)
}
