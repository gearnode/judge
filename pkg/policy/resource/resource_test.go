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

package resource

import (
	"github.com/gearnode/judge/pkg/orn"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMarshal(t *testing.T) {
	rsrc := Resource{
		Partition: "judge-org", Service: "judge-server",
		ResourceType: "policy", Resource: "*",
		AccountID: "",
	}

	assert.Equal(t, "orn:judge-org:judge-server::policy/*", Marshal(&rsrc))

	rsrc = Resource{
		Partition: "some-org", Service: "some-svc",
		ResourceType: "food", Resource: "*",
		AccountID: "gearnode",
	}

	assert.Equal(t, "orn:some-org:some-svc:gearnode:food/*", Marshal(&rsrc))
}

func TestUnmarshal(t *testing.T) {
	t.Run("returns error when the number of parts are invalid", func(t *testing.T) {
		rsrc := Resource{}
		str := "orn:foo:foo:foo"
		assert.NotNil(t, Unmarshal(str, &rsrc))

		rsrc = Resource{}
		str = "orn:foo:foo:foo:foo:foo:foo:foo:foo/ddd/d/d/d"
		assert.NotNil(t, Unmarshal(str, &rsrc))
	})

	t.Run("returns error when the first part is not orn", func(t *testing.T) {
		rsrc := Resource{}
		str := "bar:judge-org:judge-server:gearnode:policy/*"
		assert.NotNil(t, Unmarshal(str, &rsrc))

		rsrc = Resource{}
		str = ":judge-org:judge-server:gearnode:policy/*"
		assert.NotNil(t, Unmarshal(str, &rsrc))

		rsrc = Resource{}
		str = "o rn:judge-org:judge-server:gearnode:policy/*"
		assert.NotNil(t, Unmarshal(str, &rsrc))
	})

	t.Run("returns error when sub ", func(t *testing.T) {
		rsrc := Resource{}
		str := "orn:judge-org:judge-server:gearnode:policy"
		assert.NotNil(t, Unmarshal(str, &rsrc))

		rsrc = Resource{}
		str = "orn:some-org:some-server::foobar"
		assert.NotNil(t, Unmarshal(str, &rsrc))
	})
}

func TestMatch(t *testing.T) {
	t.Run("describe Partition match", func(t *testing.T) {
		t.Run("does not match with * partition", func(t *testing.T) {

			var (
				entity = orn.ORN{
					Partition:    "food-company",
					Service:      "eatService",
					ResourceType: "food",
					Resource:     "milk/goat",
				}

				resource = Resource{
					Partition:    "*",
					Service:      "eatService",
					ResourceType: "food",
					Resource:     "milk/goat",
				}
			)

			assert.False(t, resource.Match(&entity))

			resource.Partition = "food*"
			assert.False(t, resource.Match(&entity))

			resource.Partition = "*company"
			assert.False(t, resource.Match(&entity))

		})

		t.Run("does not match with different partition", func(t *testing.T) {

			var (
				entity = orn.ORN{
					Partition:    "food-company",
					Service:      "eatService",
					ResourceType: "food",
					Resource:     "milk/goat",
				}

				resource = Resource{
					Partition:    "hell-company",
					Service:      "eatService",
					ResourceType: "food",
					Resource:     "milk/goat",
				}
			)

			assert.False(t, resource.Match(&entity))

			entity.Partition = "other-company"
			assert.False(t, resource.Match(&entity))
		})

		t.Run("does match with same partition", func(t *testing.T) {

			var (
				entity = orn.ORN{
					Partition:    "food-company",
					Service:      "eatService",
					ResourceType: "food",
					Resource:     "milk/goat",
				}

				resource = Resource{
					Partition:    "food-company",
					Service:      "eatService",
					ResourceType: "food",
					Resource:     "milk/goat",
				}
			)

			assert.True(t, resource.Match(&entity))

			entity.Partition = "acme-inc"
			resource.Partition = "acme-inc"

			assert.True(t, resource.Match(&entity))
		})
	})

	t.Run("describe Service match", func(t *testing.T) {
		t.Run("does not match with * service", func(t *testing.T) {

			var (
				entity = orn.ORN{
					Partition:    "food-company",
					Service:      "eatService",
					ResourceType: "food",
					Resource:     "milk/goat",
				}

				resource = Resource{
					Partition:    "food-company",
					Service:      "*",
					ResourceType: "food",
					Resource:     "milk/goat",
				}
			)

			assert.False(t, resource.Match(&entity))

			resource.Service = "eat*"
			assert.False(t, resource.Match(&entity))

			resource.Service = "*Service"
			assert.False(t, resource.Match(&entity))
		})

		t.Run("does not match with different service", func(t *testing.T) {
			var (
				entity = orn.ORN{
					Partition:    "food-company",
					Service:      "eatService",
					ResourceType: "food",
					Resource:     "milk/goat",
				}

				resource = Resource{
					Partition:    "food-company",
					Service:      "buyService",
					ResourceType: "food",
					Resource:     "milk/goat",
				}
			)

			assert.False(t, resource.Match(&entity))

			entity.Service = "invoiceService"
			assert.False(t, resource.Match(&entity))
		})

		t.Run("does match with same service", func(t *testing.T) {
			var (
				entity = orn.ORN{
					Partition:    "food-company",
					Service:      "eatService",
					ResourceType: "food",
					Resource:     "milk/goat",
				}

				resource = Resource{
					Partition:    "food-company",
					Service:      "eatService",
					ResourceType: "food",
					Resource:     "milk/goat",
				}
			)

			assert.True(t, resource.Match(&entity))

			entity.Service = "invoiceService"
			resource.Service = "invoiceService"
			assert.True(t, resource.Match(&entity))
		})
	})

	t.Run("describe AccountID match", func(t *testing.T) {
		t.Run("does not match with * account", func(t *testing.T) {

			var (
				entity = orn.ORN{
					Partition:    "food-company",
					Service:      "eatService",
					ResourceType: "food",
					Resource:     "milk/goat",
				}

				resource = Resource{
					Partition:    "food-company",
					Service:      "eatService",
					AccountID:    "*",
					ResourceType: "food",
					Resource:     "milk/goat",
				}
			)

			assert.False(t, resource.Match(&entity))

			entity.AccountID = "gearnode"
			resource.AccountID = "gear*"
			assert.False(t, resource.Match(&entity))

			resource.AccountID = "*node"
			assert.False(t, resource.Match(&entity))
		})

		t.Run("does not match with different account", func(t *testing.T) {
			var (
				entity = orn.ORN{
					Partition:    "food-company",
					Service:      "eatService",
					AccountID:    "gearnode",
					ResourceType: "food",
					Resource:     "milk/goat",
				}

				resource = Resource{
					Partition:    "food-company",
					Service:      "buyService",
					AccountID:    "superman",
					ResourceType: "food",
					Resource:     "milk/goat",
				}
			)

			assert.False(t, resource.Match(&entity))

			resource.AccountID = "spiderman"
			assert.False(t, resource.Match(&entity))
		})

		t.Run("does match with same account", func(t *testing.T) {
			var (
				entity = orn.ORN{
					Partition:    "food-company",
					Service:      "eatService",
					ResourceType: "food",
					Resource:     "milk/goat",
				}

				resource = Resource{
					Partition:    "food-company",
					Service:      "eatService",
					ResourceType: "food",
					Resource:     "milk/goat",
				}
			)

			assert.True(t, resource.Match(&entity))

			entity.AccountID = "gearnode"
			resource.AccountID = "gearnode"
			assert.True(t, resource.Match(&entity))
		})
	})

	t.Run("describe ResourceType match", func(t *testing.T) {
		t.Run("does supports the * operator", func(t *testing.T) {

			var (
				entity = orn.ORN{
					Partition:    "food-company",
					Service:      "eatService",
					ResourceType: "food",
					Resource:     "milk/goat",
				}

				resource = Resource{
					Partition:    "food-company",
					Service:      "eatService",
					ResourceType: "*",
					Resource:     "milk/goat",
				}
			)

			assert.True(t, resource.Match(&entity))

			entity.ResourceType = "gearnode"
			assert.True(t, resource.Match(&entity))

			resource.ResourceType = "gear*"
			assert.False(t, resource.Match(&entity))

			resource.ResourceType = "*node"
			assert.False(t, resource.Match(&entity))
		})

		t.Run("does not match with different ResourceType", func(t *testing.T) {
			var (
				entity = orn.ORN{
					Partition:    "food-company",
					Service:      "eatService",
					AccountID:    "gearnode",
					ResourceType: "food",
					Resource:     "milk/goat",
				}

				resource = Resource{
					Partition:    "food-company",
					Service:      "buyService",
					AccountID:    "gearnode",
					ResourceType: "bin",
					Resource:     "milk/goat",
				}
			)

			assert.False(t, resource.Match(&entity))

			resource.ResourceType = "foo"
			assert.False(t, resource.Match(&entity))
		})

		t.Run("does match with same ResourceType", func(t *testing.T) {
			var (
				entity = orn.ORN{
					Partition:    "food-company",
					Service:      "eatService",
					ResourceType: "food",
					Resource:     "milk/goat",
				}

				resource = Resource{
					Partition:    "food-company",
					Service:      "eatService",
					ResourceType: "food",
					Resource:     "milk/goat",
				}
			)

			assert.True(t, resource.Match(&entity))

			entity.ResourceType = "stock"
			resource.ResourceType = "stock"
			assert.True(t, resource.Match(&entity))
		})
	})

	t.Run("describe Resource match", func(t *testing.T) {
		t.Run("does supports the * operator", func(t *testing.T) {

			var (
				entity = orn.ORN{
					Partition:    "food-company",
					Service:      "eatService",
					ResourceType: "food",
					Resource:     "milk/goat",
				}

				resource = Resource{
					Partition:    "food-company",
					Service:      "eatService",
					ResourceType: "food",
					Resource:     "*",
				}
			)

			assert.True(t, resource.Match(&entity))

			entity.Resource = "milk/goat"
			assert.True(t, resource.Match(&entity))

			entity.Resource = "beef/horse/foo"
			assert.True(t, resource.Match(&entity))

			resource.Resource = "gear*"
			assert.False(t, resource.Match(&entity))

			resource.Resource = "*node"
			assert.False(t, resource.Match(&entity))

			entity.Resource = "beef/horse/foo"
			resource.Resource = "milk/*"
			assert.False(t, resource.Match(&entity))

			entity.Resource = ""
			resource.Resource = ""
			assert.True(t, resource.Match(&entity))

			entity.Resource = "milk/goat/foo"
			resource.Resource = "milk/*/bar"
			assert.False(t, resource.Match(&entity))

			entity.Resource = "milk/goat/foo"
			resource.Resource = "milk/*"
			assert.True(t, resource.Match(&entity))

			entity.Resource = "milk/goat/foo"
			resource.Resource = "milk/goat"
			assert.False(t, resource.Match(&entity))

			entity.Resource = "milk/goat"
			resource.Resource = "milk/goat"
			assert.True(t, resource.Match(&entity))

			entity.Resource = "milk/goat"
			resource.Resource = "milk/goat/*"
			assert.False(t, resource.Match(&entity))

			entity.Resource = "french/milk"
			resource.Resource = "*/milk/*/goat"
			assert.False(t, resource.Match(&entity))
		})
	})
}

// Benchmarks

func BenchmarkMatch(b *testing.B) {
	var (
		entity = orn.ORN{
			Partition:    "food-company",
			Service:      "eatService",
			ResourceType: "food",
			Resource:     "milk/foo",
		}

		resource = Resource{
			Partition:    "food-company",
			Service:      "eatService",
			ResourceType: "food",
			Resource:     "milk/*",
		}
	)

	for i := 0; i < b.N; i++ {
		resource.Match(&entity)
	}
}
