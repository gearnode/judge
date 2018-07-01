package authorize

import (
	"errors"

	"github.com/gearnode/judge/pkg/orn"
	"github.com/gearnode/judge/pkg/policy"
	"github.com/gearnode/judge/pkg/storage"
)

const (
	allowAction       = "Allow"
	denyAction        = "Deny"
	notMatchStatement = ""
)

var (
	// ErrAccessDenied was return when the user does not have the sufficient permissions.
	ErrAccessDenied = errors.New("the user does not have sufficient permissions")
)

// Authorize eval a list of policy
//
// By default => access denied
// Eval all the policies every time to ensure all denied policies are eval
//
// TODO: Add context
// TODO: Add interpolcation support on policy
func Authorize(pstore storage.DB, entityORN orn.ORN, action string) (bool, error) {
	policies, err := pstore.DescribeAll("policies")
	if err != nil {
		return false, err
	}

	state := notMatchStatement
	for _, obj := range policies {
		p := obj.(judge.Policy)

		for _, statement := range p.Document.Statement {
			effect := evalStatement(&statement, &entityORN, &action)
			if effect == denyAction {
				return false, ErrAccessDenied
			} else if effect == allowAction && state == notMatchStatement {
				state = allowAction
			}
		}
	}
	if state == allowAction {
		return true, nil
	}
	return false, ErrAccessDenied
}

func evalStatement(statement *judge.Statement, entityORN *orn.ORN, action *string) string {
	if matchAction(statement.Action, *action) {
		for _, resource := range statement.Resource {
			if resource.Match(entityORN) {
				return statement.Effect
			}
		}
	}
	return notMatchStatement
}

func matchAction(actions []string, action string) bool {
	for _, v := range actions {
		if v == action {
			return true
		}
	}
	return false
}
