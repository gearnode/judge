package authorize_test

import (
	"testing"

	"github.com/gearnode/judge/pkg/orn"
	"github.com/gearnode/judge/pkg/policy"
	"github.com/gearnode/judge/pkg/storage/memory"
	"github.com/gearnode/judge/pkg/authorize"

	"github.com/stretchr/testify/assert"
)

var (
	policies = []judge.Policy{
		{
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
					{
						Effect: "Allow",
						Action: []string{"eatService:Take", "eatService:Eat", "eatService:Describe"},
						Resource: []judge.Resource{
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
						Resource: []judge.Resource{
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

func prepare() {
	for _, v := range policies {
		store.Put("policies", v.ID, v)
	}
}

func clean() {
	store = memorystore.NewMemoryStore()
}

func TestAuthorize(t *testing.T) {
	prepare()
	t.Run("it's work", func(t *testing.T) {
		entity := orn.ORN{
			Partition:    "foo-company",
			Service:      "eatService",
			ResourceType: "food",
			Resource:     "tomato",
		}

		ok, err := authorize.Authorize(store, entity, "eatService:Eat")
		assert.Nil(t, err)
		assert.True(t, ok)

		ok, err = authorize.Authorize(store, entity, "eatService:BadAction")
		assert.NotNil(t, err)
		assert.False(t, ok)

		ok, err = authorize.Authorize(store, entity, "eatService:Take")
		assert.Nil(t, err)
		assert.True(t, ok)

		ok, err = authorize.Authorize(store, entity, "eatService:Describe")
		assert.NotNil(t, err)
		assert.False(t, ok)
	})
	clean()
}

func BenchmarkAuthorize(b *testing.B) {
	prepare()
	entity := orn.ORN{
		Partition:    "foo-company",
		Service:      "eatService",
		ResourceType: "food",
		Resource:     "tomato",
	}

	for i := 0; i < b.N; i++ {
		authorize.Authorize(store, entity, "eatService:Take")
		authorize.Authorize(store, entity, "eatService:BadAction")
	}
	clean()
}
