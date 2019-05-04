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

package policy

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPolicy(t *testing.T) {
	assert := assert.New(t)

	pol, err := NewPolicy("some-name", "some description of the policy")
	assert.NoError(err)
	assert.Equal("some-name", pol.Name)
	assert.Equal("some description of the policy", pol.Description)
	assert.Equal("judge-org", pol.ID.Partition)
	assert.Equal("judge-server", pol.ID.Service)
	assert.Equal("", pol.ID.AccountID)
	assert.Equal("policy", pol.ID.ResourceType)
	assert.Equal("some-name", pol.ID.Resource)

	pol, err = NewPolicy("", "some description of the policy")
	assert.Error(err)
	assert.Equal(err, errors.New("the policy object require a non empty name"))
	assert.Nil(pol)

	pol, err = NewPolicy("some-name", "")
	assert.NoError(err)
	assert.Equal("some-name", pol.Name)
	assert.Equal("", pol.Description)
	assert.Equal("judge-org", pol.ID.Partition)
	assert.Equal("judge-server", pol.ID.Service)
	assert.Equal("", pol.ID.AccountID)
	assert.Equal("policy", pol.ID.ResourceType)
	assert.Equal("some-name", pol.ID.Resource)
}

func TestNewStatement(t *testing.T) {
	assert := assert.New(t)
	actions := []string{"edit", "show"}
	resources := []string{"orn:judge-org:policy-service::foo/bar"}
	effect := "ALLOW"

	t.Run("When the effect is not allowed", func(t *testing.T) {
		stmt, err := NewStatement("NotAllowed", actions, resources)
		assert.Error(err)
		assert.Nil(stmt)

		stmt, err = NewStatement("Allowed", actions, resources)
		assert.Error(err)
		assert.Nil(stmt)
	})

	t.Run("When the effect is allowed", func(t *testing.T) {
		stmt, err := NewStatement("ALLOW", actions, resources)
		assert.NoError(err)
		assert.Equal(stmt.Effect, "ALLOW")

		stmt, err = NewStatement("DENY", actions, resources)
		assert.NoError(err)
		assert.Equal(stmt.Effect, "DENY")
	})

	t.Run("When resources is empty", func(t *testing.T) {
		stmt, err := NewStatement(effect, actions, []string{})
		assert.Error(err)
		assert.Nil(stmt)
	})

	t.Run("When actions is empty", func(t *testing.T) {
		stmt, err := NewStatement(effect, []string{}, resources)
		assert.Error(err)
		assert.Nil(stmt)
	})

	t.Run("When actions contain at least one empty action", func(t *testing.T) {
		stmt, err := NewStatement(effect, []string{"edit", "", "show"}, resources)
		assert.Error(err)
		assert.Nil(stmt)

		stmt, err = NewStatement(effect, []string{"edit", "", "", ""}, resources)
		assert.Error(err)
		assert.Nil(stmt)
	})

	t.Run("When at least one resource is invalid", func(t *testing.T) {
		stmt, err := NewStatement(effect, actions, []string{"foo"})
		assert.Error(err)
		assert.Nil(stmt)

		stmt, err = NewStatement(effect, actions, []string{"orn:judge-org:policy-service::foo/hello", "foo", "orn:judge-org:policy-service::foo/bar"})
		assert.Error(err)
		assert.Nil(stmt)

		stmt, err = NewStatement(effect, actions, []string{"", "orn:judge-org:policy-service::foo/bar"})
		assert.Error(err)
		assert.Nil(stmt)
	})

	stmt, err := NewStatement(effect, actions, resources)
	assert.NoError(err)
	assert.Equal(stmt.Actions, actions)
	assert.Equal(len(stmt.Resources), 1)
	assert.Equal(stmt.Effect, "ALLOW")
	assert.Equal(stmt.Resources[0].Partition, "judge-org")
	assert.Equal(stmt.Resources[0].Service, "policy-service")
}
