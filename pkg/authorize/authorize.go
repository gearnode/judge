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

// Package authorize implement the policy evalution.
package authorize // import "github.com/gearnode/judge/pkg/authorize"

import (
	"errors"

	"github.com/gearnode/judge/pkg/orn"
	"github.com/gearnode/judge/pkg/policy"
	"github.com/gearnode/judge/pkg/storage"
)

const (
	allowAction = "Allow"
	denyAction  = "Deny"
	notMatch    = ""
)

var (
	// ErrAccessDenied was return when the user does not have the sufficient
	// permissions.
	ErrAccessDenied = errors.New("the user does not have sufficient permissions")
)

// Authorize decides whether a given request should be allowed or denied.
//
// The evaluation logic follows these rules:
//   By default, all requests are denied.
//   An explicit allow overrides this default.
//   An explicit deny overrides any allows.
//
// The order in which the policies are evaluated has no effect on the outcome
// of the evaluation. All policies are evaluated, and the result is always
// that the request is either allowed or denied.
func Authorize(s storage.DB, who orn.ORN, what string, something orn.ORN, context map[string]string) (bool, error) {

	defaultEffect := denyAction

	pols, err := s.DescribeAll("policies")
	if err != nil {
		return false, err
	}

	for _, data := range pols {
		pol := data.(policy.Policy)
		for _, stmt := range pol.Document.Statement {
			effect := evalStme(&stmt, &something, &what)

			if effect == denyAction { // Explicit deny
				return false, ErrAccessDenied
			} else if effect == allowAction { // Explicit allow
				defaultEffect = allowAction
			}
		}
	}

	if defaultEffect == allowAction {
		return true, nil
	}
	return false, ErrAccessDenied
}

func evalStme(stmt *policy.Statement, something *orn.ORN, what *string) string {
	if matchAction(stmt.Action, *what) {
		for _, resource := range stmt.Resource {
			if resource.Match(something) {
				return stmt.Effect
			}
		}
	}
	return notMatch
}

func matchAction(actions []string, what string) bool {
	for _, action := range actions {
		if action == what {
			return true
		}
	}
	return false
}
