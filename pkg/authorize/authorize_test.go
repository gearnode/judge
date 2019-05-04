/*
Copyright 2019 Bryan Frimin,

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

package authorize_test

import (
	"github.com/gearnode/judge/pkg/authorize"
	"github.com/gearnode/judge/pkg/orn"
	"github.com/gearnode/judge/pkg/policy"
	"github.com/gearnode/judge/pkg/policy/resource"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	policies = []*policy.Policy{
		{
			ID: orn.ORN{
				Partition:    "foo-company",
				Service:      "judge",
				AccountID:    "666e735d-d046-4b8a-b2e8-407cb837a101",
				ResourceType: "policy",
				Resource:     "someid",
			},
			Name:        "allow_eat_tomato",
			Description: "allow user to eat tomato",
			Statements: []policy.Statement{
				{
					Effect:  "Allow",
					Actions: []string{"eatService:Take", "eatService:Eat", "eatService:Describe"},
					Resources: []resource.Resource{
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
					Effect:  "Deny",
					Actions: []string{"eatService:Describe"},
					Resources: []resource.Resource{
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
	}
)

func TestAuthorize(t *testing.T) {
	t.Run("it's work", func(t *testing.T) {
		what := "eatService:Eat"
		something := orn.ORN{
			Partition:    "foo-company",
			Service:      "eatService",
			ResourceType: "food",
			Resource:     "tomato",
		}
		ctx := make(map[string]string, 1)

		err := authorize.Authorize(policies, what, something, ctx)
		assert.Nil(t, err)

		what = "eatService:BadAction"
		err = authorize.Authorize(policies, what, something, ctx)
		assert.NotNil(t, err)

		what = "eatService:Take"
		err = authorize.Authorize(policies, what, something, ctx)
		assert.Nil(t, err)

		what = "eatService:Describe"
		err = authorize.Authorize(policies, what, something, ctx)
		assert.NotNil(t, err)
	})
}

func BenchmarkAuthorize(b *testing.B) {
	something := orn.ORN{
		Partition:    "foo-company",
		Service:      "eatService",
		ResourceType: "food",
		Resource:     "tomato",
	}
	ctx := make(map[string]string, 1)

	for i := 0; i < b.N; i++ {
		authorize.Authorize(policies, "eatService:Take", something, ctx)
		authorize.Authorize(policies, "eatService:BadAction", something, ctx)
	}
}
