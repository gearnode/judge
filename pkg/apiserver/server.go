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

package apiserver // import "github.com/gearnode/judge/pkg/apiserver"

import (
	"fmt"
	"net"

	"github.com/gearnode/judge/api/judge/v1alpha1"
	"github.com/gearnode/judge/pkg/storage/memory"
	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

// Server represent a server instance. This struct store the
// server configuration.
type Server struct {
	Port  int
	Addr  string
	Creds *credentials.TransportCredentials
}

// Start a new server
func (s *Server) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Addr, s.Port))
	if err != nil {
		return err
	}

	var grpcServer *grpc.Server
	if s.Creds != nil {
		grpcServer = grpc.NewServer(grpc.Creds(*s.Creds))
	} else {
		grpcServer = grpc.NewServer()
	}

	reflection.Register(grpcServer)
	v1alpha1.RegisterJudgeServer(grpcServer, s)

	return grpcServer.Serve(lis)
}

var (
	store = memorystore.NewMemoryStore()
)

// Authorize implement judge.api.v1beta1.Judge.Authorize
func (s *Server) Authorize(ctx context.Context, in *v1alpha1.AuthorizeRequest) (*v1alpha1.AuthorizeResponse, error) {
	log.Info("Receive Authorize Request to execute")
	return &v1alpha1.AuthorizeResponse{}, nil
}

// GetPolicy implement judge.api.v1beta1.Judge.GetPolicy
func (s *Server) GetPolicy(ctx context.Context, in *v1alpha1.GetPolicyRequest) (*v1alpha1.Policy, error) {
	log.Info("Receive GetPolicy Request to execute")
	return &v1alpha1.Policy{}, nil
}

// ListPolicies implement judge.api.v1beta1.Judge.ListPolicies
func (s *Server) ListPolicies(ctx context.Context, in *v1alpha1.ListPoliciesRequest) (*v1alpha1.ListPoliciesResponse, error) {
	log.Info("Receive ListPoliciesRequest Request to execute")
	return &v1alpha1.ListPoliciesResponse{}, nil
}

// CreatePolicy implement judge.api.v1beta1.Judge.CreatePolicy
func (s *Server) CreatePolicy(ctx context.Context, in *v1alpha1.CreatePolicyRequest) (*v1alpha1.Policy, error) {
	log.Info("Receive CreatePolicy Request to execute")
	return &v1alpha1.Policy{}, nil
}

// UpdatePolicy implement judge.api.v1beta1.Judge.UpdatePolicy
func (s *Server) UpdatePolicy(ctx context.Context, in *v1alpha1.UpdatePolicyRequest) (*v1alpha1.Policy, error) {
	log.Info("Receive UpdatePolicy Resquest to execute")
	return &v1alpha1.Policy{}, nil
}

// DeletePolicy implement judge.api.v1beta1.Judge.DeletePolicy
func (s *Server) DeletePolicy(ctx context.Context, in *v1alpha1.DeletePolicyRequest) (*empty.Empty, error) {
	log.Printf("Receive DeletePolicy Resquest to execute")
	return &empty.Empty{}, nil
}
