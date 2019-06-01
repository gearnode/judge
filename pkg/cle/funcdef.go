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

type funcdef struct {
	Name    string
	ArgType string
	Arity   int
}

var (
	catalog = []*funcdef{
		&funcdef{
			Name:    "and",
			ArgType: "bool",
			Arity:   -1,
		},
		&funcdef{
			Name:    "or",
			ArgType: "bool",
			Arity:   -1,
		},
		&funcdef{
			Name:    "string:equals",
			ArgType: "string",
			Arity:   2,
		},
		&funcdef{
			Name:    "datetime:greater_than",
			ArgType: "string",
			Arity:   2,
		},
		&funcdef{
			Name:    "datetime:less_then",
			ArgType: "string",
			Arity:   2,
		},
	}
)
