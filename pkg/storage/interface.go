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
)

// ScrollOptions contains pagination and order options for scroll.
type ScrollOptions struct {
	PageToken string
	PageSize  int
	OrderBy   *ScrollOrderOptions
}

// ScrollOrderOptions contains order options for scroll.
type ScrollOrderOptions struct {
	Column string
	Order  string
}

// Storage define storage interface without any implementation.
type Storage interface {
	GetPolicy(id string) (*policy.Policy, error)
	PutPolicy(pol *policy.Policy) (*policy.Policy, error)
	DelPolicy(id string) error
	ScrollPolicy(opts *ScrollOptions) ([]*policy.Policy, string, error)
}
