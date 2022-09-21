package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/message/npool/third/mgr/v1/usedfor"
	"github.com/NpoolPlatform/third-manager/pkg/db/mixin"

	"github.com/google/uuid"
)

// EmailTemplate holds the schema definition for the EmailTemplate entity.
type EmailTemplate struct {
	ent.Schema
}

func (EmailTemplate) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the EmailTemplate.
func (EmailTemplate) Fields() []ent.Field {
	var maxLen = 8192
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("lang_id", uuid.UUID{}),
		field.String("default_to_username"),
		field.String("used_for").
			Optional().
			Default(usedfor.UsedFor_DefaultUsedFor.String()),
		field.String("sender").
			Optional().
			Default(""),
		field.JSON("reply_tos", []string{}).
			Optional().
			Default([]string{}),
		field.JSON("cc_tos", []string{}).
			Optional().
			Default([]string{}),
		field.String("subject").
			Optional().
			Default(""),
		field.String("body").
			Optional().
			Default("").
			MaxLen(maxLen),
	}
}
