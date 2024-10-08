package kafka

import (
	"context"
	"errors"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type TopicConfig struct {
	Name          string
	Partitions    int
	Replicas      int
	RetentionTime TimeDurationMilliSeconds
}

func (c TopicConfig) Validate() error {
	if c.Name == "" {
		return errors.New("topic name must not be empty")
	}

	if c.Partitions <= 0 {
		return errors.New("number of partitions must be greater than zero")
	}

	return nil
}

func ProvisionTopics(ctx context.Context, client *kafka.AdminClient, topics ...TopicConfig) error {
	if len(topics) == 0 {
		return nil
	}

	topicSpecs := make([]kafka.TopicSpecification, 0, len(topics))

	for _, topic := range topics {
		if err := topic.Validate(); err != nil {
			return fmt.Errorf("invalid topic configuration: %w", err)
		}

		topicSpec := kafka.TopicSpecification{
			Topic:         topic.Name,
			NumPartitions: topic.Partitions,
		}

		if topic.Replicas > 0 {
			topicSpec.ReplicationFactor = topic.Replicas
		}

		if topic.RetentionTime > 0 {
			topicSpec.Config["retention.ms"] = topic.RetentionTime.String()
		}

		topicSpecs = append(topicSpecs, topicSpec)
	}

	results, err := client.CreateTopics(ctx, topicSpecs)
	if err != nil {
		return fmt.Errorf("failed to create topics: %w", err)
	}

	var errs []error
	for _, result := range results {
		switch result.Error.Code() {
		case kafka.ErrNoError, kafka.ErrTopicAlreadyExists:
		default:
			errs = append(errs, fmt.Errorf("failed to create topic: %w", result.Error))
		}
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}
