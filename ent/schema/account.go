package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Comment("ID").Default(uuid.New).Immutable().Unique(),
		field.String("name").Comment("名前").NotEmpty(),
		field.Time("created_at").Comment("作成日").Default(time.Now).Immutable(),
		field.Time("updated_at").Comment("更新日").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Comment("削除日").Optional().Nillable(),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return nil
}
