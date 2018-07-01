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
	"encoding/json"
	"errors"
	"github.com/gearnode/judge/pkg/orn"
	"github.com/gearnode/judge/pkg/policy/resource"
	"github.com/gearnode/judge/pkg/storage"
	"github.com/gosimple/slug"
	"github.com/satori/go.uuid"
)

// Policy is an entity in Judge that, when attached to an identity, defines
// their permissions. Judge evaluates these policies when a principal, such as
// a user, makes a request. Permissions in the policies determine whether the
// request is allowed or denied.
type Policy struct {
	ID          string
	ORN         orn.ORN
	Name        string
	Description string
	Type        string
	Document    Document
}

// Document todo
type Document struct {
	Version   string      `json:"version"`
	Statement []Statement `json:"statement"`
}

// Statement todo
type Statement struct {
	Effect   string              `json:"effect"`
	Action   []string            `json:"action"`
	Resource []resource.Resource `json:"resource"`
}

var (
	// ErrMalformedPolicy was return when the policy is malformed.
	ErrMalformedPolicy = errors.New("malformed policy")
)

// CreatePolicy foo
func CreatePolicy(pstore storage.DB, name string, description string, doc string) (bool, error) {
	o := orn.ORN{
		Partition:    "judge-org",
		Service:      "judge-server",
		ResourceType: "policy",
		Resource:     slug.Make(name), // NOTE: keep slug dependency ?
	}

	data := make(map[string]interface{})
	if json.Unmarshal([]byte(doc), &data) != nil {
		return false, ErrMalformedPolicy
	}

	statements := []Statement{}

	// TODO: have more granular error message
	x := data["statement"].([]interface{})
	// ([]map[string]interface{})
	for _, v := range x {
		z := v.(map[string]interface{})
		statement := Statement{}
		statement.Effect = z["effect"].(string)
		if statement.Effect != "Allow" && statement.Effect != "Deny" {
			return false, ErrMalformedPolicy
		}

		for _, v := range z["action"].([]interface{}) {
			statement.Action = append(statement.Action, v.(string))
		}

		if len(statement.Action) <= 0 {
			return false, ErrMalformedPolicy
		}

		// NOTE: maybe merge resource in the ORN package
		statement.Resource = []resource.Resource{}
		// TODO: implement unmarshal func for resource type
		for _, r := range z["resource"].([]interface{}) {
			rsrc := resource.Resource{}
			err := resource.Unmarshal(r.(string), &rsrc)
			if err != nil {
				return false, err
			}
			statement.Resource = append(statement.Resource, rsrc)
		}

		if len(statement.Resource) <= 0 {
			return false, ErrMalformedPolicy
		}
		statements = append(statements, statement)
	}

	// TODO: handle uuid generation error
	policy := Policy{
		ORN:         o,
		ID:          uuid.Must(uuid.NewV4()).String(),
		Name:        name,
		Description: description,
		Type:        "STANDALONE",
		Document: Document{
			Version:   data["version"].(string),
			Statement: statements,
		},
	}

	pstore.Put("policies", policy.ID, policy)
	return true, nil
}
