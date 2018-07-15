package authorize_test

import (
	"testing"

	"github.com/gearnode/judge/pkg/authorize"
	"github.com/gearnode/judge/pkg/orn"
	"github.com/gearnode/judge/pkg/policy"
	"github.com/gearnode/judge/pkg/policy/resource"
	"github.com/gearnode/judge/pkg/storage/memory"

	"github.com/stretchr/testify/assert"
)

var (
	policies = []policy.Policy{
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
			Document: policy.Document{
				Version: "2012-10-17",
				Statement: []policy.Statement{
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

func prepare() {
	for _, pol := range policies {
		store.Put("policies", pol.Name, pol)
	}
}

func clean() {
	store = memorystore.NewMemoryStore()
}

func TestAuthorize(t *testing.T) {
	prepare()
	t.Run("it's work", func(t *testing.T) {
		who := orn.ORN{}
		what := "eatService:Eat"
		something := orn.ORN{
			Partition:    "foo-company",
			Service:      "eatService",
			ResourceType: "food",
			Resource:     "tomato",
		}
		ctx := make(map[string]string, 1)

		ok, err := authorize.Authorize(store, who, what, something, ctx)
		assert.Nil(t, err)
		assert.True(t, ok)

		what = "eatService:BadAction"
		ok, err = authorize.Authorize(store, who, what, something, ctx)
		assert.NotNil(t, err)
		assert.False(t, ok)

		what = "eatService:Take"
		ok, err = authorize.Authorize(store, who, what, something, ctx)
		assert.Nil(t, err)
		assert.True(t, ok)

		what = "eatService:Describe"
		ok, err = authorize.Authorize(store, who, what, something, ctx)
		assert.NotNil(t, err)
		assert.False(t, ok)
	})
	clean()
}

func BenchmarkAuthorize(b *testing.B) {
	prepare()

	who := orn.ORN{}
	something := orn.ORN{
		Partition:    "foo-company",
		Service:      "eatService",
		ResourceType: "food",
		Resource:     "tomato",
	}
	ctx := make(map[string]string, 1)

	for i := 0; i < b.N; i++ {
		authorize.Authorize(store, who, "eatService:Take", something, ctx)
		authorize.Authorize(store, who, "eatService:BadAction", something, ctx)
	}
	clean()
}
