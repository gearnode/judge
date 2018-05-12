package judge

import (
	"github.com/gearnode/judge/orn"
	"regexp"
	"fmt"
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
	// Replace * by .*
	var re = regexp.MustCompile(fmt.Sprintf(`(%[1]s|^)(\*)(%[1]s|$)`, orn.SubSep))
	s := string(re.ReplaceAllString(resource.Resource, "$1.*$3"))
	// Place start and end delimiters (^$)
	s = fmt.Sprintf("^%s$", s)

	matched, _ := regexp.MatchString(s, entity.Resource)
	return matched
}
