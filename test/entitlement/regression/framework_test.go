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

package framework_test

import (
	"context"
	"log/slog"
	"testing"
	"time"

	"github.com/openmeterio/openmeter/internal/credit"
	grantrepo "github.com/openmeterio/openmeter/internal/credit/postgresadapter"
	"github.com/openmeterio/openmeter/internal/ent/db"
	"github.com/openmeterio/openmeter/internal/entitlement"
	booleanentitlement "github.com/openmeterio/openmeter/internal/entitlement/boolean"
	meteredentitlement "github.com/openmeterio/openmeter/internal/entitlement/metered"
	entitlementrepo "github.com/openmeterio/openmeter/internal/entitlement/postgresadapter"
	staticentitlement "github.com/openmeterio/openmeter/internal/entitlement/static"
	"github.com/openmeterio/openmeter/internal/event/publisher"
	"github.com/openmeterio/openmeter/internal/meter"
	"github.com/openmeterio/openmeter/internal/productcatalog"
	productcatalogrepo "github.com/openmeterio/openmeter/internal/productcatalog/postgresadapter"
	streamingtestutils "github.com/openmeterio/openmeter/internal/streaming/testutils"
	"github.com/openmeterio/openmeter/internal/testutils"
	"github.com/openmeterio/openmeter/pkg/models"
)

type Dependencies struct {
	DBClient *db.Client

	GrantRepo           credit.GrantRepo
	BalanceSnapshotRepo credit.BalanceSnapshotRepo
	GrantConnector      credit.GrantConnector

	EntitlementRepo entitlement.EntitlementRepo

	EntitlementConnector        entitlement.Connector
	StaticEntitlementConnector  staticentitlement.Connector
	BooleanEntitlementConnector booleanentitlement.Connector
	MeteredEntitlementConnector meteredentitlement.Connector

	Streaming *streamingtestutils.MockStreamingConnector

	FeatureRepo      productcatalog.FeatureRepo
	FeatureConnector productcatalog.FeatureConnector

	Log *slog.Logger
}

func (d *Dependencies) Close() {
	d.DBClient.Close()
}

func setupDependencies(t *testing.T) Dependencies {
	log := slog.Default()
	ctx := context.Background()
	driver := testutils.InitPostgresDB(t)

	// init db
	dbClient := db.NewClient(db.Driver(driver))
	if err := dbClient.Schema.Create(ctx); err != nil {
		t.Fatalf("failed to migrate database %s", err)
	}

	// Init product catalog
	featureRepo := productcatalogrepo.NewPostgresFeatureRepo(dbClient, log)

	meters := []models.Meter{
		{
			Namespace:   "namespace-1",
			ID:          "meter-1",
			Slug:        "meter-1",
			WindowSize:  models.WindowSizeMinute,
			Aggregation: models.MeterAggregationCount,
		},
	}

	meterRepo := meter.NewInMemoryRepository(meters)

	featureConnector := productcatalog.NewFeatureConnector(featureRepo, meterRepo) // TODO: meter repo is needed

	// Init grants/credit
	grantRepo := grantrepo.NewPostgresGrantRepo(dbClient)
	balanceSnapshotRepo := grantrepo.NewPostgresBalanceSnapshotRepo(dbClient)

	// Init entitlements
	streaming := streamingtestutils.NewMockStreamingConnector(t)

	entitlementRepo := entitlementrepo.NewPostgresEntitlementRepo(dbClient)
	usageResetRepo := entitlementrepo.NewPostgresUsageResetRepo(dbClient)

	owner := meteredentitlement.NewEntitlementGrantOwnerAdapter(
		featureRepo,
		entitlementRepo,
		usageResetRepo,
		meterRepo,
		log,
	)

	balance := credit.NewBalanceConnector(
		grantRepo,
		balanceSnapshotRepo,
		owner,
		streaming,
		log,
	)

	grant := credit.NewGrantConnector(
		owner,
		grantRepo,
		balanceSnapshotRepo,
		time.Minute,
		publisher.NewMockTopicPublisher(t),
	)

	meteredEntitlementConnector := meteredentitlement.NewMeteredEntitlementConnector(
		streaming,
		owner,
		balance,
		grant,
		entitlementRepo,
		publisher.NewMockTopicPublisher(t),
	)

	staticEntitlementConnector := staticentitlement.NewStaticEntitlementConnector()
	booleanEntitlementConnector := booleanentitlement.NewBooleanEntitlementConnector()

	entitlementConnector := entitlement.NewEntitlementConnector(
		entitlementRepo,
		featureConnector,
		meterRepo,
		meteredEntitlementConnector,
		staticEntitlementConnector,
		booleanEntitlementConnector,
		publisher.NewMockTopicPublisher(t),
	)

	return Dependencies{
		DBClient: dbClient,

		GrantRepo:      grantRepo,
		GrantConnector: grant,

		EntitlementRepo: entitlementRepo,

		EntitlementConnector:        entitlementConnector,
		StaticEntitlementConnector:  staticEntitlementConnector,
		BooleanEntitlementConnector: booleanEntitlementConnector,
		MeteredEntitlementConnector: meteredEntitlementConnector,

		BalanceSnapshotRepo: balanceSnapshotRepo,

		Streaming: streaming,

		FeatureRepo:      featureRepo,
		FeatureConnector: featureConnector,

		Log: log,
	}
}
