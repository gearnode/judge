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
	"github.com/gearnode/judge/api/judge/v1alpha1"
	"github.com/gearnode/judge/pkg/storage/memory"
	"github.com/golang/protobuf/ptypes/empty"
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
	v1alpha1.RegisterJudgeServer(grpcServer, s)

	return grpcServer.Serve(lis)
}

var (
	store = memorystore.NewMemoryStore()
)

// Authorize todo
func (s *Server) Authorize(ctx context.Context, in *v1alpha1.AuthorizeRequest) (*v1alpha1.AuthorizeResponse, error) {
	log.Printf("Receive Authorize Request to execute")
	return &v1alpha1.AuthorizeResponse{}, nil
}

func (s *Server) GetPolicy(ctx context.Context, in *v1alpha1.GetPolicyRequest) (*v1alpha1.Policy, error) {
	log.Printf("Receive GetPolicy Request to execute")
	return &v1alpha1.Policy{}, nil
}

func (s *Server) Listpolicies(ctx context.Context, in *v1alpha1.ListPoliciesRequest) (*v1alpha1.ListPoliciesResponse, error) {
	log.Printf("Receive ListPoliciesRequest Request to execute")
	return &v1alpha1.ListPoliciesResponse{}, nil
}

func (s *Server) CreatePolicy(ctx context.Context, in *v1alpha1.CreatePolicyRequest) (*v1alpha1.Policy, error) {
	log.Printf("Receive CreatePolicy Request to execute")
	return &v1alpha1.Policy{}, nil
}

func (s *Server) UpdatePolicy(ctx context.Context, in *v1alpha1.UpdatePolicyRequest) (*v1alpha1.Policy, error) {
	log.Printf("Receive UpdateDocument Resquest to execute")
	return &v1alpha1.Policy{}, nil
}

func (s *Server) DeletePolicy(ctx context.Context, in *v1alpha1.DeletePolicyRequest) (*empty.Empty, error) {
	log.Printf("Receive DeletePolicy Resquest to execute")
	return &empty.Empty{}, nil
}
