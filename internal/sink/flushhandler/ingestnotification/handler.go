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

package ingestnotification

import (
	"context"
	"errors"
	"log/slog"

	"go.opentelemetry.io/otel/metric"

	eventmodels "github.com/openmeterio/openmeter/internal/event/models"
	"github.com/openmeterio/openmeter/internal/event/spec"
	"github.com/openmeterio/openmeter/internal/sink/flushhandler"
	sinkmodels "github.com/openmeterio/openmeter/internal/sink/models"
	"github.com/openmeterio/openmeter/openmeter/event/publisher"
	"github.com/openmeterio/openmeter/pkg/models"
)

type handler struct {
	publisher publisher.TopicPublisher
	logger    *slog.Logger
}

func NewHandler(logger *slog.Logger, metricMeter metric.Meter, publisher publisher.TopicPublisher) (flushhandler.FlushEventHandler, error) {
	handler := &handler{
		publisher: publisher,
		logger:    logger,
	}

	return flushhandler.NewFlushEventHandler(
		flushhandler.FlushEventHandlerOptions{
			Name:        "ingest_notification",
			Callback:    handler.OnFlushSuccess,
			Logger:      logger,
			MetricMeter: metricMeter,
		})
}

// OnFlushSuccess takes a look at the incoming messages and in case something is
// affecting a ledger balance it will create the relevant event.
func (h *handler) OnFlushSuccess(ctx context.Context, events []sinkmodels.SinkMessage) error {
	var finalErr error

	for _, message := range events {
		if message.Serialized == nil {
			// In case the incoming event was not supported by the parser (e.g. non-json payload)
			continue
		}

		if len(message.Meters) == 0 {
			// If the change doesn't affect a meter, we should not care about it
			continue
		}

		event, err := spec.NewCloudEvent(spec.EventSpec{
			ID:      message.Serialized.Id,
			Source:  spec.ComposeResourcePath(message.Namespace, spec.EntityEvent, message.Serialized.Id),
			Subject: spec.ComposeResourcePath(message.Namespace, spec.EntitySubjectKey, message.Serialized.Subject),
		}, EventIngested{
			Namespace:  eventmodels.NamespaceID{ID: message.Namespace},
			SubjectKey: message.Serialized.Subject,
			MeterSlugs: h.getMeterSlugsFromMeters(message.Meters),
		})
		if err != nil {
			finalErr = errors.Join(finalErr, err)
			h.logger.Error("failed to create change notification", "error", err)
			continue
		}

		if err := h.publisher.Publish(event); err != nil {
			finalErr = errors.Join(finalErr, err)
			h.logger.Error("failed to publish change notification", "error", err)
		}
	}

	return finalErr
}

func (h *handler) getMeterSlugsFromMeters(meters []models.Meter) []string {
	slugs := make([]string, len(meters))
	for i, meter := range meters {
		slugs[i] = meter.Slug
	}

	return slugs
}
