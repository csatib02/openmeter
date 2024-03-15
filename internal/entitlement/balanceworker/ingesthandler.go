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

package balanceworker

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"

	"github.com/openmeterio/openmeter/internal/entitlement"
	"github.com/openmeterio/openmeter/internal/event/spec"
	"github.com/openmeterio/openmeter/internal/productcatalog"
	"github.com/openmeterio/openmeter/internal/sink/flushhandler/ingestnotification"
)

func (w *Worker) handleIngestEvent(ctx context.Context, event ingestnotification.EventIngested) ([]*message.Message, error) {
	affectedEntitlements, err := w.GetEntitlementsAffectedByMeterSubject(ctx, event.Namespace.ID, event.MeterSlugs, event.SubjectKey)
	if err != nil {
		return nil, err
	}

	result := make([]*message.Message, 0, len(affectedEntitlements))
	for _, entitlement := range affectedEntitlements {
		messages, err := w.handleEntitlementUpdateEvent(
			ctx,
			entitlement,
			spec.ComposeResourcePath(entitlement.Namespace, spec.EntityEvent),
		)
		if err != nil {
			return nil, err
		}

		result = append(result, messages...)
	}

	return result, nil
}

func (w *Worker) GetEntitlementsAffectedByMeterSubject(ctx context.Context, namespace string, meterSlugs []string, subject string) ([]NamespacedID, error) {
	featuresByMeter, err := w.connectors.Feature.ListFeatures(ctx, productcatalog.ListFeaturesParams{
		Namespace:  namespace,
		MeterSlugs: meterSlugs,
	})
	if err != nil {
		return nil, err
	}

	featureIDs := make([]string, 0, len(featuresByMeter.Items))
	for _, feature := range featuresByMeter.Items {
		featureIDs = append(featureIDs, feature.ID)
	}

	entitlements, err := w.connectors.Entitlement.ListEntitlements(ctx, entitlement.ListEntitlementsParams{
		Namespaces:  []string{namespace},
		SubjectKeys: []string{subject},
		FeatureIDs:  featureIDs,
	})
	if err != nil {
		return nil, err
	}

	entitlementIDs := make([]NamespacedID, 0, len(entitlements.Items))
	for _, entitlement := range entitlements.Items {
		entitlementIDs = append(entitlementIDs, NamespacedID{
			ID:        entitlement.ID,
			Namespace: entitlement.Namespace,
		})
	}

	return entitlementIDs, nil
}
