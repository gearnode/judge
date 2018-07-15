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
	"github.com/gearnode/judge/pkg/orn"
	"github.com/gearnode/judge/pkg/policy/resource"
	"github.com/gearnode/judge/pkg/storage/memory"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	policies = []Policy{
		{
			ORN: orn.ORN{
				Partition:    "foo-company",
				Service:      "judge",
				AccountID:    "666e735d-d046-4b8a-b2e8-407cb837a101",
				ResourceType: "policy",
				Resource:     "someid",
			},
			Name:        "allow_eat_tomato",
			Description: "allow user to eat tomato",
			Type:        "",
			Document: Document{
				Version: "2012-10-17",
				Statement: []Statement{
					{
						Effect: "Allow",
						Action: []string{"eatService:Take", "eatService:Eat", "eatService:Describe"},
						Resource: []resource.Resource{
							{
								Partition:    "foo-company",
								Service:      "eatService",
								AccountID:    "",
								ResourceType: "food",
								Resource:     "*",
							},
							{
								Partition:    "foo-company",
								Service:      "eatService",
								AccountID:    "",
								ResourceType: "stock",
								Resource:     "tomato/*",
							},
						},
					},
					{
						Effect: "Deny",
						Action: []string{"eatService:Describe"},
						Resource: []resource.Resource{
							{
								Partition:    "foo-company",
								Service:      "eatService",
								AccountID:    "",
								ResourceType: "food",
								Resource:     "*",
							},
						},
					},
				},
			},
		},
	}

	store = memorystore.NewMemoryStore()
)

func TestNewPolicy(t *testing.T) {
	pol := NewPolicy("some-name", "some description of the policy")
	assert.Equal(t, "some-name", pol.Name)
	assert.Equal(t, "some description of the policy", pol.Description)
	assert.Equal(t, STANDALONE, pol.Type)
	assert.Equal(t, VERSION, pol.Document.Version)
	assert.Equal(t, "judge-org", pol.ORN.Partition)
	assert.Equal(t, "judge-server", pol.ORN.Service)
	assert.Equal(t, "", pol.ORN.AccountID)
	assert.Equal(t, "policy", pol.ORN.ResourceType)
	assert.Equal(t, "some-name", pol.ORN.Resource)
}

func prepare() {
	for _, pol := range policies {
		store.Put("policies", pol.Name, pol)
	}
}

func clean() {
	store = memorystore.NewMemoryStore()
}

func emptyDatabase(t *testing.T) {
	l, err := store.DescribeAll("policies")
	assert.Nil(t, err)
	assert.Equal(t, 0, len(l))
}

func TestCreatePolicy(t *testing.T) {
	emptyDatabase(t)

	ok, err := CreatePolicy(
		store,
		"policy demo",
		"some description",
		`
		{
			"version": "2018-02-15",
			"statement": [
				{
					"effect": "Allow",
					"action": ["eatService:Take"],
					"resource": ["orn:judge-org:policy-service::foo/bar"]
				}
			]
		}
		`,
	)

	assert.True(t, ok)
	assert.Nil(t, err)
	l, err := store.DescribeAll("policies")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(l))
	clean()

	emptyDatabase(t)
	ok, err = CreatePolicy(
		store,
		"policy demo",
		"some description",
		`
		{
			"version": "2018-02-15",
			"statement": [
				{
					"effect": "",
					"action": ["eatService:Take"],
					"resource": ["orn:judge-org:policy-service::foo/bar"]
				}
			]
		}
		`,
	)

	assert.False(t, ok)
	assert.NotNil(t, err)
	emptyDatabase(t)
	clean()

	emptyDatabase(t)
	ok, err = CreatePolicy(
		store,
		"policy demo",
		"some description",
		`
		{
			"version": "2018-02-15",
			"statement": [
				{
					"effect": "Other",
					"action": ["eatService:Take"],
					"resource": ["orn:judge-org:policy-service::foo/bar"]
				}
			]
		}
		`,
	)

	assert.False(t, ok)
	assert.NotNil(t, err)
	emptyDatabase(t)
	clean()

	emptyDatabase(t)
	ok, err = CreatePolicy(
		store,
		"policy demo",
		"some description",
		`
		{
			"version": "2018-02-15",
			"statement": [
				{
					"effect": "Deny",
					"action": ["eatService:Take"],
					"resource": []
				}
			]
		}
		`,
	)

	assert.False(t, ok)
	assert.NotNil(t, err)
	emptyDatabase(t)
	clean()

	emptyDatabase(t)
	ok, err = CreatePolicy(
		store,
		"policy demo",
		"some description",
		`
		{
			"version": "2018-02-15",
			"statement": [
				{
					"effect": "Deny",
					"action": [],
					"resource": ["orn:judge-org:policy-service::foo/bar"]
				}
			]
		}
		`,
	)

	assert.False(t, ok)
	assert.NotNil(t, err)
	emptyDatabase(t)
	clean()
}
