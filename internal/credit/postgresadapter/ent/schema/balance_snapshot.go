package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/openmeterio/openmeter/internal/credit"
	"github.com/openmeterio/openmeter/pkg/framework/entutils"
)

type BalanceSnapshot struct {
	ent.Schema
}

func (BalanceSnapshot) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entutils.NamespaceMixin{},
		entutils.TimeMixin{}, // maybe we don't need this
	}
}

func (BalanceSnapshot) Fields() []ent.Field {
	return []ent.Field{
		field.String("owner_id").GoType(credit.GrantOwner("")).Immutable().SchemaType(map[string]string{
			dialect.Postgres: "char(26)",
		}),
		field.JSON("grant_balances", credit.GrantBalanceMap{}).Immutable().SchemaType(map[string]string{
			dialect.Postgres: "jsonb",
		}),
		field.Float("balance").Immutable().SchemaType(map[string]string{
			dialect.Postgres: "numeric",
		}),
		field.Float("overage").Immutable().SchemaType(map[string]string{
			dialect.Postgres: "numeric",
		}),
		field.Time("at").Immutable(),
	}
}

func (BalanceSnapshot) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("namespace", "at"),
		index.Fields("namespace", "balance"),
		index.Fields("namespace", "balance", "at"),
	}
}

func (BalanceSnapshot) Edges() []ent.Edge {
	return []ent.Edge{}
}
