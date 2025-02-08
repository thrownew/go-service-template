package servers

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type (
	GRPCServer struct {
		grpc   *grpc.Server
		health *health.Server
	}

	GRPCService interface {
		Register(grpc.ServiceRegistrar)
	}
)

func NewServer(services ...GRPCService) *GRPCServer {
	s := grpc.NewServer()

	reflection.Register(s)

	hs := health.NewServer()
	hs.SetServingStatus("", grpc_health_v1.HealthCheckResponse_NOT_SERVING)

	s.RegisterService(&grpc_health_v1.Health_ServiceDesc, hs)

	// Register services
	for _, srv := range services {
		srv.Register(s)
	}

	return &GRPCServer{
		grpc:   s,
		health: hs,
	}
}

func (s *GRPCServer) Ready(ready bool) {
	if ready {
		s.health.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)
	} else {
		s.health.SetServingStatus("", grpc_health_v1.HealthCheckResponse_NOT_SERVING)
	}
}

func (s *GRPCServer) Start(addr string) error {
	ls, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("server: grpc: listen `%s`: %w", addr, err)
	}
	go func() {
		if sErr := s.grpc.Serve(ls); sErr != nil && !errors.Is(sErr, grpc.ErrServerStopped) {
			panic(fmt.Errorf("server: grpc: serve `%s`: %w", addr, sErr))
		}
	}()
	slog.Info(fmt.Sprintf("Server: grpc: Listen and Serve `%s`", addr))
	return nil
}

func (s *GRPCServer) Stop(_ context.Context) {
	s.grpc.GracefulStop()
	slog.Info("Server: grpc: Shutdown")
}
