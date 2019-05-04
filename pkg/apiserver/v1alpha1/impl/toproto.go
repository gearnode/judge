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

package apiserver

import (
	pb "github.com/gearnode/judge/pkg/apiserver/v1alpha1"
	"github.com/gearnode/judge/pkg/policy"
	"github.com/gearnode/judge/pkg/policy/resource"
)

// toProto serialize the Policy go struct in the v1alpha1 policy proto struct.
func toProto(pol *policy.Policy) *pb.Policy {
	proto := pb.Policy{
		PolicyId:    pol.ID.String(),
		Name:        pol.Name,
		Description: pol.Description,
		Statements:  make([]*pb.Statement, len(pol.Statements)),
	}

	for i := range pol.Statements {
		statement := pb.Statement{
			Effect:    pb.Statement_Effect(pb.Statement_Effect_value[pol.Statements[i].Effect]),
			Actions:   pol.Statements[i].Actions,
			Resources: make([]string, len(pol.Statements[i].Resources)),
		}

		for j, v := range pol.Statements[i].Resources {
			statement.Resources[j] = resource.Marshal(&v)
		}

		proto.Statements[i] = &statement
	}

	return &proto
}
