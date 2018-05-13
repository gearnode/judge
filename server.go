package judge

import (
	"crypto/tls"
	"fmt"
	"github.com/gearnode/judge/orn"
	"github.com/gearnode/judge/svc"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

// Server represent a server instance. This struct store the
// server configuration.
type Server struct {
	Port int
	Addr string
	Cert *tls.Certificate
}

// Start a new server
func (s *Server) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Addr, s.Port))
	if err != nil {
		return err
	}

	creds := credentials.NewServerTLSFromCert(s.Cert)
	opts := []grpc.ServerOption{grpc.Creds(creds)}

	grpcServer := grpc.NewServer(opts...)
	svc.RegisterJudgeServer(grpcServer, s)

	return grpcServer.Serve(lis)
}

// Authorize todo
func (s *Server) Authorize(ctx context.Context, in *svc.AuthorizeRequest) (*svc.AuthorizeResponse, error) {
	log.Printf("Receive Authorize Request for execute %s on %s", in.Action, in.Orn)
	ok, err := Authorize([]Policy{}, orn.ORN{}, "")
	return &svc.AuthorizeResponse{Authorize: ok}, err
}
