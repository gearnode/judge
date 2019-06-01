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
)

type Constant struct {
	Value interface{}
	Type  string
}

func (c *Constant) Children() []Expr {
	return nil
}

func (c *Constant) Eval() Expr {
	return &Constant{Value: c.Value, Type: c.Type}
}

func (c *Constant) String() string {
	return fmt.Sprintf("%v", c.Value)
}

func (c *Constant) GoString() string {
	return fmt.Sprintf("&Constant{Value: %q, Type: %q}", c.Value, c.Type)
}
