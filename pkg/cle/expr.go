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

type Tree struct {
	Root Expr
}

func (t *Tree) Walk(fn func(Expr) error) error {
	return t.Root.Walk(fn)
}

type Expr interface {
	Eval() Expr
	Walk(func(Expr) error) error
	GoString() string
	String() string
}
