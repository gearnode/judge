/*
Copyright 2018 The Judge Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package judge

import (
	"crypto/tls"
	"fmt"
	"github.com/gearnode/judge/pkg/apiserver/svc"
	"github.com/gearnode/judge/pkg/orn"
	"github.com/gearnode/judge/pkg/policy"
	"github.com/gearnode/judge/pkg/storage/memory"
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

var (
	store = memorystore.NewMemoryStore()
)

// Authorize todo
func (s *Server) Authorize(ctx context.Context, in *svc.AuthorizeRequest) (*svc.AuthorizeResponse, error) {
	log.Printf("Receive Authorize Request for execute %s on %s", in.GetAction(), in.GetOrn())
	o := orn.ORN{}
	orn.Unmarshal(in.GetOrn(), &o)
	ok, err := judge.Authorize(store, o, in.GetAction())
	return &svc.AuthorizeResponse{Authorize: ok}, err
}

// CreatePolicy todo
func (s *Server) CreatePolicy(ctx context.Context, in *svc.CreatePolicyRequest) (*svc.CreatePolicyResponse, error) {
	log.Printf("Receive CreatePolicy Request to execute")
	ok, err := judge.CreatePolicy(store, in.GetName(), in.GetDescription(), in.GetDocument())
	return &svc.CreatePolicyResponse{Ok: ok, Error: err.Error()}, err
}
