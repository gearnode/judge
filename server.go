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

type PolicyStore interface {
	GetAll() []Policy
	Put(Policy) bool
	Flush() bool
}

type MemoryStore struct {
	policies []Policy
	PolicyStore
}

func (s *MemoryStore) GetAll() []Policy {
	return s.policies
}

func (s *MemoryStore) Put(policy Policy) bool {
	s.policies = append(s.policies, policy)
	return true
}

func (s *MemoryStore) Flush() bool {
	s.policies = []Policy{}
	return true
}

var (
	store = MemoryStore{}
)

// Authorize todo
func (s *Server) Authorize(ctx context.Context, in *svc.AuthorizeRequest) (*svc.AuthorizeResponse, error) {
	log.Printf("Receive Authorize Request for execute %s on %s", in.GetAction(), in.GetOrn())
	o := orn.ORN{}
	orn.Unmarshal(in.GetOrn(), &o)
	ok, err := Authorize(&store, o, in.GetAction())
	return &svc.AuthorizeResponse{Authorize: ok}, err
}

func (s *Server) CreatePolicy(ctx context.Context, in *svc.CreatePolicyRequest) (*svc.CreatePolicyResponse, error) {
	log.Printf("Receive CreatePolicy Request to execute")
	ok, err := CreatePolicy(&store, in.GetName(), in.GetDescription(), in.GetDocument())
	return &svc.CreatePolicyResponse{Ok: ok, Error: err.Error()}, err
}
