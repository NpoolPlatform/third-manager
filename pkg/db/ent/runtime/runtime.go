// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"context"

	"github.com/NpoolPlatform/third-manager/pkg/db/ent/appcontact"
	"github.com/NpoolPlatform/third-manager/pkg/db/ent/appemailtemplate"
	"github.com/NpoolPlatform/third-manager/pkg/db/ent/appsmstemplate"
	"github.com/NpoolPlatform/third-manager/pkg/db/ent/schema"
	"github.com/google/uuid"

	"entgo.io/ent"
	"entgo.io/ent/privacy"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	appcontactMixin := schema.AppContact{}.Mixin()
	appcontact.Policy = privacy.NewPolicies(appcontactMixin[0], schema.AppContact{})
	appcontact.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := appcontact.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	appcontactMixinFields0 := appcontactMixin[0].Fields()
	_ = appcontactMixinFields0
	appcontactFields := schema.AppContact{}.Fields()
	_ = appcontactFields
	// appcontactDescCreatedAt is the schema descriptor for created_at field.
	appcontactDescCreatedAt := appcontactMixinFields0[0].Descriptor()
	// appcontact.DefaultCreatedAt holds the default value on creation for the created_at field.
	appcontact.DefaultCreatedAt = appcontactDescCreatedAt.Default.(func() uint32)
	// appcontactDescUpdatedAt is the schema descriptor for updated_at field.
	appcontactDescUpdatedAt := appcontactMixinFields0[1].Descriptor()
	// appcontact.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	appcontact.DefaultUpdatedAt = appcontactDescUpdatedAt.Default.(func() uint32)
	// appcontact.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	appcontact.UpdateDefaultUpdatedAt = appcontactDescUpdatedAt.UpdateDefault.(func() uint32)
	// appcontactDescDeletedAt is the schema descriptor for deleted_at field.
	appcontactDescDeletedAt := appcontactMixinFields0[2].Descriptor()
	// appcontact.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	appcontact.DefaultDeletedAt = appcontactDescDeletedAt.Default.(func() uint32)
	// appcontactDescUsedFor is the schema descriptor for used_for field.
	appcontactDescUsedFor := appcontactFields[2].Descriptor()
	// appcontact.DefaultUsedFor holds the default value on creation for the used_for field.
	appcontact.DefaultUsedFor = appcontactDescUsedFor.Default.(string)
	// appcontact.UsedForValidator is a validator for the "used_for" field. It is called by the builders before save.
	appcontact.UsedForValidator = appcontactDescUsedFor.Validators[0].(func(string) error)
	// appcontactDescSender is the schema descriptor for sender field.
	appcontactDescSender := appcontactFields[3].Descriptor()
	// appcontact.DefaultSender holds the default value on creation for the sender field.
	appcontact.DefaultSender = appcontactDescSender.Default.(string)
	// appcontactDescAccount is the schema descriptor for account field.
	appcontactDescAccount := appcontactFields[4].Descriptor()
	// appcontact.DefaultAccount holds the default value on creation for the account field.
	appcontact.DefaultAccount = appcontactDescAccount.Default.(string)
	// appcontactDescAccountType is the schema descriptor for account_type field.
	appcontactDescAccountType := appcontactFields[5].Descriptor()
	// appcontact.DefaultAccountType holds the default value on creation for the account_type field.
	appcontact.DefaultAccountType = appcontactDescAccountType.Default.(string)
	// appcontactDescID is the schema descriptor for id field.
	appcontactDescID := appcontactFields[0].Descriptor()
	// appcontact.DefaultID holds the default value on creation for the id field.
	appcontact.DefaultID = appcontactDescID.Default.(func() uuid.UUID)
	appemailtemplateFields := schema.AppEmailTemplate{}.Fields()
	_ = appemailtemplateFields
	// appemailtemplateDescBody is the schema descriptor for body field.
	appemailtemplateDescBody := appemailtemplateFields[9].Descriptor()
	// appemailtemplate.BodyValidator is a validator for the "body" field. It is called by the builders before save.
	appemailtemplate.BodyValidator = appemailtemplateDescBody.Validators[0].(func(string) error)
	// appemailtemplateDescCreateAt is the schema descriptor for create_at field.
	appemailtemplateDescCreateAt := appemailtemplateFields[10].Descriptor()
	// appemailtemplate.DefaultCreateAt holds the default value on creation for the create_at field.
	appemailtemplate.DefaultCreateAt = appemailtemplateDescCreateAt.Default.(func() uint32)
	// appemailtemplateDescUpdateAt is the schema descriptor for update_at field.
	appemailtemplateDescUpdateAt := appemailtemplateFields[11].Descriptor()
	// appemailtemplate.DefaultUpdateAt holds the default value on creation for the update_at field.
	appemailtemplate.DefaultUpdateAt = appemailtemplateDescUpdateAt.Default.(func() uint32)
	// appemailtemplate.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	appemailtemplate.UpdateDefaultUpdateAt = appemailtemplateDescUpdateAt.UpdateDefault.(func() uint32)
	// appemailtemplateDescID is the schema descriptor for id field.
	appemailtemplateDescID := appemailtemplateFields[0].Descriptor()
	// appemailtemplate.DefaultID holds the default value on creation for the id field.
	appemailtemplate.DefaultID = appemailtemplateDescID.Default.(func() uuid.UUID)
	appsmstemplateFields := schema.AppSMSTemplate{}.Fields()
	_ = appsmstemplateFields
	// appsmstemplateDescUsedFor is the schema descriptor for used_for field.
	appsmstemplateDescUsedFor := appsmstemplateFields[3].Descriptor()
	// appsmstemplate.DefaultUsedFor holds the default value on creation for the used_for field.
	appsmstemplate.DefaultUsedFor = appsmstemplateDescUsedFor.Default.(string)
	// appsmstemplateDescSubject is the schema descriptor for subject field.
	appsmstemplateDescSubject := appsmstemplateFields[4].Descriptor()
	// appsmstemplate.DefaultSubject holds the default value on creation for the subject field.
	appsmstemplate.DefaultSubject = appsmstemplateDescSubject.Default.(string)
	// appsmstemplateDescMessage is the schema descriptor for message field.
	appsmstemplateDescMessage := appsmstemplateFields[5].Descriptor()
	// appsmstemplate.DefaultMessage holds the default value on creation for the message field.
	appsmstemplate.DefaultMessage = appsmstemplateDescMessage.Default.(string)
	// appsmstemplateDescCreateAt is the schema descriptor for create_at field.
	appsmstemplateDescCreateAt := appsmstemplateFields[6].Descriptor()
	// appsmstemplate.DefaultCreateAt holds the default value on creation for the create_at field.
	appsmstemplate.DefaultCreateAt = appsmstemplateDescCreateAt.Default.(func() uint32)
	// appsmstemplateDescUpdateAt is the schema descriptor for update_at field.
	appsmstemplateDescUpdateAt := appsmstemplateFields[7].Descriptor()
	// appsmstemplate.DefaultUpdateAt holds the default value on creation for the update_at field.
	appsmstemplate.DefaultUpdateAt = appsmstemplateDescUpdateAt.Default.(func() uint32)
	// appsmstemplate.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	appsmstemplate.UpdateDefaultUpdateAt = appsmstemplateDescUpdateAt.UpdateDefault.(func() uint32)
	// appsmstemplateDescID is the schema descriptor for id field.
	appsmstemplateDescID := appsmstemplateFields[0].Descriptor()
	// appsmstemplate.DefaultID holds the default value on creation for the id field.
	appsmstemplate.DefaultID = appsmstemplateDescID.Default.(func() uuid.UUID)
}

const (
	Version = "v0.11.2"                                         // Version of ent codegen.
	Sum     = "h1:UM2/BUhF2FfsxPHRxLjQbhqJNaDdVlOwNIAMLs2jyto=" // Sum of ent codegen.
)