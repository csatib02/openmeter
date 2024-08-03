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
	"context"
	"encoding/json"
	"time"

	"github.com/openmeterio/openmeter/internal/entitlement"
	"github.com/openmeterio/openmeter/internal/productcatalog"
	"github.com/openmeterio/openmeter/pkg/clock"
	"github.com/openmeterio/openmeter/pkg/recurrence"
)

type Connector interface {
	entitlement.SubTypeConnector
}

type connector struct {
	granularity time.Duration
}

func NewStaticEntitlementConnector() Connector {
	return &connector{
		granularity: time.Minute,
	}
}

func (c *connector) GetValue(entitlement *entitlement.Entitlement, at time.Time) (entitlement.EntitlementValue, error) {
	static, err := ParseFromGenericEntitlement(entitlement)
	if err != nil {
		return nil, err
	}

	return &StaticEntitlementValue{
		Config: static.Config,
	}, nil
}

func (c *connector) BeforeCreate(model entitlement.CreateEntitlementInputs, feature productcatalog.Feature) (*entitlement.CreateEntitlementRepoInputs, error) {
	model.EntitlementType = entitlement.EntitlementTypeStatic

	if model.MeasureUsageFrom != nil ||
		model.IssueAfterReset != nil ||
		model.IsSoftLimit != nil {
		return nil, &entitlement.InvalidValueError{Type: model.EntitlementType, Message: "Invalid inputs for type"}
	}

	// validate that config is JSON parseable
	if model.Config == nil {
		return nil, &entitlement.InvalidValueError{Type: model.EntitlementType, Message: "Config is required"}
	}

	if !json.Valid(model.Config) {
		return nil, &entitlement.InvalidValueError{Type: model.EntitlementType, Message: "Config is not valid JSON"}
	}

	var usagePeriod *entitlement.UsagePeriod
	var currentUsagePeriod *recurrence.Period

	if model.UsagePeriod != nil {
		usagePeriod = model.UsagePeriod

		calculatedPeriod, err := usagePeriod.GetCurrentPeriodAt(clock.Now())
		if err != nil {
			return nil, err
		}

		currentUsagePeriod = &calculatedPeriod
	}

	return &entitlement.CreateEntitlementRepoInputs{
		Namespace:          model.Namespace,
		FeatureID:          feature.ID,
		FeatureKey:         feature.Key,
		SubjectKey:         model.SubjectKey,
		EntitlementType:    model.EntitlementType,
		Metadata:           model.Metadata,
		UsagePeriod:        model.UsagePeriod,
		CurrentUsagePeriod: currentUsagePeriod,
		Config:             model.Config,
	}, nil
}

func (c *connector) AfterCreate(ctx context.Context, entitlement *entitlement.Entitlement) error {
	return nil
}

type StaticEntitlementValue struct {
	Config []byte `json:"config,omitempty"`
}

var _ entitlement.EntitlementValue = &StaticEntitlementValue{}

func (s *StaticEntitlementValue) HasAccess() bool {
	return true
}
