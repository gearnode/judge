package orn_test

import (
	"testing"

	"github.com/gearnode/judge/orn"
	"github.com/stretchr/testify/assert"
)

func TestUnmarshal(t *testing.T) {

	t.Run("does Unmarshal the ORN successfully", func(t *testing.T) {
		expected := orn.ORN{}
		err := orn.Unmarshal("orn:partition:service:account:resourcetype/resource", &expected)

		assert.Nil(t, err)

		assert.Equal(t, expected.Partition, "partition")
		assert.Equal(t, expected.Service, "service")
		assert.Equal(t, expected.AccountID, "account")
		assert.Equal(t, expected.ResourceType, "resourcetype")
		assert.Equal(t, expected.Resource, "resource")
	})

	t.Run("does Unmarshal the ORN successfully with multiple / in the resource", func(t *testing.T) {
		expected := orn.ORN{}
		err := orn.Unmarshal("orn:acme:judge:account:user/837/bar", &expected)

		assert.Nil(t, err)

		assert.Equal(t, expected.Partition, "acme")
		assert.Equal(t, expected.Service, "judge")
		assert.Equal(t, expected.AccountID, "account")
		assert.Equal(t, expected.ResourceType, "user")
		assert.Equal(t, expected.Resource, "837/bar")
	})

	t.Run("does Unmarshal the ORN successfully without account id", func(t *testing.T) {
		expected := orn.ORN{}
		err := orn.Unmarshal("orn:acme:judge::user/837/bar", &expected)

		assert.Nil(t, err)

		assert.Equal(t, expected.Partition, "acme")
		assert.Equal(t, expected.Service, "judge")
		assert.Equal(t, expected.AccountID, "")
		assert.Equal(t, expected.ResourceType, "user")
		assert.Equal(t, expected.Resource, "837/bar")
	})
}

func TestMarshal(t *testing.T) {
	o := orn.ORN{Partition: "foo", Service: "bar", AccountID: "baz", ResourceType: "biz", Resource: "fiz"}
	expected := orn.Marshal(&o)

	assert.Equal(t, expected, "orn:foo:bar:baz:biz/fiz")
}
