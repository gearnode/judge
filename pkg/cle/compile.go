/*
 * Copyright 2019 Bryan Frimin
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cle

import (
	"encoding/json"
	"fmt"
	"time"
)

/* TODO:
 * - Limit the nested to 7
 * - Have var type
 * - Check type before cast
 * - Error message
 */

func Compile(s []byte) (Expr, error) {
	var raw interface{}

	err := json.Unmarshal(s, &raw)
	if err != nil {
		return nil, err
	}

	return decodeExpr(raw)
}

var functionCatalog = map[string]bool{
	"and":                           true,
	"or":                            true,
	"string:equals":                 true,
	"string:not_equals":             true,
	"string:equals_ignore_case":     true,
	"string:not_equals_ignore_case": true,
	"datetime:greater_than":         true,
	"datetime:less_then":            true,
}

var functionArity = map[string]int{
	"and":                           -1,
	"or":                            -1,
	"string:equals":                 2,
	"string:not_equals":             2,
	"string:equals_ignore_case":     2,
	"string:not_equals_ignore_case": 2,
	"datetime:greater_than":         2,
	"datetime:less_then":            2,
}

var functionType = map[string]string{
	"and":                           "bool",
	"or":                            "bool",
	"string:equals":                 "string",
	"string:not_equals":             "string",
	"string:equals_ignore_case":     "string",
	"string:not_equals_ignore_case": "string",
	"datetime:greater_than":         "datetime",
	"datetime:less_then":            "datetime",
}

func decodeExpr(value interface{}) (Expr, error) {
	switch v := value.(type) {
	case string:
		t, err := time.Parse(time.RFC3339, v)
		if err == nil {
			return &Constant{Value: t, Type: "datetime"}, nil
		}
		return &Constant{Value: v, Type: "string"}, nil
	case int64:
		// TOOD: float 64
		return &Constant{Value: v, Type: "integer"}, nil
	case bool:
		return &Constant{Value: v, Type: "boolean"}, nil
	case []interface{}:
		if len(v) == 0 {
			return nil, fmt.Errorf("Funcall should have argument(s)")
		}

		funcName, ok := v[0].(string)
		if !ok {
			return nil, fmt.Errorf("bad function name")
		}

		if !functionCatalog[funcName] {
			return nil, fmt.Errorf("The function does not exist")
		}

		// TODO change cond order
		if functionArity[funcName] != len(v[1:]) && functionArity[funcName] != -1 {
			return nil, fmt.Errorf("The function does have the right number of argument(s)")
		}

		funcArgs := make([]Expr, len(v)-1)

		for _, arg := range v[1:] {
			args, err := decodeExpr(arg)
			if err != nil {
				return nil, err
			}
			funcArgs = append(funcArgs, args)
		}

		return &Funcall{Name: funcName, Args: funcArgs}, nil
	}

	return nil, fmt.Errorf("unknown type")
}
