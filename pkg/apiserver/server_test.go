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
	"github.com/gearnode/judge/pkg/apiserver/v1alpha1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

var (
	srv                  = Server{}
	invalidArgumentError = status.Error(codes.InvalidArgument, "the policy object require a non empty name")

	createPolicyRequest = v1alpha1.CreatePolicyRequest{
		Policy: &v1alpha1.Policy{
			Name:        "Demo Policy",
			Description: "A simple demo policy",
			Document: &v1alpha1.Document{
				Version: "v1alpha1",
				Statements: []*v1alpha1.Statement{
					&v1alpha1.Statement{
						Effect:    "Allow",
						Actions:   []string{"listUser", "showUser"},
						Resources: []string{"orn:judgetest:judge::user/*"},
					},
					&v1alpha1.Statement{
						Effect:    "Allow",
						Actions:   []string{"editUser", "deleteUser"},
						Resources: []string{"orn:judgetest:judge::user/1"},
					},
				},
			},
		},
	}
)

func TestCreatePolicy(t *testing.T) {
	var err error
	var pol *v1alpha1.Policy

	// When policy have an empty name.
	_, err = srv.CreatePolicy(context.Background(), &v1alpha1.CreatePolicyRequest{})
	assert.NotNil(t, err)
	assert.Equal(t, invalidArgumentError, err)

	// When the create policy request is valid (when no ORN was provide the name is used to generate the ORN).
	pol, err = srv.CreatePolicy(context.Background(), &createPolicyRequest)
	assert.Nil(t, err)
	assert.Equal(t, "orn:judge-org:judge-server::policy/demo-policy", pol.Orn)
	assert.Equal(t, "Demo Policy", pol.Name)
	assert.Equal(t, "A simple demo policy", pol.Description)
	assert.Equal(t, "v1alpha1", pol.Document.Version)
	assert.Equal(t, 2, len(pol.Document.Statements))
	assert.Equal(t, pol.Document.Statements[0].Effect, "Allow")
	assert.Equal(t, pol.Document.Statements[0].Actions, []string{"listUser", "showUser"})
	assert.Equal(t, pol.Document.Statements[0].Resources, []string{"orn:judgetest:judge::user/*"})
	assert.Equal(t, pol.Document.Statements[1].Effect, "Allow")
	assert.Equal(t, pol.Document.Statements[1].Actions, []string{"editUser", "deleteUser"})
	assert.Equal(t, pol.Document.Statements[1].Resources, []string{"orn:judgetest:judge::user/1"})

	// When the create policy request is valid and set manualy the orn.
	createPolicyRequest.Orn = "orn:judgetest:judge::super-policy/super-policy-policy"
	pol, err = srv.CreatePolicy(context.Background(), &createPolicyRequest)
	assert.Nil(t, err)
	assert.Equal(t, "orn:judgetest:judge::super-policy/super-policy-policy", pol.Orn)

	assert.Equal(t, "Demo Policy", pol.Name)
	assert.Equal(t, "A simple demo policy", pol.Description)
	assert.Equal(t, "v1alpha1", pol.Document.Version)
	assert.Equal(t, 2, len(pol.Document.Statements))
	assert.Equal(t, pol.Document.Statements[0].Effect, "Allow")
	assert.Equal(t, pol.Document.Statements[0].Actions, []string{"listUser", "showUser"})
	assert.Equal(t, pol.Document.Statements[0].Resources, []string{"orn:judgetest:judge::user/*"})
	assert.Equal(t, pol.Document.Statements[1].Effect, "Allow")
	assert.Equal(t, pol.Document.Statements[1].Actions, []string{"editUser", "deleteUser"})
	assert.Equal(t, pol.Document.Statements[1].Resources, []string{"orn:judgetest:judge::user/1"})
}
