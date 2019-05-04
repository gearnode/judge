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
	"github.com/gearnode/judge/pkg/authorize"
	"github.com/gearnode/judge/pkg/orn"
	"github.com/gearnode/judge/pkg/policy"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

// Authorize decides whether a given authorization request should be allowed or denied.
func (s *Service) Authorize(ctx context.Context, in *pb.AuthorizeRequest) (*pb.AuthorizeResponse, error) {
	var object orn.ORN
	err := orn.Unmarshal(in.GetObject(), &object)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// TODO: @gearnode set execution subject with the current authenticated user
	in.Context["execution.subject"] = "some subject id"
	in.Context["execution.time"] = fmt.Sprintf("%d", time.Now().Unix())

	// TODO: @gearnode populates policies with stored policies
	policies := make([]*policy.Policy, 0)

	err = authorize.Authorize(policies, in.GetAction(), object, in.GetContext())
	if err != nil {
		return &pb.AuthorizeResponse{Authorized: false, Explain: err.Error()}, nil
	}

	return &pb.AuthorizeResponse{Authorized: true}, nil
}
