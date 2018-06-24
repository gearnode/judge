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

// Package storage contains code for storages (like the redis
// storage).
package storage // import "github.com/gearnode/judge/pkg/storage"

// DB represents the database generic interface.
type DB interface {
	DescribeAll(table string) ([]interface{}, error)
	Describe(table string, id string) (interface{}, error)
	Put(table string, id string, object interface{}) error
	Delete(table string, id string) error
}
