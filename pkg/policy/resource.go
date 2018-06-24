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

package judge

import (
	"github.com/gearnode/judge/pkg/orn"
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
