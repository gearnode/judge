/*
Copyright 2019 Bryan Frimin.

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

package policy

import (
	"errors"
	"fmt"
	"github.com/gearnode/judge/pkg/orn"
	"github.com/gearnode/judge/pkg/policy/resource"
	"github.com/gosimple/slug"
)

// Policy is an entity in Judge that, when attached to an identity, defines
// their permissions. Judge evaluates these policies when a principal, such as
// a user, makes a request. Permissions in the policies determine whether the
// request is allowed or denied.
type Policy struct {
	// ID element specifies a global unique identifier for the policy.
	ID orn.ORN `json:"orn"`

	// Name element specifies a user friendly name for the policy.
	Name string `json:"name"`

	// Description element specifies description/usage about the policy.
	Description string `json:"description"`

	// Statements contains a list of Statement.
	Statements []Statement `json:"statement"`
}

// The Statement element is the main element for a policy. It defines
// permissions.
type Statement struct {
	// Effect element is required and specifies whether the statement
	// results in an allow or an explicit deny. Valid values for Effect are
	// Allow and Deny.
	Effect string `json:"effect"`

	// Actions element describes the specific action or actions that will
	// be allowed or denied.
	Actions []string `json:"action"`

	// Resources element specifies the object or objects that the statement
	// covers.
	Resources []resource.Resource `json:"resource"`
}

const (
	PARTITION = "judge-org"
	SERVICE   = "judge-server"
)

var (
	// ErrMalformedPolicy was return when the policy is malformed.
	ErrMalformedPolicy = errors.New("malformed policy")
)

// NewPolicy create a new policy document.
func NewPolicy(name string, description string) (*Policy, error) {
	if name == "" {
		return nil, errors.New("the policy object require a non empty name")
	}

	return &Policy{
		ID: orn.ORN{
			Partition: PARTITION, Service: SERVICE,
			ResourceType: "policy", Resource: slug.Make(name),
		},
		Name:        name,
		Description: description,
	}, nil
}

func NewStatement(effect string, actions []string, resources []string) (*Statement, error) {
	if effect != "ALLOW" && effect != "DENY" {
		return nil, fmt.Errorf("the effect %q is not supported. Supported effects are %q or %q", "ALLOW", "DENY", effect)
	}

	if len(resources) == 0 {
		return nil, errors.New("the statement object require at least one resource")
	}

	if len(actions) == 0 {
		return nil, errors.New("the statement object require at least one action")
	}

	for _, action := range actions {
		if action == "" {
			return nil, errors.New("the statement object does not support empty action")
		}
	}

	statement := Statement{
		Effect:    effect,
		Actions:   actions,
		Resources: make([]resource.Resource, len(resources)),
	}

	for i := range resources {
		err := resource.Unmarshal(resources[i], &statement.Resources[i])
		if err != nil {
			return nil, err
		}
	}

	return &statement, nil
}
