package judge

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gearnode/judge/orn"
	"github.com/gosimple/slug"
	"github.com/satori/go.uuid"
	"strings"
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
	Effect   string     `json:"effect"`
	Action   []string   `json:"action"`
	Resource []Resource `json:"resource"`
}

const (
	allowAction       = "Allow"
	denyAction        = "Deny"
	notMatchStatement = ""
)

var (
	// ErrAccessDenied was return when the user does not have the sufficient permissions.
	ErrAccessDenied = errors.New("the user does not have sufficient permissions")
	// ErrMalformedPolicy was return when the policy is malformed.
	ErrMalformedPolicy = errors.New("malformed policy")
)

// Authorize eval a list of policy
//
// By default => access denied
// Eval all the policies every time to ensure all denied policies are eval
//
// TODO: Add context
// TODO: Add interpolcation support on policy
func Authorize(pstore PolicyStore, entityORN orn.ORN, action string) (bool, error) {
	policies := pstore.GetAll()
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

func CreatePolicy(pstore PolicyStore, name string, description string, doc string) (bool, error) {
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
		statement.Resource = []Resource{}
		// TODO: implement unmarshal func for resource type
		for _, r := range z["resource"].([]interface{}) {
			resource := Resource{}
			err := UnmarshalResource(r.(string), &resource)
			if err != nil {
				return false, err
			}
			statement.Resource = append(statement.Resource, resource)
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

	pstore.Put(policy)
	return true, nil
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
func UnmarshalResource(data string, orn *Resource) error {
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
func Marshal(orn *Resource) string {
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
		if (s[i] < 'a' || s[i] > 'z') && s[i] != '-' && s[i] != '*' {
			return false
		}
	}
	return true
}
