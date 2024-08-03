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

package snapshot

import (
	"errors"
	"time"

	"github.com/openmeterio/openmeter/internal/entitlement"
	"github.com/openmeterio/openmeter/internal/event/models"
	"github.com/openmeterio/openmeter/internal/event/spec"
	"github.com/openmeterio/openmeter/internal/productcatalog"
	"github.com/openmeterio/openmeter/pkg/recurrence"
)

const (
	snapshotEventName spec.EventName = "entitlement.snapshot"
)

type BalanceOperationType string

const (
	BalanceOperationUpdate BalanceOperationType = "update"
	BalanceOperationDelete BalanceOperationType = "delete"
)

type EntitlementValue struct {
	// Balance Only available for metered entitlements. Metered entitlements are built around a balance calculation where feature usage is deducted from the issued grants. Balance represents the remaining balance of the entitlement, it's value never turns negative.
	Balance *float64 `json:"balance,omitempty"`

	// Config Only available for static entitlements. The JSON parsable config of the entitlement.
	Config *string `json:"config,omitempty"`

	// HasAccess Whether the subject has access to the feature. Shared across all entitlement types.
	HasAccess *bool `json:"hasAccess,omitempty"`

	// Overage Only available for metered entitlements. Overage represents the usage that wasn't covered by grants, e.g. if the subject had a total feature usage of 100 in the period but they were only granted 80, there would be 20 overage.
	Overage *float64 `json:"overage,omitempty"`

	// Usage Only available for metered entitlements. Returns the total feature usage in the current period.
	Usage *float64 `json:"usage,omitempty"`
}

type SnapshotEvent struct {
	Entitlement entitlement.Entitlement `json:"entitlement"`
	Namespace   models.NamespaceID      `json:"namespace"`
	Subject     models.SubjectKeyAndID  `json:"subject"`
	Feature     productcatalog.Feature  `json:"feature"`
	// Operation is delete if the entitlement gets deleted, in that case the balance object is empty
	Operation BalanceOperationType `json:"operation"`

	// CalculatedAt specifies when the balance calculation was performed. It can be used to verify
	// in edge-worker if the store already contains the required item.
	CalculatedAt *time.Time `json:"calculatedAt,omitempty"`

	Balance            *EntitlementValue  `json:"balance,omitempty"`
	CurrentUsagePeriod *recurrence.Period `json:"currentUsagePeriod,omitempty"`
}

var SnapshotEventSpec = spec.EventTypeSpec{
	Subsystem: entitlement.EventSubsystem,
	Name:      snapshotEventName,
	Version:   "v1",
}

func (e SnapshotEvent) Spec() *spec.EventTypeSpec {
	return &SnapshotEventSpec
}

func (e SnapshotEvent) Validate() error {
	if e.Operation != BalanceOperationDelete && e.Operation != BalanceOperationUpdate {
		return errors.New("operation must be either delete or update")
	}

	if e.Entitlement.ID == "" {
		return errors.New("entitlementId is required")
	}

	if err := e.Namespace.Validate(); err != nil {
		return err
	}

	if err := e.Subject.Validate(); err != nil {
		return err
	}

	if e.Feature.ID == "" {
		return errors.New("feature ID must be set")
	}

	if e.Operation == BalanceOperationUpdate {
		if e.CalculatedAt == nil {
			return errors.New("calculatedAt is required for balance update")
		}

		if e.Balance == nil {
			return errors.New("balance is required for balance update")
		}

		if e.CurrentUsagePeriod == nil {
			return errors.New("currentUsagePeriod is required for balance update")
		}
	}

	return nil
}
