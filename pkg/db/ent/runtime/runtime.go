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
	// appcontactDescAppID is the schema descriptor for app_id field.
	appcontactDescAppID := appcontactFields[1].Descriptor()
	// appcontact.DefaultAppID holds the default value on creation for the app_id field.
	appcontact.DefaultAppID = appcontactDescAppID.Default.(func() uuid.UUID)
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
	appemailtemplateMixin := schema.AppEmailTemplate{}.Mixin()
	appemailtemplate.Policy = privacy.NewPolicies(appemailtemplateMixin[0], schema.AppEmailTemplate{})
	appemailtemplate.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := appemailtemplate.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	appemailtemplateMixinFields0 := appemailtemplateMixin[0].Fields()
	_ = appemailtemplateMixinFields0
	appemailtemplateFields := schema.AppEmailTemplate{}.Fields()
	_ = appemailtemplateFields
	// appemailtemplateDescCreatedAt is the schema descriptor for created_at field.
	appemailtemplateDescCreatedAt := appemailtemplateMixinFields0[0].Descriptor()
	// appemailtemplate.DefaultCreatedAt holds the default value on creation for the created_at field.
	appemailtemplate.DefaultCreatedAt = appemailtemplateDescCreatedAt.Default.(func() uint32)
	// appemailtemplateDescUpdatedAt is the schema descriptor for updated_at field.
	appemailtemplateDescUpdatedAt := appemailtemplateMixinFields0[1].Descriptor()
	// appemailtemplate.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	appemailtemplate.DefaultUpdatedAt = appemailtemplateDescUpdatedAt.Default.(func() uint32)
	// appemailtemplate.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	appemailtemplate.UpdateDefaultUpdatedAt = appemailtemplateDescUpdatedAt.UpdateDefault.(func() uint32)
	// appemailtemplateDescDeletedAt is the schema descriptor for deleted_at field.
	appemailtemplateDescDeletedAt := appemailtemplateMixinFields0[2].Descriptor()
	// appemailtemplate.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	appemailtemplate.DefaultDeletedAt = appemailtemplateDescDeletedAt.Default.(func() uint32)
	// appemailtemplateDescBody is the schema descriptor for body field.
	appemailtemplateDescBody := appemailtemplateFields[9].Descriptor()
	// appemailtemplate.BodyValidator is a validator for the "body" field. It is called by the builders before save.
	appemailtemplate.BodyValidator = appemailtemplateDescBody.Validators[0].(func(string) error)
	// appemailtemplateDescID is the schema descriptor for id field.
	appemailtemplateDescID := appemailtemplateFields[0].Descriptor()
	// appemailtemplate.DefaultID holds the default value on creation for the id field.
	appemailtemplate.DefaultID = appemailtemplateDescID.Default.(func() uuid.UUID)
	appsmstemplateMixin := schema.AppSMSTemplate{}.Mixin()
	appsmstemplate.Policy = privacy.NewPolicies(appsmstemplateMixin[0], schema.AppSMSTemplate{})
	appsmstemplate.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := appsmstemplate.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	appsmstemplateMixinFields0 := appsmstemplateMixin[0].Fields()
	_ = appsmstemplateMixinFields0
	appsmstemplateFields := schema.AppSMSTemplate{}.Fields()
	_ = appsmstemplateFields
	// appsmstemplateDescCreatedAt is the schema descriptor for created_at field.
	appsmstemplateDescCreatedAt := appsmstemplateMixinFields0[0].Descriptor()
	// appsmstemplate.DefaultCreatedAt holds the default value on creation for the created_at field.
	appsmstemplate.DefaultCreatedAt = appsmstemplateDescCreatedAt.Default.(func() uint32)
	// appsmstemplateDescUpdatedAt is the schema descriptor for updated_at field.
	appsmstemplateDescUpdatedAt := appsmstemplateMixinFields0[1].Descriptor()
	// appsmstemplate.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	appsmstemplate.DefaultUpdatedAt = appsmstemplateDescUpdatedAt.Default.(func() uint32)
	// appsmstemplate.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	appsmstemplate.UpdateDefaultUpdatedAt = appsmstemplateDescUpdatedAt.UpdateDefault.(func() uint32)
	// appsmstemplateDescDeletedAt is the schema descriptor for deleted_at field.
	appsmstemplateDescDeletedAt := appsmstemplateMixinFields0[2].Descriptor()
	// appsmstemplate.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	appsmstemplate.DefaultDeletedAt = appsmstemplateDescDeletedAt.Default.(func() uint32)
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
	// appsmstemplateDescID is the schema descriptor for id field.
	appsmstemplateDescID := appsmstemplateFields[0].Descriptor()
	// appsmstemplate.DefaultID holds the default value on creation for the id field.
	appsmstemplate.DefaultID = appsmstemplateDescID.Default.(func() uuid.UUID)
}

const (
	Version = "v0.11.2"                                         // Version of ent codegen.
	Sum     = "h1:UM2/BUhF2FfsxPHRxLjQbhqJNaDdVlOwNIAMLs2jyto=" // Sum of ent codegen.
)
