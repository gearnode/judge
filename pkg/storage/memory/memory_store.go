/*
Copyright 2019 Bryan Frimin.

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
	"fmt"
	"sync"

	"github.com/gearnode/judge/pkg/policy"
	"github.com/gearnode/judge/pkg/storage"
)

// MemoryStore is a simple in-memory storage implementing the Storage interface.
type MemoryStore struct {
	mux sync.RWMutex
	storage.Storage
	policies map[string]*policy.Policy
}

// NewMemoryStore create a new memory store
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{policies: make(map[string]*policy.Policy)}
}

// GetPolicy returns the policy with the given identifier.
func (s *MemoryStore) GetPolicy(id string) (*policy.Policy, error) {
	s.mux.RLock()
	defer s.mux.RUnlock()

	if s.policies[id] != nil {
		return s.policies[id], nil
	}

	return nil, fmt.Errorf("policy with id %q does not exist", id)
}

func (s *MemoryStore) PutPolicy(pol *policy.Policy) (*policy.Policy, error) {
	s.mux.Lock()
	defer s.mux.Unlock()

	s.policies[pol.ID.String()] = pol

	return s.policies[pol.ID.String()], nil
}

func (s *MemoryStore) DelPolicy(id string) error {
	s.mux.Lock()
	defer s.mux.Unlock()

	if s.policies[id] != nil {
		delete(s.policies, id)
		return nil
	}

	return fmt.Errorf("policy with id %q does not exist", id)
}

func (s *MemoryStore) ScrollPolicy(opts *storage.ScrollOptions) ([]*policy.Policy, string, error) {
	return []*policy.Policy{}, "", nil
}
