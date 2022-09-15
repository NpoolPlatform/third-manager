package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
)

// AppEmailTemplate holds the schema definition for the AppEmailTemplate entity.
type AppEmailTemplate struct {
	ent.Schema
}

// Fields of the AppEmailTemplate.
func (AppEmailTemplate) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("lang_id", uuid.UUID{}),
		field.String("default_to_username"),
		field.String("used_for"),
		field.String("sender"),
		field.JSON("reply_tos", []string{}),
		field.JSON("cc_tos", []string{}),
		field.String("subject"),
		field.String("body").MaxLen(8192),
		field.Uint32("create_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("update_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}).
			UpdateDefault(func() uint32 {
				return uint32(time.Now().Unix())
			}),
	}
}

// Edges of the AppEmailTemplate.
func (AppEmailTemplate) Edges() []ent.Edge {
	return nil
}

func (AppEmailTemplate) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "lang_id", "used_for").
			Unique(),
	}
}
