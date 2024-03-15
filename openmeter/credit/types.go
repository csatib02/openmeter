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
	"github.com/openmeterio/openmeter/internal/credit"
)

type (
	BalanceConnector                        = credit.BalanceConnector
	BalanceHistoryParams                    = credit.BalanceHistoryParams
	BalanceSnapshotRepo                     = credit.BalanceSnapshotRepo
	CreateGrantInput                        = credit.CreateGrantInput
	DBCreateGrantInput                      = credit.GrantRepoCreateGrantInput
	Engine                                  = credit.Engine
	ExpirationPeriod                        = credit.ExpirationPeriod
	ExpirationPeriodDuration                = credit.ExpirationPeriodDuration
	Grant                                   = credit.Grant
	GrantBalanceMap                         = credit.GrantBalanceMap
	GrantBalanceNoSavedBalanceForOwnerError = credit.GrantBalanceNoSavedBalanceForOwnerError
	GrantBalanceSnapshot                    = credit.GrantBalanceSnapshot
	GrantBurnDownHistory                    = credit.GrantBurnDownHistory
	GrantBurnDownHistorySegment             = credit.GrantBurnDownHistorySegment
	GrantConnector                          = credit.GrantConnector
	GrantRepo                               = credit.GrantRepo
	GrantNotFoundError                      = credit.GrantNotFoundError
	GrantOrderBy                            = credit.GrantOrderBy
	GrantOwner                              = credit.GrantOwner
	GrantUsage                              = credit.GrantUsage
	GrantUsageTerminationReason             = credit.GrantUsageTerminationReason
	ListGrantsParams                        = credit.ListGrantsParams
	NamespacedGrantOwner                    = credit.NamespacedGrantOwner
	OwnerConnector                          = credit.OwnerConnector
	Pagination                              = credit.Pagination
	QueryUsageFn                            = credit.QueryUsageFn
	SegmentTerminationReason                = credit.SegmentTerminationReason
)
