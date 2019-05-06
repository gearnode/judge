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
	"fmt"
	"strconv"
	"time"
)

type Constant struct {
	Value interface{}
	Type  string
}

func (c *Constant) Eval() Expr {
	switch c.Type {
	case "string":
		return &Constant{Value: c.Value, Type: c.Type}
	}
	return &Constant{Value: c.Value, Type: c.Type}
}

func (c *Constant) String() string {
	switch v := c.Value.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case bool:
		return strconv.FormatBool(v)
	case time.Time:
		return v.String()
	}
	return "unknown type"
}

func (c *Constant) GoString() string {
	return fmt.Sprintf("&Constant{Value: %q}", c.Value)
}
