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

package credit

import (
	"log/slog"
	"time"

	"github.com/openmeterio/openmeter/internal/credit"
	"github.com/openmeterio/openmeter/internal/event/publisher"
	"github.com/openmeterio/openmeter/openmeter/streaming"
)

// TODO: adapters have to be exported here

func NewBalanceConnector(
	gc GrantRepo,
	bsc BalanceSnapshotRepo,
	oc OwnerConnector,
	sc streaming.Connector,
	log *slog.Logger,
) BalanceConnector {
	return credit.NewBalanceConnector(gc, bsc, oc, sc, log)
}

func NewGrantConnector(
	oc OwnerConnector,
	db GrantRepo,
	bsdb BalanceSnapshotRepo,
	granularity time.Duration,
	publisher publisher.TopicPublisher,
) GrantConnector {
	return credit.NewGrantConnector(oc, db, bsdb, granularity, publisher)
}
