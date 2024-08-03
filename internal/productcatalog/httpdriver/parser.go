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
	"github.com/openmeterio/openmeter/api"
	"github.com/openmeterio/openmeter/internal/productcatalog"
	"github.com/openmeterio/openmeter/pkg/convert"
)

func MaptFeatureToResponse(f productcatalog.Feature) api.Feature {
	return api.Feature{
		CreatedAt:           &f.CreatedAt,
		DeletedAt:           nil,
		UpdatedAt:           &f.UpdatedAt,
		Id:                  &f.ID,
		Key:                 f.Key,
		Metadata:            convert.MapToPointer(f.Metadata),
		Name:                f.Name,
		ArchivedAt:          f.ArchivedAt,
		MeterGroupByFilters: convert.MapToPointer(f.MeterGroupByFilters),
		MeterSlug:           f.MeterSlug,
	}
}
