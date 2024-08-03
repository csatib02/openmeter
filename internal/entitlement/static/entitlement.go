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

package staticentitlement

import (
	"github.com/openmeterio/openmeter/internal/entitlement"
)

type Entitlement struct {
	entitlement.GenericProperties

	Config []byte `json:"config,omitempty"`
}

func ParseFromGenericEntitlement(model *entitlement.Entitlement) (*Entitlement, error) {
	if model.EntitlementType != entitlement.EntitlementTypeStatic {
		return nil, &entitlement.WrongTypeError{Expected: entitlement.EntitlementTypeStatic, Actual: model.EntitlementType}
	}

	if model.Config == nil {
		return nil, &entitlement.InvalidValueError{Type: model.EntitlementType, Message: "Config is required"}
	}

	return &Entitlement{
		GenericProperties: model.GenericProperties,
		Config:            model.Config,
	}, nil
}
