package dogs

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"pupa/servers"
	srv "pupa/servers/grpc/pupa"

	"google.golang.org/grpc"
)

var (
	_ srv.DogServiceServer = (*Server)(nil)
	_ servers.GRPCService  = (*Server)(nil)
)

type (
	Server struct {
		srv.UnsafeDogServiceServer

		repo *Repository
	}
)

func NewServer(repo *Repository) *Server {
	return &Server{
		repo: repo,
	}
}

func (s *Server) Register(r grpc.ServiceRegistrar) {
	r.RegisterService(&srv.DogService_ServiceDesc, s)
}

func (s *Server) DogIsGoodBoyV1(ctx context.Context, req *srv.DogIsGoodBoyV1Request) (*srv.DogIsGoodBoyV1Response, error) {
	dog, err := s.repo.DogByName(ctx, req.GetName())
	if err != nil {
		if errors.Is(err, ErrorNotFound) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &srv.DogIsGoodBoyV1Response{
		IsGoodBoy: dog.GoodBoy,
	}, nil
}
