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

// Package orn represents ORN data structure in Go types.
package orn

import (
	"errors"
	"fmt"
	"strings"
)

// ORN represent a deserialized ORN in the Go Struct format.
type ORN struct {

	// The partition that the resource is in.
	Partition string

	// The service namespace that identifies the service where the resource live.
	Service string

	// The owner of the resource.
	AccountID string

	// The type of the resource.
	ResourceType string

	// The path of the resource.
	Resource string
}

const (
	// FirstPart represent the first part of an unmarshal ORN.
	FirstPart = "orn"

	// PartSep is the value used to separate ORN parts when ORN is marshaled.
	PartSep = ":"

	// PartSize is the number of piece of an ORN.
	PartSize = 5

	// SubSep is the sperator used to seprate the Resource and ResourceType.
	SubSep = "/"

	// SubSize is the number of piece of an ResourceType/Resource
	SubSize = 2 // sub parts size
)

var (
	// ErrMalformed is returned when the ORN appears to be invalid.
	ErrMalformed = errors.New("malformed ORN")
)

// Unmarshal accepts an ORN string and attempts to split it into constiuent parts.
func Unmarshal(data string, orn *ORN) error {
	parts := strings.Split(data, PartSep)
	if len(parts) != PartSize {
		return ErrMalformed
	} else if parts[0] != FirstPart {
		return ErrMalformed
	}

	sub := strings.SplitN(parts[4], SubSep, 2)
	if len(sub) != SubSize {
		return ErrMalformed
	}

	// Don't validate the last part because this part contain / to seperate
	// resourcetype/resource
	for i := 0; i < len(parts)-1; i++ {
		if !containsOnlyPermitedChar(parts[i]) {
			fmt.Println("bad char", parts[i])
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

// Marshal accepts an ORN Struct and attempts to join it into constiuent parts.
func Marshal(orn *ORN) string {
	return fmt.Sprintf(
		"orn:%s:%s:%s:%s/%s",
		orn.Partition,
		orn.Service,
		orn.AccountID,
		orn.ResourceType,
		orn.Resource,
	)
}

func containsOnlyPermitedChar(s string) bool {
	for i := 0; i < len(s); i++ {
		if (s[i] < 'a' || s[i] > 'z') && s[i] != '-' {
			return false
		}
	}
	return true
}
