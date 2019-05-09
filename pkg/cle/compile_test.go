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
	"testing"
)

func TestFoo(t *testing.T) {
	tree, err := Compile([]byte(`
	["and",
		["string:equals", "judge:current_identiy", "gearnode"],
		["and",
			["datetime:greater_than", "judge:current_time", "2013-06-30T00:00:00Z"],
			["datetime:less_then", "judge:current_time", "2020-06-30T00:00:00Z"]
		]
	]
	`))

	fmt.Printf("Tree: %#v\n", tree)
	fmt.Printf("Error: %v\n", err)
}
