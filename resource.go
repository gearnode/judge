package judge

import (
	"github.com/gearnode/judge/orn"
	"strings"
)

const (
	allowAll = "*"
)

// Resource todo
type Resource orn.ORN

// Match todo
func (resource *Resource) Match(entity *orn.ORN) bool {
	return resource.haveSamePartition(entity) && resource.haveSameService(entity) &&
		resource.haveSameAccountID(entity) && resource.haveMatchResourceType(entity) &&
		resource.haveMatchResourcePath(entity)
}

func (resource *Resource) haveSamePartition(entity *orn.ORN) bool {
	return resource.Partition == entity.Partition
}

func (resource *Resource) haveSameService(entity *orn.ORN) bool {
	return resource.Service == entity.Service
}

func (resource *Resource) haveSameAccountID(entity *orn.ORN) bool {
	return resource.AccountID == entity.AccountID
}

func (resource *Resource) haveMatchResourceType(entity *orn.ORN) bool {
	return resource.ResourceType == entity.ResourceType || resource.ResourceType == allowAll
}

func (resource *Resource) haveMatchResourcePath(entity *orn.ORN) bool {
	resourcePathParts := strings.Split(resource.Resource, orn.SubSep)
	entityPathParts := strings.Split(entity.Resource, orn.SubSep)
	x := len(resourcePathParts)
	y := len(entityPathParts)

	if y < x {
		return false
	}

	for i := 0; i < x && i < y; i++ {
		if resourcePathParts[i] != entityPathParts[i] && resourcePathParts[i] != allowAll {
			return false
		}

		resourceLastPart := i+1 >= x
		entityLastPart := i+1 >= y

		if resourceLastPart && x < y && resourcePathParts[i] == allowAll {
			return true
		}

		if resourceLastPart && entityLastPart &&
			(resourcePathParts[i] == entityPathParts[i] || resourcePathParts[i] == allowAll) {
			return true
		}
	}
	return false
}
