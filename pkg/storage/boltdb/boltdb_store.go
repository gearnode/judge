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

// Package boltdbstore contains an in memory implementation of
// the Judge storage interface.
package boltdbstore // import "github.com/gearnode/judge/pkg/storage/boltdb"

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/boltdb/bolt"
)

type DB struct {
	store *bolt.DB
}

func NewBoltdbStore() (*DB, error) {
	store, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		return &DB{}, err
	}

	err = store.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("policies"))
		if err != nil {
			return fmt.Errorf("could not create policies bucket: %v", err)
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("could not set up buckets, %v", err)
	}
	return &DB{store: store}, nil
}

func (db *DB) Describe(table string, id string) (interface{}, error) {
	var val bytes.Buffer
	dec := gob.NewDecoder(&val)

	err := db.store.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(table))
		v := b.Get([]byte(id))

		err := dec.Decode(v)
		if err != nil {
			return err
		}

		return nil
	})

	return val, err
}

func (db *DB) Put(table string, id string, data interface{}) error {
	var val bytes.Buffer
	enc := gob.NewEncoder(&val)

	err := db.store.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(table))
		enc.Encode(data)
		err := b.Put([]byte(id), val.Bytes())
		return err
	})
	return err
}

func (db *DB) DescribeAll(table string) ([]interface{}, error) {
	var val []interface{}

	db.store.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(table))

		b.ForEach(func(k, v []byte) error {
			var x bytes.Buffer
			dec := gob.NewDecoder(&x)
			val = append(val, dec.Decode(v))
			return nil
		})
		return nil
	})

	return val, nil
}
