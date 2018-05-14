package judge_test

import (
	"github.com/gearnode/judge"
	"github.com/gearnode/judge/orn"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	policies = []judge.Policy{
		judge.Policy{
			ID: "someid",
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
			Document: judge.Document{
				Version: "2012-10-17",
				Statement: []judge.Statement{
					judge.Statement{
						Effect: "Allow",
						Action: []string{"eatService:Take", "eatService:Eat", "eatService:Describe"},
						Resource: []judge.Resource{
							judge.Resource{
								Partition:    "foo-company",
								Service:      "eatService",
								AccountID:    "",
								ResourceType: "food",
								Resource:     "*",
							},
							judge.Resource{
								Partition:    "foo-company",
								Service:      "eatService",
								AccountID:    "",
								ResourceType: "stock",
								Resource:     "tomato/*",
							},
						},
					},
					judge.Statement{
						Effect: "Deny",
						Action: []string{"eatService:Describe"},
						Resource: []judge.Resource{
							judge.Resource{
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
)

func TestAuthorize(t *testing.T) {
	t.Run("it's work", func(t *testing.T) {
		entity := orn.ORN{
			Partition:    "foo-company",
			Service:      "eatService",
			ResourceType: "food",
			Resource:     "tomato",
		}

		ok, err := judge.Authorize(policies, entity, "eatService:Eat")
		assert.Nil(t, err)
		assert.True(t, ok)

		ok, err = judge.Authorize(policies, entity, "eatService:BadAction")
		assert.NotNil(t, err)
		assert.False(t, ok)

		ok, err = judge.Authorize(policies, entity, "eatService:Take")
		assert.Nil(t, err)
		assert.True(t, ok)

		ok, err = judge.Authorize(policies, entity, "eatService:Describe")
		assert.NotNil(t, err)
		assert.False(t, ok)
	})
}

func BenchmarkAuthorize(b *testing.B) {
	entity := orn.ORN{
		Partition:    "foo-company",
		Service:      "eatService",
		ResourceType: "food",
		Resource:     "tomato",
	}

	for i := 0; i < b.N; i++ {
		judge.Authorize(policies, entity, "eatService:Take")
		judge.Authorize(policies, entity, "eatService:BadAction")
	}
}
