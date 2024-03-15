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

// Package config loads application configuration.
package config

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/openmeterio/openmeter/pkg/models"
)

// Configuration holds any kind of Configuration that comes from the outside world and
// is necessary for running the application.
type Configuration struct {
	Address     string
	Environment string

	Telemetry TelemetryConfig

	Namespace   NamespaceConfiguration
	Ingest      IngestConfiguration
	Aggregation AggregationConfiguration
	Sink        SinkConfiguration
	Dedupe      DedupeConfiguration
	Portal      PortalConfiguration

	Meters []*models.Meter
}

// Validate validates the configuration.
func (c Configuration) Validate() error {
	if c.Address == "" {
		return errors.New("server address is required")
	}

	if err := c.Telemetry.Validate(); err != nil {
		return fmt.Errorf("telemetry: %w", err)
	}

	if err := c.Namespace.Validate(); err != nil {
		return fmt.Errorf("namespace: %w", err)
	}

	if err := c.Ingest.Validate(); err != nil {
		return fmt.Errorf("ingest: %w", err)
	}

	if err := c.Aggregation.Validate(); err != nil {
		return fmt.Errorf("aggregation: %w", err)
	}

	if err := c.Sink.Validate(); err != nil {
		return fmt.Errorf("sink: %w", err)
	}

	if err := c.Dedupe.Validate(); err != nil {
		return fmt.Errorf("dedupe: %w", err)
	}

	if err := c.Portal.Validate(); err != nil {
		return fmt.Errorf("portal: %w", err)
	}

	for _, m := range c.Meters {
		// Namespace is not configurable on per meter level
		m.Namespace = c.Namespace.Default

		// set default window size
		if m.WindowSize == "" {
			m.WindowSize = models.WindowSizeMinute
		}

		if err := m.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// Configure configures some defaults in the Viper instance.
func Configure(v *viper.Viper, flags *pflag.FlagSet) {
	// Viper settings
	v.AddConfigPath(".")

	// Environment variable settings
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.AllowEmptyEnv(true)
	v.AutomaticEnv()

	// Server configuration
	flags.String("address", ":8888", "Server address")
	_ = v.BindPFlag("address", flags.Lookup("address"))
	v.SetDefault("address", ":8888")

	// Environment used for identifying the service environment
	v.SetDefault("environment", "unknown")

	configureTelemetry(v, flags)

	ConfigureNamespace(v)
	ConfigureIngest(v)
	ConfigureAggregation(v)
	ConfigureSink(v)
	ConfigureDedupe(v)
	ConfigurePortal(v)
}
