package grpcserver

import (
	"errors"
	"fmt"
	"net"
	"os"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	"github.com/zarix908/passwords_store/pkg/key"
	"github.com/zarix908/passwords_store/pkg/keymanager"
	keymanagerpb "github.com/zarix908/passwords_store/pkg/pb/keymanager/v1"
	"github.com/zarix908/passwords_store/pkg/storage"
	"github.com/zarix908/passwords_store/pkg/xlog"
)

func NewRunGrpcServerCommand() *cobra.Command {
	return &cobra.Command{
		Use:       "rungrpc <address>",
		Short:     "Run grpc server at specified <address>",
		ValidArgs: []string{"<address>"},
		Example:   "passheimdal rungrpc <address>",
		RunE:      runGrpcServer,
	}
}

func runGrpcServer(_ *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("invalid argument count")
	}

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout})

	s := grpc.NewServer(grpc.ChainUnaryInterceptor(
		xlog.UnaryServerInterceptor(logger),
	))

	keymanagerpb.RegisterKeyServiceServer(s, keymanager.NewKeyServiceServer(storage.New, key.NewKeyGenerator()))

	addr := args[0]
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen server: %w", err)
	}
	logger.Info().Str("address", addr).Msg("Listening...")

	if err := s.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}
