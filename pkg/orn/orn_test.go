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

package orn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func DoesUnmarshalORNWithSuccess(t *testing.T) {
	assert := assert.New(t)
	expected := ORN{}
	err := Unmarshal("orn:partition:service:account:resourcetype/resource", &expected)

	assert.NoError(err)
	assert.Equal(expected.Partition, "partition")
	assert.Equal(expected.Service, "service")
	assert.Equal(expected.AccountID, "account")
	assert.Equal(expected.ResourceType, "resourcetype")
	assert.Equal(expected.Resource, "resource")
}

func DoesUnmarshalORNWithMultipleResourceWithSuccess(t *testing.T) {
	assert := assert.New(t)
	expected := ORN{}
	err := Unmarshal("orn:acme:judge:account:user/837/bar", &expected)

	assert.NoError(err)
	assert.Equal(expected.Partition, "acme")
	assert.Equal(expected.Service, "judge")
	assert.Equal(expected.AccountID, "account")
	assert.Equal(expected.ResourceType, "user")
	assert.Equal(expected.Resource, "837/bar")
}

func DoesUnmarshalORNWithoutAccountIDWithSuccess(t *testing.T) {
	assert := assert.New(t)
	expected := ORN{}
	err := Unmarshal("orn:acme:judge::user/837/bar", &expected)

	assert.NoError(err)

	assert.Equal(expected.Partition, "acme")
	assert.Equal(expected.Service, "judge")
	assert.Equal(expected.AccountID, "")
	assert.Equal(expected.ResourceType, "user")
	assert.Equal(expected.Resource, "837/bar")
}

func TestUnmarshal(t *testing.T) {
	DoesUnmarshalORNWithSuccess(t)
	DoesUnmarshalORNWithMultipleResourceWithSuccess(t)
	DoesUnmarshalORNWithoutAccountIDWithSuccess(t)
}

func TestMarshal(t *testing.T) {
	assert := assert.New(t)
	id := ORN{Partition: "foo", Service: "bar", AccountID: "baz", ResourceType: "biz", Resource: "fiz"}
	expected := Marshal(&id)
	assert.Equal(expected, "orn:foo:bar:baz:biz/fiz")
}

func TestString(t *testing.T) {
	assert := assert.New(t)
	id := ORN{Partition: "foo", Service: "bar", AccountID: "baz", ResourceType: "biz", Resource: "fiz"}
	assert.Equal(id.String(), "orn:foo:bar:baz:biz/fiz")
}
