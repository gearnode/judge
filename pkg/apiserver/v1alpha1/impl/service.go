/*
Copyright 2019 Bryan Frimin.

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

package apiserver

import (
	pb "github.com/gearnode/judge/pkg/apiserver/v1alpha1"
	"github.com/gearnode/judge/pkg/storage"
	"google.golang.org/grpc"
)

// Service contains dependencies used by the service and the service implementation.
type Service struct {
	state storage.Storage
}

// NewService create a new service with a default configuration.
func NewService(store storage.Storage) *Service {
	return &Service{state: store}
}

// Register register the service to an gRPC server.
func (s *Service) Register(server *grpc.Server) {
	pb.RegisterJudgeServer(server, s)
}
