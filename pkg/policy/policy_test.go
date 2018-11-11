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

package policy

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPolicy(t *testing.T) {
	pol, err := NewPolicy("some-name", "some description of the policy")
	assert.Nil(t, err)
	assert.Equal(t, "some-name", pol.Name)
	assert.Equal(t, "some description of the policy", pol.Description)
	assert.Equal(t, STANDALONE, pol.Type)
	assert.Equal(t, VERSION, pol.Document.Version)
	assert.Equal(t, "judge-org", pol.ORN.Partition)
	assert.Equal(t, "judge-server", pol.ORN.Service)
	assert.Equal(t, "", pol.ORN.AccountID)
	assert.Equal(t, "policy", pol.ORN.ResourceType)
	assert.Equal(t, "some-name", pol.ORN.Resource)

	pol, err = NewPolicy("", "some description of the policy")
	assert.NotNil(t, err)
	assert.Equal(t, err, errors.New("the policy object require a non empty name"))
	assert.Equal(t, pol, &Policy{})

	pol, err = NewPolicy("some-name", "")
	assert.Nil(t, err)
	assert.Equal(t, "some-name", pol.Name)
	assert.Equal(t, "", pol.Description)
	assert.Equal(t, STANDALONE, pol.Type)
	assert.Equal(t, VERSION, pol.Document.Version)
	assert.Equal(t, "judge-org", pol.ORN.Partition)
	assert.Equal(t, "judge-server", pol.ORN.Service)
	assert.Equal(t, "", pol.ORN.AccountID)
	assert.Equal(t, "policy", pol.ORN.ResourceType)
	assert.Equal(t, "some-name", pol.ORN.Resource)
}

func TestNewStatement(t *testing.T) {
	actions := []string{"edit", "show"}
	resources := []string{"orn:judge-org:policy-service::foo/bar"}
	effect := "Allow"

	t.Run("When the effect is not allowed", func(t *testing.T) {
		stmt, err := NewStatement("NotAllowed", actions, resources)
		assert.NotNil(t, err)
		assert.Equal(t, stmt, &Statement{})

		stmt, err = NewStatement("Allowed", actions, resources)
		assert.NotNil(t, err)
		assert.Equal(t, stmt, &Statement{})
	})

	t.Run("When the effect is allowed", func(t *testing.T) {
		stmt, err := NewStatement("Allow", actions, resources)
		assert.Nil(t, err)
		assert.Equal(t, stmt.Effect, "Allow")

		stmt, err = NewStatement("Deny", actions, resources)
		assert.Nil(t, err)
		assert.Equal(t, stmt.Effect, "Deny")
	})

	t.Run("When resources is empty", func(t *testing.T) {
		stmt, err := NewStatement(effect, actions, []string{})
		assert.NotNil(t, err)
		assert.Equal(t, stmt, &Statement{})
	})

	t.Run("When actions is empty", func(t *testing.T) {
		stmt, err := NewStatement(effect, []string{}, resources)
		assert.NotNil(t, err)
		assert.Equal(t, stmt, &Statement{})
	})

	t.Run("When actions contain at least one empty action", func(t *testing.T) {
		stmt, err := NewStatement(effect, []string{"edit", "", "show"}, resources)
		assert.NotNil(t, err)
		assert.Equal(t, stmt, &Statement{})

		stmt, err = NewStatement(effect, []string{"edit", "", "", ""}, resources)
		assert.NotNil(t, err)
		assert.Equal(t, stmt, &Statement{})
	})

	t.Run("When at least one resource is invalid", func(t *testing.T) {
		stmt, err := NewStatement(effect, actions, []string{"foo"})
		assert.NotNil(t, err)
		assert.Equal(t, stmt, &Statement{})

		stmt, err = NewStatement(effect, actions, []string{"orn:judge-org:policy-service::foo/hello", "foo", "orn:judge-org:policy-service::foo/bar"})
		assert.NotNil(t, err)
		assert.Equal(t, stmt, &Statement{})

		stmt, err = NewStatement(effect, actions, []string{"", "orn:judge-org:policy-service::foo/bar"})
		assert.NotNil(t, err)
		assert.Equal(t, stmt, &Statement{})
	})

	stmt, err := NewStatement(effect, actions, resources)
	assert.Nil(t, err)
	assert.Equal(t, stmt.Action, actions)
	assert.Equal(t, len(stmt.Resource), 1)
	assert.Equal(t, stmt.Effect, "Allow")
	assert.Equal(t, stmt.Resource[0].Partition, "judge-org")
	assert.Equal(t, stmt.Resource[0].Service, "policy-service")
}
