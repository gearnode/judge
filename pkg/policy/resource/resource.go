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

package resource // import "github.com/gearnode/judge/pkg/policy/resource"

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gearnode/judge/pkg/orn"
)

const (
	allowAll  = "*"
	firstPart = "orn"
	partSep   = ":"
	partSize  = 5
	subSep    = "/"
	subSize   = 2 // sub parts size
)

var (
	// ErrMalformed is returned when the ORN appears to be invalid.
	ErrMalformed = errors.New("malformed ORN")
)

// Marshal accepts an ORN Struct and attempts to join it into constiuent parts.
func Marshal(rsrc *Resource) string {
	return fmt.Sprintf(
		"orn:%s:%s:%s:%s/%s",
		rsrc.Partition,
		rsrc.Service,
		rsrc.AccountID,
		rsrc.ResourceType,
		rsrc.Resource,
	)
}

// UnmarshalResource accepts an string and attempts convert this string in
// Resource Go struct.
func Unmarshal(data string, orn *Resource) error {
	parts := strings.Split(data, partSep)
	if len(parts) != partSize {
		return ErrMalformed
	} else if parts[0] != firstPart {
		return ErrMalformed
	}

	// TODO: Not force policy to have two sub.
	// This change force breaking change on the current API.
	sub := strings.SplitN(parts[4], subSep, 2)
	if len(sub) != subSize {
		return ErrMalformed
	}

	// Don't validate the last part because this part contain / to seperate
	// resourcetype/resource
	for i := 0; i < len(parts)-1; i++ {
		if !containsOnlyPermitedChar(parts[i]) {
			return ErrMalformed
		}
	}

	if !containsOnlyPermitedChar(sub[0]) {
		return ErrMalformed
	}

	orn.Partition = parts[1]
	orn.Service = parts[2]
	orn.AccountID = parts[3]
	orn.ResourceType = sub[0]
	orn.Resource = sub[1]
	return nil
}

func containsOnlyPermitedChar(s string) bool {
	for i := 0; i < len(s); i++ {
		if (s[i] < 'a' || s[i] > 'z') && s[i] != '-' && s[i] != '*' {
			return false
		}
	}
	return true
}

// Resource foo
type Resource orn.ORN

// Match
func (resource *Resource) Match(something *orn.ORN) bool {
	return resource.haveSamePartition(something) &&
		resource.haveSameService(something) &&
		resource.haveSameAccountID(something) &&
		resource.haveMatchResourceType(something) &&
		resource.haveMatchResourcePath(something)
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
	return resource.ResourceType == entity.ResourceType ||
		resource.ResourceType == allowAll
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
		if resourcePathParts[i] != entityPathParts[i] &&
			resourcePathParts[i] != allowAll {
			return false
		}

		resourceLastPart := i+1 >= x
		entityLastPart := i+1 >= y

		if resourceLastPart && x < y && resourcePathParts[i] == allowAll {
			return true
		}

		if resourceLastPart && entityLastPart &&
			(resourcePathParts[i] == entityPathParts[i] ||
				resourcePathParts[i] == allowAll) {
			return true
		}
	}
	return false
}
