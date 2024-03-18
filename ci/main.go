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

import (
	"context"
)

const (
	// Alpine is required for our current build (due to Kafka and CGO), but it doesn't seem to work well with golangci-lint
	goVersion      = "1.22.0"
	goBuildVersion = goVersion + "-alpine3.18@sha256:2745a45f77ae2e7be569934fa9a111f067d04c767f54577e251d9b101250e46b"

	golangciLintVersion = "v1.56.1"
	spectralVersion     = "6.11"
	kafkaVersion        = "3.6"
	clickhouseVersion   = "23.3.9.55"
	redisVersion        = "7.0.12"

	helmDocsVersion = "v1.11.3"
	helmVersion     = "3.13.2"

	ghVersion = "2.42.1"

	alpineBaseImage = "alpine:3.19.1@sha256:c5b1261d6d3e43071626931fc004f70149baeba2c8ec672bd4f27761f8e1ad6b"
	xxBaseImage     = "tonistiigi/xx:1.3.0@sha256:904fe94f236d36d65aeb5a2462f88f2c537b8360475f6342e7599194f291fb7e"
)

type Ci struct {
	// Project source directory
	// This will become useful once pulling from remote becomes available
	//
	// +private
	Source *Directory
}

func New(
	// Checkout the repository (at the designated ref) and use it as the source directory instead of the local one.
	// +optional
	checkout string,
) *Ci {
	var source *Directory

	if checkout != "" {
		source = dag.Git("https://github.com/openmeterio/openmeter.git", GitOpts{
			KeepGitDir: true,
		}).Branch(checkout).Tree()
	} else {
		source = projectDir()
	}
	return &Ci{
		Source: source,
	}
}

func (m *Ci) Ci(ctx context.Context) error {
	p := newPipeline(ctx)

	p.addJobs(
		wrapSyncable(m.Test()),
		m.Lint().All,

		// TODO: run trivy scan on container(s?)
		// TODO: version should be the commit hash (if any?)?
		wrapSyncables(m.Build().containerImages("ci")),

		// TODO: run trivy scan on helm chart
		wrapSyncable(m.Build().HelmChart("openmeter", "0.0.0")),
		wrapSyncable(m.Build().HelmChart("benthos-collector", "0.0.0")),

		wrapSyncables(m.releaseAssets("ci")),
	)

	return p.wait()
}

func (m *Ci) Test() *Container {
	return dag.Go().
		WithSource(m.Source).
		Exec([]string{"go", "test", "-v", "./..."})
}
