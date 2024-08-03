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

package flushhandler

import "github.com/openmeterio/openmeter/internal/sink/flushhandler"

type (
	FlushEventHandler = flushhandler.FlushEventHandler
	FlushCallback     = flushhandler.FlushCallback
)

// FlushHandlers
type (
	DrainCompleteFunc  = flushhandler.DrainCompleteFunc
	FlushEventHandlers = flushhandler.FlushEventHandlers
)

func NewFlushEventHandlers() *FlushEventHandlers {
	return flushhandler.NewFlushEventHandlers()
}

// FlushHandler
type (
	FlushEventHandlerOptions = flushhandler.FlushEventHandlerOptions
)

func NewFlushEventHandler(opts FlushEventHandlerOptions) (FlushEventHandler, error) {
	return flushhandler.NewFlushEventHandler(opts)
}
