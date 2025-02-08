package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"pupa/app"
	"pupa/cmd/wof"

	"pupa/databases"
	"pupa/deps"
	"pupa/logs"
	"pupa/servers"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:               "pupa",
		Short:             "pupa",
		Version:           "v1.0.0",
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, _ []string) error {
			c, err := deps.NewContainer(
				deps.Provide(func() context.Context {
					return gracefulContext(cmd.Context())
				}),
				logs.Provide(),
				databases.Provide(),
				servers.Provide(),
				app.Provide(),
			)
			if err != nil {
				return fmt.Errorf("container: %w", err)
			}
			return c.Invoke(run)
		},
	}
	cmd.AddCommand(wof.NewCommand())
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}

func run(ctx context.Context, s *servers.GRPCServer) error {
	if err := s.Start(":8080"); err != nil {
		return fmt.Errorf("server: start: %w", err)
	}

	s.Ready(true)

	slog.Info("Service ready 🚀")
	<-ctx.Done()
	slog.Info("Buy buy! 👋")

	s.Ready(false)

	sCtx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()
	s.Stop(sCtx)

	return nil
}

func gracefulContext(ctx context.Context) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		<-sigs
		cancel()
	}()
	return ctx
}
