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

package httpdriver

import (
	"github.com/openmeterio/openmeter/internal/productcatalog/httpdriver"
	"github.com/openmeterio/openmeter/openmeter/namespace/namespacedriver"
	"github.com/openmeterio/openmeter/openmeter/productcatalog"
	"github.com/openmeterio/openmeter/pkg/framework/transport/httptransport"
)

type (
	CreateFeatureHandler = httpdriver.CreateFeatureHandler
	DeleteFeatureHandler = httpdriver.DeleteFeatureHandler
	GetFeatureHandler    = httpdriver.GetFeatureHandler
	ListFeaturesHandler  = httpdriver.ListFeaturesHandler
	FeatureHandler       = httpdriver.FeatureHandler
)

func NewFeatureHandler(
	connector productcatalog.FeatureConnector,
	namespaceDecoder namespacedriver.NamespaceDecoder,
	options ...httptransport.HandlerOption,
) FeatureHandler {
	return httpdriver.NewFeatureHandler(connector, namespaceDecoder, options...)
}
