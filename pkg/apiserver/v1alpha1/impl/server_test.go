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

package apiserver

import (
	"context"
	pb "github.com/gearnode/judge/pkg/apiserver/v1alpha1"
	"github.com/gearnode/judge/pkg/storage/memory"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

var (
	svc                  = NewService(memorystore.NewMemoryStore())
	invalidArgumentError = status.Error(codes.InvalidArgument, "the policy object require a non empty name")

	createPolicyRequest = pb.CreatePolicyRequest{
		Policy: &pb.Policy{
			Name:        "Demo Policy",
			Description: "A simple demo policy",
			Statements: []*pb.Statement{
				{
					Effect:    pb.Statement_ALLOW,
					Actions:   []string{"listUser", "showUser"},
					Resources: []string{"orn:judgetest:judge::user/*"},
				},
				{
					Effect:    pb.Statement_ALLOW,
					Actions:   []string{"editUser", "deleteUser"},
					Resources: []string{"orn:judgetest:judge::user/1"},
				},
			},
		},
	}
)

func TestCreatePolicy(t *testing.T) {
	var err error
	var pol *pb.Policy

	// When policy have an empty name.
	_, err = svc.CreatePolicy(context.Background(), &pb.CreatePolicyRequest{})
	assert.NotNil(t, err)
	assert.Equal(t, invalidArgumentError, err)

	// When the create policy request is valid (when no ORN was provide the name is used to generate the ORN).
	pol, err = svc.CreatePolicy(context.Background(), &createPolicyRequest)
	assert.Nil(t, err)
	assert.Equal(t, "orn:judge-org:judge-server::policy/demo-policy", pol.PolicyId)
	assert.Equal(t, "Demo Policy", pol.Name)
	assert.Equal(t, "A simple demo policy", pol.Description)
	assert.Equal(t, 2, len(pol.Statements))
	assert.Equal(t, pol.Statements[0].Effect.String(), "ALLOW")
	assert.Equal(t, pol.Statements[0].Actions, []string{"listUser", "showUser"})
	assert.Equal(t, pol.Statements[0].Resources, []string{"orn:judgetest:judge::user/*"})
	assert.Equal(t, pol.Statements[1].Effect.String(), "ALLOW")
	assert.Equal(t, pol.Statements[1].Actions, []string{"editUser", "deleteUser"})
	assert.Equal(t, pol.Statements[1].Resources, []string{"orn:judgetest:judge::user/1"})

	// When the create policy request is valid and set manualy the orn.
	createPolicyRequest.PolicyId = "orn:judgetest:judge::super-policy/super-policy-policy"
	pol, err = svc.CreatePolicy(context.Background(), &createPolicyRequest)
	assert.Nil(t, err)
	assert.Equal(t, "orn:judgetest:judge::super-policy/super-policy-policy", pol.PolicyId)

	assert.Equal(t, "Demo Policy", pol.Name)
	assert.Equal(t, "A simple demo policy", pol.Description)
	assert.Equal(t, 2, len(pol.Statements))
	assert.Equal(t, pol.Statements[0].Effect.String(), "ALLOW")
	assert.Equal(t, pol.Statements[0].Actions, []string{"listUser", "showUser"})
	assert.Equal(t, pol.Statements[0].Resources, []string{"orn:judgetest:judge::user/*"})
	assert.Equal(t, pol.Statements[1].Effect.String(), "ALLOW")
	assert.Equal(t, pol.Statements[1].Actions, []string{"editUser", "deleteUser"})
	assert.Equal(t, pol.Statements[1].Resources, []string{"orn:judgetest:judge::user/1"})
}

func TestGetPolicy(t *testing.T) {
	var err error
	var pol *pb.Policy

	_, err = svc.GetPolicy(context.Background(), &pb.GetPolicyRequest{})
	assert.NotNil(t, err)

	createPolicyRequest.PolicyId = "orn:judgetest:judge::super-policy/policy-policy"
	pol, err = svc.CreatePolicy(context.Background(), &createPolicyRequest)
	assert.Nil(t, err)
	assert.Equal(t, "orn:judgetest:judge::super-policy/policy-policy", pol.PolicyId)

	_, err = svc.GetPolicy(context.Background(), &pb.GetPolicyRequest{PolicyId: "orn:judgetest:judge::super-policy/policy-policy"})
	assert.Nil(t, err)
	assert.Equal(t, "orn:judgetest:judge::super-policy/policy-policy", pol.PolicyId)
	assert.Equal(t, "Demo Policy", pol.Name)
	assert.Equal(t, "A simple demo policy", pol.Description)
	assert.Equal(t, 2, len(pol.Statements))
	assert.Equal(t, pol.Statements[0].Effect.String(), "ALLOW")
	assert.Equal(t, pol.Statements[0].Actions, []string{"listUser", "showUser"})
	assert.Equal(t, pol.Statements[0].Resources, []string{"orn:judgetest:judge::user/*"})
	assert.Equal(t, pol.Statements[1].Effect.String(), "ALLOW")
	assert.Equal(t, pol.Statements[1].Actions, []string{"editUser", "deleteUser"})
	assert.Equal(t, pol.Statements[1].Resources, []string{"orn:judgetest:judge::user/1"})
}
