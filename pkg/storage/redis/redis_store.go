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

package redisstore

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/gearnode/judge/pkg/policy"
	"github.com/gearnode/judge/pkg/storage"
	"github.com/go-redis/redis"
)

// RedisStore is  a simple redis storage implementing the Storage interface.
type RedisStore struct {
	conn *redis.Client
	storage.Storage
}

// Options is a alias type for https://godoc.org/github.com/go-redis/redis#Options
type Options = redis.Options

// NewRedisStore create a new redis store
func NewRedisStore(cfg *Options) (*RedisStore, error) {
	client := redis.NewClient(cfg)

	pong, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	if pong != "PONG" {
		return nil, errors.New("failed to ping redis")
	}

	return &RedisStore{conn: client}, nil
}

// GetPolicy returns the policy with the given identifier.
func (s *RedisStore) GetPolicy(id string) (*policy.Policy, error) {
	val, err := s.conn.Get("pol:" + id).Result()
	if err != nil {
		return nil, err
	}

	if val == "" {
		return nil, fmt.Errorf("the policy with %q ID does not exist", id)
	}

	var pol policy.Policy

	data := bytes.NewBufferString(val)
	gob.NewDecoder(data).Decode(&pol)

	return &pol, nil
}

// PutPolicy upsert the policy with the given identifier.
func (s *RedisStore) PutPolicy(pol *policy.Policy) (*policy.Policy, error) {
	var data bytes.Buffer
	err := gob.NewEncoder(&data).Encode(&pol)

	if err != nil {
		return nil, err
	}

	val, err := s.conn.Set("pol:"+pol.ID.String(), data.String(), 0).Result()

	if err != nil || val != "OK" {
		return nil, err
	}

	return pol, nil
}

// DelPolicy delete the policy with the given identifier.
func (s *RedisStore) DelPolicy(id string) error {

	val, err := s.conn.Del("pol:" + id).Result()

	if err != nil {
		return err
	}

	if val != 1 {
		return fmt.Errorf("the policy with %q ID does not exist", id)
	}

	return nil
}

func (s *RedisStore) ScrollPolicy(opts *storage.ScrollOptions) ([]*policy.Policy, string, error) {
	return nil, "", errors.New("not implemented yet")
}
