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

package sink

import (
	"context"
	"fmt"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/huandu/go-sqlbuilder"

	"github.com/openmeterio/openmeter/internal/streaming/clickhouse_connector"
)

type Storage interface {
	BatchInsert(ctx context.Context, messages []SinkMessage) error
}

type ClickHouseStorageConfig struct {
	ClickHouse clickhouse.Conn
	Database   string
}

func NewClickhouseStorage(config ClickHouseStorageConfig) *ClickHouseStorage {
	return &ClickHouseStorage{
		config: config,
	}
}

type ClickHouseStorage struct {
	config ClickHouseStorageConfig
}

func (c *ClickHouseStorage) BatchInsert(ctx context.Context, messages []SinkMessage) error {
	query := InsertEventsQuery{
		Database: c.config.Database,
		Messages: messages,
	}
	sql, args, err := query.ToSQL()
	if err != nil {
		return err
	}

	err = c.config.ClickHouse.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("failed to batch insert events: %w", err)
	}

	return nil
}

type InsertEventsQuery struct {
	Database string
	Messages []SinkMessage
}

func (q InsertEventsQuery) ToSQL() (string, []interface{}, error) {
	tableName := clickhouse_connector.GetEventsTableName(q.Database)

	query := sqlbuilder.ClickHouse.NewInsertBuilder()
	query.InsertInto(tableName)
	query.Cols("namespace", "validation_error", "id", "type", "source", "subject", "time", "data")

	for _, message := range q.Messages {
		var eventErr string
		if message.Error != nil {
			eventErr = message.Error.Error()
		}

		query.Values(
			message.Namespace,
			eventErr,
			message.Serialized.Id,
			message.Serialized.Type,
			message.Serialized.Source,
			message.Serialized.Subject,
			message.Serialized.Time,
			message.Serialized.Data,
		)
	}

	sql, args := query.Build()
	return sql, args, nil
}
