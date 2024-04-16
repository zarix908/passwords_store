package main

import (
	"context"
	"os"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/zarix908/passwords_store/cmd/grpcserver"
)

func main() {
	rootCmd := &cobra.Command{
		Use:         "passheimdal",
		Short:       "Passheimdal - simple passwords manager based on age and sops.",
		Example:     "passheimdal <command>",
		ValidArgs:   []string{"<coomand>"},
		Annotations: map[string]string{cobra.BashCompCustom: "rungrpc"},
	}

	rootCmd.AddCommand(grpcserver.NewRunGrpcServerCommand())

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout})
	ctx := logger.WithContext(context.Background())

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		logger.Err(err).Msg("Command execution failed.")
		os.Exit(1)
	}
}
