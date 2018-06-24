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

// Package memorystore contains an in memory implementation of
// the Judge storage interface.
package memorystore // import "github.com/gearnode/judge/pkg/storage/memory"

import (
	"errors"
	"sync"

	"github.com/gearnode/judge/pkg/storage"
)

// MemoryTable foo
type memoryTable map[string]memoryRow

// MemoryRow foo
type memoryRow map[string]interface{}

// MemoryStore is a simple in-memory storage using the DB interface.
type MemoryStore struct {
	data memoryTable
	mux  sync.RWMutex

	storage.DB
}

// NewMemoryStore create a new memory store
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{data: memoryTable{"settings": memoryRow{}}}
}

// Data return the raw internal memory state.
func (store *MemoryStore) Data() memoryTable {
	store.mux.RLock()
	data := store.data
	store.mux.RUnlock()
	return data
}

// DescribeAll return all element in the given table
func (store *MemoryStore) DescribeAll(table string) ([]interface{}, error) {
	store.mux.RLock()
	partition := store.data[table]
	store.mux.RUnlock()
	data := make([]interface{}, len(partition))

	if partition == nil {
		partition = memoryRow{}
	}

	i := 0
	for _, v := range partition {
		data[i] = v
		i++
	}

	return data, nil
}

// Describe return the elememt with the given ID in the given table
func (store *MemoryStore) Describe(table string, id string) (interface{}, error) {
	store.mux.RLock()
	partition := store.data[table]
	store.mux.RUnlock()

	if partition == nil {
		partition = memoryRow{}
	}

	elem := partition[id]

	if elem == nil {
		return nil, errors.New("record not exist")
	}

	return elem, nil
}

// Put upsert record in the given table
func (store *MemoryStore) Put(table string, id string, object interface{}) error {
	store.mux.Lock()
	partition := store.data[table]

	if partition == nil {
		partition = memoryRow{}
	}

	partition[id] = object

	store.data[table] = partition
	store.mux.Unlock()
	return nil
}

// Delete remove an element in the given table
func (store *MemoryStore) Delete(table string, id string) error {
	return nil
}
