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

package quickstart

import (
	"context"
	"net/http"
	"os"
	"testing"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2/event"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	api "github.com/openmeterio/openmeter/api/client/go"
	"github.com/openmeterio/openmeter/pkg/models"
)

func initClient(t *testing.T) *api.ClientWithResponses {
	t.Helper()

	address := os.Getenv("OPENMETER_ADDRESS")
	if address == "" {
		t.Skip("OPENMETER_ADDRESS not set")
	}

	client, err := api.NewClientWithResponses(address)
	require.NoError(t, err)

	return client
}

func TestQuickstart(t *testing.T) {
	client := initClient(t)

	// TODO: read these from JSON files to make it easier to keep things in sync
	{
		ev := cloudevents.New()
		ev.SetID("00001")
		ev.SetSource("service-0")
		ev.SetType("request")
		ev.SetSubject("customer-1")
		ev.SetTime(time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC))
		_ = ev.SetData("application/json", map[string]string{
			"method": "GET",
			"route":  "/hello",
		})

		require.EventuallyWithT(t, func(t *assert.CollectT) {
			resp, err := client.IngestEventWithResponse(context.Background(), ev)
			require.NoError(t, err)
			require.Equal(t, http.StatusNoContent, resp.StatusCode())
		}, 30*time.Second, time.Second)
	}

	{
		ev := cloudevents.New()
		ev.SetID("00002")
		ev.SetSource("service-0")
		ev.SetType("request")
		ev.SetSubject("customer-1")
		ev.SetTime(time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC))
		_ = ev.SetData("application/json", map[string]string{
			"method": "GET",
			"route":  "/hello",
		})

		require.EventuallyWithT(t, func(t *assert.CollectT) {
			resp, err := client.IngestEventWithResponse(context.Background(), ev)
			require.NoError(t, err)
			require.Equal(t, http.StatusNoContent, resp.StatusCode())
		}, 30*time.Second, time.Second)
	}

	{
		ev := cloudevents.New()
		ev.SetID("00003")
		ev.SetSource("service-0")
		ev.SetType("request")
		ev.SetSubject("customer-1")
		ev.SetTime(time.Date(2023, time.January, 2, 0, 0, 0, 0, time.UTC))
		_ = ev.SetData("application/json", map[string]string{
			"method": "GET",
			"route":  "/hello",
		})

		require.EventuallyWithT(t, func(t *assert.CollectT) {
			resp, err := client.IngestEventWithResponse(context.Background(), ev)
			require.NoError(t, err)
			require.Equal(t, http.StatusNoContent, resp.StatusCode())
		}, 30*time.Second, time.Second)
	}

	assert.EventuallyWithT(t, func(t *assert.CollectT) {
		windowSize := models.WindowSizeHour

		resp, err := client.QueryMeterWithResponse(context.Background(), "api_requests_total", &api.QueryMeterParams{
			WindowSize: &windowSize,
		})
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode())

		require.Len(t, resp.JSON200.Data, 2)
		assert.Equal(t, float64(2), resp.JSON200.Data[0].Value)
		assert.Equal(t, float64(1), resp.JSON200.Data[1].Value)
	}, 30*time.Second, time.Second)
}
