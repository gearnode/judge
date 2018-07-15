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

package policy // import "github.com/gearnode/judge/pkg/policy"

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gearnode/judge/pkg/orn"
	"github.com/gearnode/judge/pkg/policy/resource"
	"github.com/gearnode/judge/pkg/storage"
	"github.com/gosimple/slug"
)

// Policy is an entity in Judge that, when attached to an identity, defines
// their permissions. Judge evaluates these policies when a principal, such as
// a user, makes a request. Permissions in the policies determine whether the
// request is allowed or denied.
type Policy struct {
	// ORN element specifies a global unique identifier for the policy.
	ORN orn.ORN `json:"orn"`

	// Name element specifies a user friendly name for the policy.
	Name string `json:"name"`

	// Description element specifies description/usage about the policy.
	Description string `json:"description"`

	// Type element specifies the type for the policy.
	Type string `json:"type"`

	// Document contains all statements for the policy.
	Document Document `json:"document"`
}

// Document contains statements and version of these statements.
type Document struct {
	// Version is the Statement version.
	Version string `json:"version"`

	// Statement contains a list of Statement.
	Statement []Statement `json:"statement"`
}

// The Statement element is the main element for a policy. It defines
// permissions.
type Statement struct {
	// The Effect element is required and specifies whether the statement
	// results in an allow or an explicit deny. Valid values for Effect are
	// Allow and Deny.
	Effect string `json:"effect"`

	// The Action element describes the specific action or actions that will
	// be allowed or denied.
	Action []string `json:"action"`

	// The Resource element specifies the object or objects that the statement
	// covers.
	Resource []resource.Resource `json:"resource"`
}

const (
	PARTITION  = "judge-org"
	SERVICE    = "judge-server"
	STANDALONE = "STANDALONE"
	VERSION    = "v1alpha1"
)

var (
	// ErrMalformedPolicy was return when the policy is malformed.
	ErrMalformedPolicy = errors.New("malformed policy")
)

// NewPolicy create a new policy document.
func NewPolicy(name string, description string) *Policy {
	return &Policy{
		ORN: orn.ORN{
			Partition: PARTITION, Service: SERVICE,
			ResourceType: "policy", Resource: slug.Make(name),
		},
		Name:        name,
		Description: description,
		Type:        STANDALONE,
		Document: Document{
			Version: VERSION,
		},
	}
}

func NewStatement(effect string, actions []string, resources []string) (*Statement, error) {
	if effect != "Allow" && effect != "Deny" {
		return &Statement{}, fmt.Errorf("The effect %s is not supported."+
			" Supported effects are \"Allow\" or \"Deny\"", effect)
	}

	if len(resources) == 0 {
		return &Statement{}, errors.New("The statement object require at least" +
			" one resource.")
	}

	if len(actions) == 0 {
		return &Statement{}, errors.New("The statement object require at least" +
			" one action.")
	}

	for _, action := range actions {
		if action == "" {
			return &Statement{}, errors.New("The statement object does not support" +
				" empty action.")
		}
	}

	stmt := Statement{
		Effect:   effect,
		Action:   actions,
		Resource: make([]resource.Resource, len(resources)),
	}

	for i, rsrc := range resources {
		err := resource.Unmarshal(rsrc, &stmt.Resource[i])
		if err != nil {
			return &Statement{}, err
		}
	}

	return &stmt, nil
}

// CreatePolicy foo
func CreatePolicy(store storage.DB, name string, description string, doc string) (bool, error) {
	pol := NewPolicy(name, description)

	// TODO: use reflect for data casting

	document := make(map[string]interface{})
	err := json.Unmarshal([]byte(doc), &document)
	if err != nil {
		return false, err
	}

	stmts := document["statement"].([]interface{})
	pol.Document.Statement = make([]Statement, len(stmts))

	for i, raw := range stmts {
		data := raw.(map[string]interface{})
		stmt, err := NewStatement(
			data["effect"].(string),
			toStringArray(data["action"].([]interface{})),
			toStringArray(data["resource"].([]interface{})),
		)

		if err != nil {
			return false, err
		}

		pol.Document.Statement[i] = *stmt
	}

	data := make(map[string]interface{})
	if json.Unmarshal([]byte(doc), &data) != nil {
		return false, ErrMalformedPolicy
	}

	store.Put("policies", pol.Name, pol)
	return true, nil
}

func toStringArray(data []interface{}) []string {
	strs := make([]string, len(data))

	for i := range data {
		strs[i] = data[i].(string)
	}

	return strs
}
