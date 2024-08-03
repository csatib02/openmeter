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

package kafka

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"

	watermillkafka "github.com/openmeterio/openmeter/internal/watermill/driver/kafka"
)

const (
	PartitionKeyMetadataKey = watermillkafka.PartitionKeyMetadataKey
)

type (
	Publisher = watermillkafka.Publisher
)

func NewPublisher(producer *kafka.Producer) *Publisher {
	return watermillkafka.NewPublisher(producer)
}

func AddPartitionKeyFromSubject(watermillIn *message.Message, cloudEvent event.Event) (*message.Message, error) {
	return watermillkafka.AddPartitionKeyFromSubject(watermillIn, cloudEvent)
}
