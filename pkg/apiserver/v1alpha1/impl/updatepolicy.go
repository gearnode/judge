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
	"github.com/gearnode/judge/pkg/policy"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UpdatePolicy upsert a policy. This function create the policy when the ID does not exist or update
// the policy with the same ID. If you need creation only use the CreatePolicy function. This function
// allow End-User to override the default auto generate policy ID.
func (s *Service) UpdatePolicy(ctx context.Context, in *pb.UpdatePolicyRequest) (*pb.Policy, error) {
	pol, err := policy.NewPolicy(
		in.GetPolicy().GetName(),
		in.GetPolicy().GetDescription(),
	)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
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

	pol, err = s.state.PutPolicy(pol)
	if err != nil {
		// TODO: log put error
		return nil, status.Error(codes.Internal, "internal error")
	}

	return toProto(pol), nil
}
