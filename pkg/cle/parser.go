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
)

/* TODO:
 * - Limit the nested to 7
 * - Have var type
 * - Check type before cast
 * - Error message
 */

func Parse(s []byte) (Expr, error) {
	var raw interface{}

	err := json.Unmarshal(s, &raw)
	if err != nil {
		return nil, err
	}

	return parseExpr(raw)
}

func parseExpr(value interface{}) (Expr, error) {
	switch v := value.(type) {
	case string:
		return &Constant{Value: v, Type: "string"}, nil
	case []interface{}:
		return parseFuncall(v)
	}

	return nil, fmt.Errorf("syntax error")
}

func parseFuncall(v []interface{}) (*Funcall, error) {
	if len(v) == 0 {
		return nil, fmt.Errorf("Funcall should have argument(s)")
	}

	name, ok := v[0].(string)
	if !ok {
		return nil, fmt.Errorf("bad function name")
	}

	args := make([]Expr, len(v)-1)

	for i, arg := range v[1:] {
		x, err := parseExpr(arg)
		if err != nil {
			return nil, err
		}
		args[i] = x
	}

	return &Funcall{Name: name, Args: args}, nil
}
