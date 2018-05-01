package judge

import (
	"errors"
	"github.com/gearnode/judge/orn"
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
	Version   string
	Statement []Statement
}

// Statement todo
type Statement struct {
	Effect   string
	Action   []string
	Resource []Resource
}

const (
	allowAction       = "Allow"
	denyAction        = "Deny"
	notMatchStatement = ""
)

var (
	// ErrAccessDenied was return when the does not have the sufficient permissions.
	ErrAccessDenied = errors.New("the user does not have sufficient permissions")
)

// Authorize eval a list of policy
//
// By default => access denied
// Eval all the policies every time to ensure all denied policies are eval
//
// TODO: Add context
// TODO: Add interpolcation support on policy
func Authorize(policies []Policy, entityORN orn.ORN, action string) (bool, error) {
	state := notMatchStatement
	for _, policy := range policies {
		for _, statement := range policy.Document.Statement {
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

func evalStatement(statement *Statement, entityORN *orn.ORN, action *string) string {
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
