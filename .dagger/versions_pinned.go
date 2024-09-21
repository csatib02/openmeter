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

package main

const (
	// Alpine is required for our current build (due to Kafka and CGO), but it doesn't seem to work well with golangci-lint
	goVersion = "1.23.0"

	kafkaVersion      = "3.6"
	clickhouseVersion = "24.5.5.78"
	redisVersion      = "7.0.12"
	postgresVersion   = "14.9"
	svixVersion       = "v1.29"

	// TODO: add update mechanism for versions below

	// Alpine is required for our current build (due to Kafka and CGO), but it doesn't seem to work well with golangci-lint
	goBuildVersion = goVersion + "-alpine3.20@sha256:d0b31558e6b3e4cc59f6011d79905835108c919143ebecc58f35965bf79948f4"
	xxBaseImage    = "tonistiigi/xx:1.5.0@sha256:0c6a569797744e45955f39d4f7538ac344bfb7ebf0a54006a0a4297b153ccf0f"

	alpineBaseImage = "alpine:3.20.2@sha256:0a4eaa0eecf5f8c050e5bba433f58c052be7587ee8af3e8b3910ef9ab5fbe9f5"

	atlasImage = "arigaio/atlas:0.26.1@sha256:f77152f5458255410d2e59ad80a0fc661524067a90d297a0b62e6bda23437893"
)
