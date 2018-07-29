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

package orn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshal(t *testing.T) {

	t.Run("does Unmarshal the ORN successfully", func(t *testing.T) {
		expected := ORN{}
		err := Unmarshal("orn:partition:service:account:resourcetype/resource", &expected)

		assert.Nil(t, err)

		assert.Equal(t, expected.Partition, "partition")
		assert.Equal(t, expected.Service, "service")
		assert.Equal(t, expected.AccountID, "account")
		assert.Equal(t, expected.ResourceType, "resourcetype")
		assert.Equal(t, expected.Resource, "resource")
	})

	t.Run("does Unmarshal the ORN successfully with multiple / in the resource", func(t *testing.T) {
		expected := ORN{}
		err := Unmarshal("orn:acme:judge:account:user/837/bar", &expected)

		assert.Nil(t, err)

		assert.Equal(t, expected.Partition, "acme")
		assert.Equal(t, expected.Service, "judge")
		assert.Equal(t, expected.AccountID, "account")
		assert.Equal(t, expected.ResourceType, "user")
		assert.Equal(t, expected.Resource, "837/bar")
	})

	t.Run("does Unmarshal the ORN successfully without account id", func(t *testing.T) {
		expected := ORN{}
		err := Unmarshal("orn:acme:judge::user/837/bar", &expected)

		assert.Nil(t, err)

		assert.Equal(t, expected.Partition, "acme")
		assert.Equal(t, expected.Service, "judge")
		assert.Equal(t, expected.AccountID, "")
		assert.Equal(t, expected.ResourceType, "user")
		assert.Equal(t, expected.Resource, "837/bar")
	})
}

func TestMarshal(t *testing.T) {
	o := ORN{Partition: "foo", Service: "bar", AccountID: "baz", ResourceType: "biz", Resource: "fiz"}
	expected := Marshal(&o)

	assert.Equal(t, expected, "orn:foo:bar:baz:biz/fiz")
}
