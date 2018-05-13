package judge

import (
	"fmt"
	"github.com/gearnode/judge/api"
	"github.com/gearnode/judge/orn"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct{}

func Start(addr string, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		return err
	}

	s := Server{}
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	api.RegisterJudgeServer(grpcServer, &s)

	return grpcServer.Serve(lis)
}

func (s *Server) Authorize(ctx context.Context, in *api.AuthorizeRequest) (*api.AuthorizeResponse, error) {
	log.Printf("Receive Authorize Request for execute %s on %s", in.Action, in.Orn)
	ok, err := Authorize([]Policy{}, orn.ORN{}, "")
	return &api.AuthorizeResponse{Authorize: ok}, err
}
