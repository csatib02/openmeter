// Copyright © 2024 Tailfin Cloud Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/openmeterio/openmeter/pkg/framework/entutils"
)

type Example1 struct {
	ent.Schema
}

func (Example1) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entutils.TimeMixin{},
	}
}

func (Example1) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Unique().
			Immutable(),
		field.String("example_value_1"),
	}
}

func (Example1) Indexes() []ent.Index {
	return []ent.Index{}
}

func (Example1) Edges() []ent.Edge {
	return []ent.Edge{}
}
