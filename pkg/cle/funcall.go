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
	"strings"
)

type Funcall struct {
	Name string
	Args []Expr
}

func (f *Funcall) Walk(fn func(Expr) error) error {
	if err := fn(f); err != nil {
		return err
	}

	for _, arg := range f.Args {
		if err := arg.Walk(fn); err != nil {
			return err
		}
	}
	return nil
}

func (f *Funcall) Eval() Expr {
	return f
}

func (f *Funcall) String() string {
	return "funcall(" + f.Name + "/" + strconv.Itoa(2) + ")"
}

func (f *Funcall) GoString() string {
	args := make([]string, len(f.Args))

	for i := range f.Args {
		args[i] = f.Args[i].GoString()
	}

	return fmt.Sprintf("&Funcall{Name: %s, Args: %s}", f.Name, strings.Join(args, ", "))
}
