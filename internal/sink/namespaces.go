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

package sink

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/oliveagle/jsonpath"

	"github.com/openmeterio/openmeter/internal/ingest/kafkaingest/serializer"
	"github.com/openmeterio/openmeter/pkg/models"
)

type MeterStore struct {
	Meters []models.Meter
}

type NamespaceStore struct {
	namespaces map[string]*MeterStore
}

func NewNamespaceStore() *NamespaceStore {
	return &NamespaceStore{
		namespaces: make(map[string]*MeterStore),
	}
}

func (n *NamespaceStore) AddMeter(meter models.Meter) {
	if n.namespaces[meter.Namespace] == nil {
		n.namespaces[meter.Namespace] = &MeterStore{
			Meters: []models.Meter{meter},
		}
	} else {
		n.namespaces[meter.Namespace].Meters = append(n.namespaces[meter.Namespace].Meters, meter)
	}
}

// ValidateEvent validates a single event by matching it with the corresponding meter if any
func (a *NamespaceStore) ValidateEvent(ctx context.Context, event serializer.CloudEventsKafkaPayload, namespace string) error {
	namespaceStore := a.namespaces[namespace]
	if namespaceStore == nil {
		// We drop events from unknown org
		return NewProcessingError(fmt.Sprintf("namespace not found: %s", namespace), DROP)
	}

	// Validate a single event against multiple meters
	var foundMeter bool
	for _, meter := range namespaceStore.Meters {
		if meter.EventType == event.Type {
			foundMeter = true
			err := validateEventWithMeter(meter, event)
			if err != nil {
				return err
			}
			// A single event can match multiple meters so we cannot break the loop early
		}
	}

	if !foundMeter {
		// Mark as invalid so we can show it to the user
		return NewProcessingError(fmt.Sprintf("no meter found for event type: %s", event.Type), INVALID)
	}

	return nil
}

// validateEventWithMeter validates a single event against a single meter
func validateEventWithMeter(meter models.Meter, ev serializer.CloudEventsKafkaPayload) *ProcessingError {
	// Parse CloudEvents data as JSON, currently we only support JSON encoding
	var data interface{}
	err := json.Unmarshal([]byte(ev.Data), &data)
	if err != nil {
		return NewProcessingError("cannot unmarshal event data as json", INVALID)
	}

	// We can skip count events as they don't have value property
	if meter.Aggregation == models.MeterAggregationCount {
		return nil
	}

	// Parse and validate value
	valueRaw, err := jsonpath.JsonPathLookup(data, meter.ValueProperty)
	if err != nil {
		return NewProcessingError(fmt.Sprintf("event data is missing value property at %s", meter.ValueProperty), INVALID)
	}
	if valueRaw == nil {
		return NewProcessingError("event data value cannot be null", INVALID)
	}

	if valueStr, ok := valueRaw.(string); ok {
		_, err = strconv.ParseFloat(valueStr, 64)
		if err != nil {
			return NewProcessingError(fmt.Sprintf("event data value cannot be parsed as float64: %s", valueStr), INVALID)
		}
	} else if _, ok := valueRaw.(float64); ok {

	} else {
		return NewProcessingError("event data value property cannot be parsed", INVALID)
	}

	return nil
}
