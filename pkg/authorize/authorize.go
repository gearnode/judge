/*
Copyright 2019 Bryan Frimin,

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

package authorize

import (
	"errors"
	"github.com/gearnode/judge/pkg/orn"
	"github.com/gearnode/judge/pkg/policy"
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
func Authorize(policies []*policy.Policy, action string, object orn.ORN, _context map[string]string) error {
	finalEffect := denyAction // By default deny

	for _, policy := range policies {
		for _, statement := range policy.Statements {

			// TODO: @gearnode Add context support when policy object support condition
			//                 The context should be evaluate at the last check to avoid
			//                 long evaluation on not matching statement.
			currentEffect := evalStatement(statement, object, action)

			// Explicit deny (stop the evaluation)
			// No allow can override an deny action.
			if currentEffect == denyAction {
				return ErrAccessDenied
			}

			// Explicit allow (modify the result effect and don't break the evaluation).
			// Next statement can deny this statement allow.
			if currentEffect == allowAction {
				finalEffect = allowAction
			}
		}
	}

	if finalEffect == allowAction {
		return nil
	}

	return ErrAccessDenied
}

func evalStatement(statement policy.Statement, object orn.ORN, action string) string {
	if matchAction(statement.Actions, action) {
		for _, resource := range statement.Resources {
			if resource.Match(&object) {
				return statement.Effect
			}
		}
	}
	return notMatch
}

func matchAction(actions []string, action string) bool {
	for i := range actions {
		if actions[i] == action {
			return true
		}
	}
	return false
}
