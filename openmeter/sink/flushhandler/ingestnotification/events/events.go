package events

import (
	"errors"

	"github.com/openmeterio/openmeter/openmeter/event/metadata"
	"github.com/openmeterio/openmeter/openmeter/event/models"
	"github.com/openmeterio/openmeter/openmeter/watermill/marshaler"
)

const (
	EventSubsystem metadata.EventSubsystem = "ingest"
)

type EventBatchedIngest struct {
	Events []IngestEventData `json:"events"`
}

var (
	_ marshaler.Event = EventBatchedIngest{}

	batchIngestEventType = metadata.EventType{
		Subsystem: EventSubsystem,
		Name:      "events.ingested",
		Version:   "v1",
	}
	batchIngestEventName  = metadata.GetEventName(batchIngestEventType)
	EventVersionSubsystem = batchIngestEventType.VersionSubsystem()
)

func (b EventBatchedIngest) EventName() string {
	return batchIngestEventName
}

func (b EventBatchedIngest) Validate() error {
	if len(b.Events) == 0 {
		return errors.New("events must not be empty")
	}

	var finalErr error

	for _, e := range b.Events {
		if err := e.Validate(); err != nil {
			finalErr = errors.Join(finalErr, err)
		}
	}

	return finalErr
}

func (b EventBatchedIngest) EventMetadata() metadata.EventMetadata {
	return metadata.EventMetadata{
		Source: metadata.ComposeResourcePathRaw(string(EventSubsystem)),
	}
}

type IngestEventData struct {
	Namespace  models.NamespaceID `json:"namespace"`
	SubjectKey string             `json:"subjectKey"`

	// MeterSlugs contain the list of slugs that are affected by the event. We
	// should not use meterIDs as they are not something present in the open source
	// version, thus any code that is in opensource should not rely on them.
	MeterSlugs []string `json:"meterSlugs"`
}

func (i IngestEventData) Validate() error {
	if err := i.Namespace.Validate(); err != nil {
		return err
	}

	if i.SubjectKey == "" {
		return errors.New("subjectKey must be set")
	}

	if len(i.MeterSlugs) == 0 {
		return errors.New("meterSlugs must not be empty")
	}

	return nil
}
