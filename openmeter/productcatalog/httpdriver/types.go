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

import "github.com/openmeterio/openmeter/internal/productcatalog/httpdriver"

// requests
type (
	CreateFeatureHandlerRequest = httpdriver.CreateFeatureHandlerRequest
	DeleteFeatureHandlerRequest = httpdriver.DeleteFeatureHandlerRequest
	GetFeatureHandlerRequest    = httpdriver.GetFeatureHandlerRequest
	ListFeaturesHandlerRequest  = httpdriver.ListFeaturesHandlerRequest
)

// responses
type (
	CreateFeatureHandlerResponse = httpdriver.CreateFeatureHandlerResponse
	DeleteFeatureHandlerResponse = httpdriver.DeleteFeatureHandlerResponse
	GetFeatureHandlerResponse    = httpdriver.GetFeatureHandlerResponse
	ListFeaturesHandlerResponse  = httpdriver.ListFeaturesHandlerResponse
)

// params
type (
	DeleteFeatureHandlerParams = httpdriver.DeleteFeatureHandlerParams
	GetFeatureHandlerParams    = httpdriver.GetFeatureHandlerParams
	ListFeaturesHandlerParams  = httpdriver.ListFeaturesHandlerParams
)
