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

package registry

import (
	"log/slog"
	"time"

	"github.com/openmeterio/openmeter/internal/ent/db"
	"github.com/openmeterio/openmeter/internal/meter"
	"github.com/openmeterio/openmeter/openmeter/credit"
	creditpgadapter "github.com/openmeterio/openmeter/openmeter/credit/postgresadapter"
	"github.com/openmeterio/openmeter/openmeter/entitlement"
	booleanentitlement "github.com/openmeterio/openmeter/openmeter/entitlement/boolean"
	meteredentitlement "github.com/openmeterio/openmeter/openmeter/entitlement/metered"
	entitlementpgadapter "github.com/openmeterio/openmeter/openmeter/entitlement/postgresadapter"
	staticentitlement "github.com/openmeterio/openmeter/openmeter/entitlement/static"
	"github.com/openmeterio/openmeter/openmeter/event/publisher"
	"github.com/openmeterio/openmeter/openmeter/productcatalog"
	productcatalogpgadapter "github.com/openmeterio/openmeter/openmeter/productcatalog/postgresadapter"
	"github.com/openmeterio/openmeter/openmeter/streaming"
)

type Entitlement struct {
	Feature            productcatalog.FeatureConnector
	EntitlementOwner   credit.OwnerConnector
	CreditBalance      credit.BalanceConnector
	Grant              credit.GrantConnector
	MeteredEntitlement meteredentitlement.Connector
	Entitlement        entitlement.EntitlementConnector
}

type EntitlementOptions struct {
	DatabaseClient     *db.Client
	StreamingConnector streaming.Connector
	Logger             *slog.Logger
	MeterRepository    meter.Repository
	Publisher          publisher.TopicPublisher
}

func GetEntitlementRegistry(opts EntitlementOptions) *Entitlement {
	// Initialize database adapters
	featureDBAdapter := productcatalogpgadapter.NewPostgresFeatureDBAdapter(opts.DatabaseClient, opts.Logger)
	entitlementDBAdapter := entitlementpgadapter.NewPostgresEntitlementDBAdapter(opts.DatabaseClient)
	usageResetDBAdapter := entitlementpgadapter.NewPostgresUsageResetDBAdapter(opts.DatabaseClient)
	grantDBAdapter := creditpgadapter.NewPostgresGrantDBAdapter(opts.DatabaseClient)
	balanceSnashotDBAdapter := creditpgadapter.NewPostgresBalanceSnapshotDBAdapter(opts.DatabaseClient)

	// Initialize connectors
	featureConnector := productcatalog.NewFeatureConnector(featureDBAdapter, opts.MeterRepository)
	entitlementOwnerConnector := meteredentitlement.NewEntitlementGrantOwnerAdapter(
		featureDBAdapter,
		entitlementDBAdapter,
		usageResetDBAdapter,
		opts.MeterRepository,
		opts.Logger,
	)
	creditBalanceConnector := credit.NewBalanceConnector(
		grantDBAdapter,
		balanceSnashotDBAdapter,
		entitlementOwnerConnector,
		opts.StreamingConnector,
		opts.Logger,
	)
	grantConnector := credit.NewGrantConnector(
		entitlementOwnerConnector,
		grantDBAdapter,
		balanceSnashotDBAdapter,
		time.Minute,
		opts.Publisher,
	)
	meteredEntitlementConnector := meteredentitlement.NewMeteredEntitlementConnector(
		opts.StreamingConnector,
		entitlementOwnerConnector,
		creditBalanceConnector,
		grantConnector,
		entitlementDBAdapter,
		opts.Publisher,
	)
	entitlementConnector := entitlement.NewEntitlementConnector(
		entitlementDBAdapter,
		featureConnector,
		opts.MeterRepository,
		meteredEntitlementConnector,
		staticentitlement.NewStaticEntitlementConnector(),
		booleanentitlement.NewBooleanEntitlementConnector(),
		opts.Publisher,
	)

	return &Entitlement{
		Feature:            featureConnector,
		EntitlementOwner:   entitlementOwnerConnector,
		CreditBalance:      creditBalanceConnector,
		Grant:              grantConnector,
		MeteredEntitlement: meteredEntitlementConnector,
		Entitlement:        entitlementConnector,
	}
}
