package interceptor

import (
	grpcZap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func ZapLoggerInterceptor() grpc.UnaryServerInterceptor {
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
	return grpcZap.UnaryServerInterceptor(zapPrd, opt)

}
