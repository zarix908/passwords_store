package xlog

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func UnaryServerInterceptor(logger zerolog.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context, req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		ctx = logger.WithContext(ctx)
		logger.Info().Str("method", info.FullMethod).Msg("request")

		resp, err := handler(ctx, req)
		if err != nil {
			code := status.Code(err)
			logger.Err(err).Fields(map[string]interface{}{
				"method": info.FullMethod,
				"code":   code,
			}).Msg("request failed")

			if code == codes.Unknown {
				return nil, fmt.Errorf("uknown error")
			}

			return nil, err
		}

		return resp, nil
	}
}
