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
	"time"

	"github.com/spf13/viper"
)

type EventsConfiguration struct {
	Enabled      bool
	SystemEvents EventSubsystemConfiguration
	IngestEvents EventSubsystemConfiguration
}

func (c EventsConfiguration) Validate() error {
	return c.SystemEvents.Validate()
}

type EventSubsystemConfiguration struct {
	Enabled bool
	Topic   string

	AutoProvision AutoProvisionConfiguration
}

func (c EventSubsystemConfiguration) Validate() error {
	if !c.Enabled {
		return nil
	}

	if c.Topic == "" {
		return errors.New("topic name is required")
	}
	return c.AutoProvision.Validate()
}

type AutoProvisionConfiguration struct {
	Enabled    bool
	Partitions int
}

func (c AutoProvisionConfiguration) Validate() error {
	if c.Enabled && c.Partitions < 1 {
		return errors.New("partitions must be greater than 0")
	}
	return nil
}

type DLQConfiguration struct {
	Enabled       bool
	Topic         string
	AutoProvision AutoProvisionConfiguration
	Throttle      ThrottleConfiguration
}

func (c DLQConfiguration) Validate() error {
	if !c.Enabled {
		return nil
	}

	if c.Topic == "" {
		return errors.New("topic name is required")
	}

	if err := c.Throttle.Validate(); err != nil {
		return fmt.Errorf("throttle: %w", err)
	}

	return nil
}

type ThrottleConfiguration struct {
	Enabled  bool
	Count    int64
	Duration time.Duration
}

func (c ThrottleConfiguration) Validate() error {
	if !c.Enabled {
		return nil
	}

	if c.Count <= 0 {
		return errors.New("count must be greater than 0")
	}

	if c.Duration <= 0 {
		return errors.New("duration must be greater than 0")
	}

	return nil
}

type RetryConfiguration struct {
	MaxRetries      int
	InitialInterval time.Duration
}

func (c RetryConfiguration) Validate() error {
	if c.MaxRetries <= 0 {
		return errors.New("max retries must be greater than 0")
	}

	if c.InitialInterval <= 0 {
		return errors.New("initial interval must be greater than 0")
	}

	return nil
}

func ConfigureEvents(v *viper.Viper) {
	// TODO: after the system events are fully implemented, we should enable them by default
	v.SetDefault("events.enabled", false)
	v.SetDefault("events.systemEvents.enabled", true)
	v.SetDefault("events.systemEvents.topic", "om_sys.api_events")
	v.SetDefault("events.systemEvents.autoProvision.enabled", true)
	v.SetDefault("events.systemEvents.autoProvision.partitions", 4)

	v.SetDefault("events.ingestEvents.enabled", true)
	v.SetDefault("events.ingestEvents.topic", "om_sys.ingest_events")
	v.SetDefault("events.ingestEvents.autoProvision.enabled", true)
	v.SetDefault("events.ingestEvents.autoProvision.partitions", 8)
}
