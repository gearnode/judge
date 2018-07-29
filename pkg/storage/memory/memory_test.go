/*
Copyright 2018 The Judge Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package memorystore

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type ValueTest struct {
	ID    string
	Value string
}

func TestDescribeAll(t *testing.T) {
}

func TestDescribe(t *testing.T) {
	store := NewMemoryStore()
	assert.Equal(t, 0, len(store.Data()["fridge"]))

	_, err := store.Describe("fridge", "someID")
	assert.NotNil(t, err)

	milk := ValueTest{ID: "milk", Value: "2 bottles"}
	err = store.Put("fridge", milk.ID, milk)
	assert.Nil(t, err)

	_, err = store.Describe("", "someID")
	assert.NotNil(t, err)

	expected, err := store.Describe("fridge", milk.ID)
	assert.Nil(t, err)
	assert.Equal(t, milk, expected.(ValueTest))
}

func TestPut(t *testing.T) {
	store := NewMemoryStore()

	// Ensure memory is empty
	assert.Equal(t, 0, len(store.Data()["values"]))

	// Insert data
	v := ValueTest{ID: "foobar", Value: "hello world"}
	store.Put("values", v.ID, v)

	assert.Equal(t, 1, len(store.Data()["values"]))

	// Insert other data
	v = ValueTest{ID: "bazbar", Value: "hello world 2"}
	store.Put("values", v.ID, v)

	assert.Equal(t, 2, len(store.Data()["values"]))

	// Insert already exist key
	store.Put("values", v.ID, v)

	assert.Equal(t, 2, len(store.Data()["values"]))
}

func TestDelete(t *testing.T) {
}
