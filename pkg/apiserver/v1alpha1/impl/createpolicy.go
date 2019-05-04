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
	pb "github.com/gearnode/judge/pkg/apiserver/v1alpha1"
	"github.com/gearnode/judge/pkg/orn"
	"github.com/gearnode/judge/pkg/policy"
	"github.com/gearnode/judge/pkg/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreatePolicy create a policy. This function should not allow upsert and returns error when the
// policy already exist. This function allow End-User to override the default auto generate policy ID.
func (s *Service) CreatePolicy(ctx context.Context, in *pb.CreatePolicyRequest) (*pb.Policy, error) {
	pol, err := policy.NewPolicy(
		in.GetPolicy().GetName(),
		in.GetPolicy().GetDescription(),
	)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if in.GetPolicyId() != "" {
		err := orn.Unmarshal(in.GetPolicyId(), &pol.ID)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	for _, statement := range in.GetPolicy().GetStatements() {
		stmt, err := policy.NewStatement(
			statement.GetEffect().String(),
			statement.GetActions(),
			statement.GetResources(),
		)

		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		pol.Statements = append(pol.Statements, *stmt)
	}

	if policyExist(s.state, pol.ID.String()) {
		return nil, status.Error(codes.AlreadyExists, "the policy already exists")
	}

	pol, err = s.state.PutPolicy(pol)
	if err != nil {
		// TODO: log put error
		return nil, status.Error(codes.Internal, "internal error")
	}

	return toProto(pol), nil
}

func policyExist(store storage.Storage, id string) bool {
	_, err := store.GetPolicy(id)
	if err != nil {
		return false
	}
	return true
}
