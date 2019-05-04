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
	"context"
	"fmt"
	pb "github.com/gearnode/judge/pkg/apiserver/v1alpha1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DeletePolicy delete the policy with the given policy ID from the state. This function can or not can returns
// error when the given policy ID does not exist it's depend of the storage implementation.
func (s *Service) DeletePolicy(ctx context.Context, in *pb.DeletePolicyRequest) (*empty.Empty, error) {
	err := s.state.DelPolicy(in.GetPolicyId())
	if err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("the policy with the orn %q is not found", in.GetPolicyId()))
	}

	return &empty.Empty{}, nil
}
