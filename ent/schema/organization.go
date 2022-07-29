package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Organization holds the schema definition for the Organization entity.
type Organization struct {
	ent.Schema
}

// Fields of the Organization.
func (Organization) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Comment("ID").Default(uuid.New).Immutable().Unique(),
		field.String("name").Comment("組織名").NotEmpty(),
		field.Time("created_at").Comment("作成日").Default(time.Now).Immutable(),
		field.Time("updated_at").Comment("更新日").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Comment("削除日").Optional().Nillable(),
	}
}

// Edges of the Organization.
func (Organization) Edges() []ent.Edge {
	return nil
}
