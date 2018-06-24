package judge_test

import (
	"github.com/gearnode/judge/pkg/orn"
	"github.com/gearnode/judge/pkg/policy"
	"github.com/gearnode/judge/pkg/storage/memory"
	"github.com/stretchr/testify/assert"
	"testing"
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

		ok, err := judge.Authorize(store, entity, "eatService:Eat")
		assert.Nil(t, err)
		assert.True(t, ok)

		ok, err = judge.Authorize(store, entity, "eatService:BadAction")
		assert.NotNil(t, err)
		assert.False(t, ok)

		ok, err = judge.Authorize(store, entity, "eatService:Take")
		assert.Nil(t, err)
		assert.True(t, ok)

		ok, err = judge.Authorize(store, entity, "eatService:Describe")
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
		judge.Authorize(store, entity, "eatService:Take")
		judge.Authorize(store, entity, "eatService:BadAction")
	}
	clean()
}

func emptyDatabase(t *testing.T) {
	l, err := store.DescribeAll("policies")
	assert.Nil(t, err)
	assert.Equal(t, 0, len(l))
}

func TestCreatePolicy(t *testing.T) {
	emptyDatabase(t)

	ok, err := judge.CreatePolicy(
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
	ok, err = judge.CreatePolicy(
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
	ok, err = judge.CreatePolicy(
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
	ok, err = judge.CreatePolicy(
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
	ok, err = judge.CreatePolicy(
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
