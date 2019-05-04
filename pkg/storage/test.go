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

package storage

import (
	"github.com/gearnode/judge/pkg/policy"
	"github.com/stretchr/testify/assert"
	"testing"
)

// StorageTestSuite define the test suite to ensure storage implementation
// is valid and compliant with the storage specification.
func StorageTestSuite(t *testing.T, store Storage) {
	assert := assert.New(t)

	var err error

	_, err = store.GetPolicy("notExistingID")
	assert.Error(err)

	err = store.DelPolicy("notExistingID")
	assert.Error(err)

	expected, err := policy.NewPolicy("test pol", "")
	assert.NoError(err)

	tested, err := store.PutPolicy(expected)
	assert.NoError(err)
	assert.Equal(expected.ID.String(), tested.ID.String())
	assert.Equal(expected.Name, tested.Name)
	assert.Equal(expected.Description, tested.Description)
	assert.Len(tested.Statements, len(expected.Statements))

	expected.Name = "a new name"
	tested, err = store.PutPolicy(expected)
	assert.NoError(err)
	assert.Equal(expected.ID.String(), tested.ID.String())
	assert.Equal(expected.Name, tested.Name)
	assert.Equal(expected.Description, tested.Description)
	assert.Len(tested.Statements, len(expected.Statements))

	tested, err = store.GetPolicy(expected.ID.String())
	assert.NoError(err)
	assert.Equal(expected.ID.String(), tested.ID.String())
	assert.Equal(expected.Name, tested.Name)
	assert.Equal(expected.Description, tested.Description)
	assert.Len(tested.Statements, len(expected.Statements))

	err = store.DelPolicy(expected.ID.String())
	assert.NoError(err)

	err = store.DelPolicy(expected.ID.String())
	assert.Error(err)

	_, err = store.GetPolicy(expected.ID.String())
	assert.Error(err)
}
