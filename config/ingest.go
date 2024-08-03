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

package config

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/spf13/viper"

	pkgkafka "github.com/openmeterio/openmeter/pkg/kafka"
)

type IngestConfiguration struct {
	Kafka KafkaIngestConfiguration
}

// Validate validates the configuration.
func (c IngestConfiguration) Validate() error {
	if err := c.Kafka.Validate(); err != nil {
		return fmt.Errorf("kafka: %w", err)
	}

	return nil
}

type KafkaIngestConfiguration struct {
	KafkaConfiguration `mapstructure:",squash"`

	Partitions          int
	EventsTopicTemplate string
}

// Validate validates the configuration.
func (c KafkaIngestConfiguration) Validate() error {
	if c.EventsTopicTemplate == "" {
		return errors.New("events topic template is required")
	}

	if err := c.KafkaConfiguration.Validate(); err != nil {
		return err
	}
	return nil
}

type KafkaConfiguration struct {
	Broker           string
	SecurityProtocol string
	SaslMechanisms   string
	SaslUsername     string
	SaslPassword     string

	StatsInterval pkgkafka.TimeDurationMilliSeconds

	// BrokerAddressFamily defines the IP address family to be used for network communication with Kafka cluster
	BrokerAddressFamily pkgkafka.BrokerAddressFamily
	// SocketKeepAliveEnable defines if TCP socket keep-alive is enabled to prevent closing idle connections
	// by Kafka brokers.
	SocketKeepAliveEnabled bool
	// TopicMetadataRefreshInterval defines how frequently the Kafka client needs to fetch metadata information
	// (brokers, topic, partitions, etc) from the Kafka cluster.
	// The 5 minutes default value is appropriate for mostly static Kafka clusters, but needs to be lowered
	// in case of large clusters where changes are more frequent.
	// This value must not be set to value lower than 10s.
	TopicMetadataRefreshInterval pkgkafka.TimeDurationMilliSeconds

	// Enable contexts for extensive debugging of librdkafka.
	// See: https://github.com/confluentinc/librdkafka/blob/master/INTRODUCTION.md#debug-contexts
	DebugContexts pkgkafka.DebugContexts
}

func (c KafkaConfiguration) Validate() error {
	if c.Broker == "" {
		return errors.New("broker is required")
	}

	if c.StatsInterval > 0 && c.StatsInterval.Duration() < 5*time.Second {
		return errors.New("StatsInterval must be >=5s")
	}

	if c.TopicMetadataRefreshInterval > 0 && c.TopicMetadataRefreshInterval.Duration() < 10*time.Second {
		return errors.New("topic metadata refresh interval must be >=10s")
	}

	return nil
}

// CreateKafkaConfig creates a Kafka config map.
func (c KafkaConfiguration) CreateKafkaConfig() kafka.ConfigMap {
	config := kafka.ConfigMap{
		"bootstrap.servers": c.Broker,

		// Required for logging
		"go.logs.channel.enable": true,
	}

	// This is needed when using localhost brokers on OSX,
	// since the OSX resolver will return the IPv6 addresses first.
	// See: https://github.com/openmeterio/openmeter/issues/321
	if c.BrokerAddressFamily != "" {
		config["broker.address.family"] = c.BrokerAddressFamily
	} else if strings.Contains(c.Broker, "localhost") || strings.Contains(c.Broker, "127.0.0.1") {
		config["broker.address.family"] = pkgkafka.BrokerAddressFamilyIPv4
	}

	if c.SecurityProtocol != "" {
		config["security.protocol"] = c.SecurityProtocol
	}

	if c.SaslMechanisms != "" {
		config["sasl.mechanism"] = c.SaslMechanisms
	}

	if c.SaslUsername != "" {
		config["sasl.username"] = c.SaslUsername
	}

	if c.SaslPassword != "" {
		config["sasl.password"] = c.SaslPassword
	}

	if c.StatsInterval > 0 {
		config["statistics.interval.ms"] = c.StatsInterval
	}

	if c.SocketKeepAliveEnabled {
		config["socket.keepalive.enable"] = c.SocketKeepAliveEnabled
	}

	// The `topic.metadata.refresh.interval.ms` defines the frequency the Kafka client needs to retrieve metadata
	// from Kafka cluster. While `metadata.max.age.ms` defines the interval after the metadata cache maintained
	// on client side becomes invalid. Setting the former will automatically adjust the value of the latter to avoid
	// misconfiguration where the entries in metadata cache are evicted prior metadata refresh.
	if c.TopicMetadataRefreshInterval > 0 {
		config["topic.metadata.refresh.interval.ms"] = c.TopicMetadataRefreshInterval
		config["metadata.max.age.ms"] = 3 * c.TopicMetadataRefreshInterval
	}

	if len(c.DebugContexts) > 0 {
		config["debug"] = c.DebugContexts.String()
	}

	return config
}

// Configure configures some defaults in the Viper instance.
func ConfigureIngest(v *viper.Viper) {
	v.SetDefault("ingest.kafka.broker", "127.0.0.1:29092")
	v.SetDefault("ingest.kafka.securityProtocol", "")
	v.SetDefault("ingest.kafka.saslMechanisms", "")
	v.SetDefault("ingest.kafka.saslUsername", "")
	v.SetDefault("ingest.kafka.saslPassword", "")
	v.SetDefault("ingest.kafka.partitions", 1)
	v.SetDefault("ingest.kafka.eventsTopicTemplate", "om_%s_events")
}
