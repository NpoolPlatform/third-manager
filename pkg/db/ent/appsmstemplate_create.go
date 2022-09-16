// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/third-manager/pkg/db/ent/appsmstemplate"
	"github.com/google/uuid"
)

// AppSMSTemplateCreate is the builder for creating a AppSMSTemplate entity.
type AppSMSTemplateCreate struct {
	config
	mutation *AppSMSTemplateMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (astc *AppSMSTemplateCreate) SetCreatedAt(u uint32) *AppSMSTemplateCreate {
	astc.mutation.SetCreatedAt(u)
	return astc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (astc *AppSMSTemplateCreate) SetNillableCreatedAt(u *uint32) *AppSMSTemplateCreate {
	if u != nil {
		astc.SetCreatedAt(*u)
	}
	return astc
}

// SetUpdatedAt sets the "updated_at" field.
func (astc *AppSMSTemplateCreate) SetUpdatedAt(u uint32) *AppSMSTemplateCreate {
	astc.mutation.SetUpdatedAt(u)
	return astc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (astc *AppSMSTemplateCreate) SetNillableUpdatedAt(u *uint32) *AppSMSTemplateCreate {
	if u != nil {
		astc.SetUpdatedAt(*u)
	}
	return astc
}

// SetDeletedAt sets the "deleted_at" field.
func (astc *AppSMSTemplateCreate) SetDeletedAt(u uint32) *AppSMSTemplateCreate {
	astc.mutation.SetDeletedAt(u)
	return astc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (astc *AppSMSTemplateCreate) SetNillableDeletedAt(u *uint32) *AppSMSTemplateCreate {
	if u != nil {
		astc.SetDeletedAt(*u)
	}
	return astc
}

// SetAppID sets the "app_id" field.
func (astc *AppSMSTemplateCreate) SetAppID(u uuid.UUID) *AppSMSTemplateCreate {
	astc.mutation.SetAppID(u)
	return astc
}

// SetLangID sets the "lang_id" field.
func (astc *AppSMSTemplateCreate) SetLangID(u uuid.UUID) *AppSMSTemplateCreate {
	astc.mutation.SetLangID(u)
	return astc
}

// SetUsedFor sets the "used_for" field.
func (astc *AppSMSTemplateCreate) SetUsedFor(s string) *AppSMSTemplateCreate {
	astc.mutation.SetUsedFor(s)
	return astc
}

// SetNillableUsedFor sets the "used_for" field if the given value is not nil.
func (astc *AppSMSTemplateCreate) SetNillableUsedFor(s *string) *AppSMSTemplateCreate {
	if s != nil {
		astc.SetUsedFor(*s)
	}
	return astc
}

// SetSubject sets the "subject" field.
func (astc *AppSMSTemplateCreate) SetSubject(s string) *AppSMSTemplateCreate {
	astc.mutation.SetSubject(s)
	return astc
}

// SetNillableSubject sets the "subject" field if the given value is not nil.
func (astc *AppSMSTemplateCreate) SetNillableSubject(s *string) *AppSMSTemplateCreate {
	if s != nil {
		astc.SetSubject(*s)
	}
	return astc
}

// SetMessage sets the "message" field.
func (astc *AppSMSTemplateCreate) SetMessage(s string) *AppSMSTemplateCreate {
	astc.mutation.SetMessage(s)
	return astc
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (astc *AppSMSTemplateCreate) SetNillableMessage(s *string) *AppSMSTemplateCreate {
	if s != nil {
		astc.SetMessage(*s)
	}
	return astc
}

// SetID sets the "id" field.
func (astc *AppSMSTemplateCreate) SetID(u uuid.UUID) *AppSMSTemplateCreate {
	astc.mutation.SetID(u)
	return astc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (astc *AppSMSTemplateCreate) SetNillableID(u *uuid.UUID) *AppSMSTemplateCreate {
	if u != nil {
		astc.SetID(*u)
	}
	return astc
}

// Mutation returns the AppSMSTemplateMutation object of the builder.
func (astc *AppSMSTemplateCreate) Mutation() *AppSMSTemplateMutation {
	return astc.mutation
}

// Save creates the AppSMSTemplate in the database.
func (astc *AppSMSTemplateCreate) Save(ctx context.Context) (*AppSMSTemplate, error) {
	var (
		err  error
		node *AppSMSTemplate
	)
	if err := astc.defaults(); err != nil {
		return nil, err
	}
	if len(astc.hooks) == 0 {
		if err = astc.check(); err != nil {
			return nil, err
		}
		node, err = astc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppSMSTemplateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = astc.check(); err != nil {
				return nil, err
			}
			astc.mutation = mutation
			if node, err = astc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(astc.hooks) - 1; i >= 0; i-- {
			if astc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = astc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, astc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AppSMSTemplate)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AppSMSTemplateMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (astc *AppSMSTemplateCreate) SaveX(ctx context.Context) *AppSMSTemplate {
	v, err := astc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (astc *AppSMSTemplateCreate) Exec(ctx context.Context) error {
	_, err := astc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (astc *AppSMSTemplateCreate) ExecX(ctx context.Context) {
	if err := astc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (astc *AppSMSTemplateCreate) defaults() error {
	if _, ok := astc.mutation.CreatedAt(); !ok {
		if appsmstemplate.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized appsmstemplate.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := appsmstemplate.DefaultCreatedAt()
		astc.mutation.SetCreatedAt(v)
	}
	if _, ok := astc.mutation.UpdatedAt(); !ok {
		if appsmstemplate.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized appsmstemplate.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := appsmstemplate.DefaultUpdatedAt()
		astc.mutation.SetUpdatedAt(v)
	}
	if _, ok := astc.mutation.DeletedAt(); !ok {
		if appsmstemplate.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized appsmstemplate.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := appsmstemplate.DefaultDeletedAt()
		astc.mutation.SetDeletedAt(v)
	}
	if _, ok := astc.mutation.UsedFor(); !ok {
		v := appsmstemplate.DefaultUsedFor
		astc.mutation.SetUsedFor(v)
	}
	if _, ok := astc.mutation.Subject(); !ok {
		v := appsmstemplate.DefaultSubject
		astc.mutation.SetSubject(v)
	}
	if _, ok := astc.mutation.Message(); !ok {
		v := appsmstemplate.DefaultMessage
		astc.mutation.SetMessage(v)
	}
	if _, ok := astc.mutation.ID(); !ok {
		if appsmstemplate.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized appsmstemplate.DefaultID (forgotten import ent/runtime?)")
		}
		v := appsmstemplate.DefaultID()
		astc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (astc *AppSMSTemplateCreate) check() error {
	if _, ok := astc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AppSMSTemplate.created_at"`)}
	}
	if _, ok := astc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "AppSMSTemplate.updated_at"`)}
	}
	if _, ok := astc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "AppSMSTemplate.deleted_at"`)}
	}
	if _, ok := astc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "AppSMSTemplate.app_id"`)}
	}
	if _, ok := astc.mutation.LangID(); !ok {
		return &ValidationError{Name: "lang_id", err: errors.New(`ent: missing required field "AppSMSTemplate.lang_id"`)}
	}
	if _, ok := astc.mutation.UsedFor(); !ok {
		return &ValidationError{Name: "used_for", err: errors.New(`ent: missing required field "AppSMSTemplate.used_for"`)}
	}
	if _, ok := astc.mutation.Subject(); !ok {
		return &ValidationError{Name: "subject", err: errors.New(`ent: missing required field "AppSMSTemplate.subject"`)}
	}
	if _, ok := astc.mutation.Message(); !ok {
		return &ValidationError{Name: "message", err: errors.New(`ent: missing required field "AppSMSTemplate.message"`)}
	}
	return nil
}

func (astc *AppSMSTemplateCreate) sqlSave(ctx context.Context) (*AppSMSTemplate, error) {
	_node, _spec := astc.createSpec()
	if err := sqlgraph.CreateNode(ctx, astc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (astc *AppSMSTemplateCreate) createSpec() (*AppSMSTemplate, *sqlgraph.CreateSpec) {
	var (
		_node = &AppSMSTemplate{config: astc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: appsmstemplate.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appsmstemplate.FieldID,
			},
		}
	)
	_spec.OnConflict = astc.conflict
	if id, ok := astc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := astc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appsmstemplate.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := astc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appsmstemplate.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := astc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appsmstemplate.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := astc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appsmstemplate.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := astc.mutation.LangID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appsmstemplate.FieldLangID,
		})
		_node.LangID = value
	}
	if value, ok := astc.mutation.UsedFor(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appsmstemplate.FieldUsedFor,
		})
		_node.UsedFor = value
	}
	if value, ok := astc.mutation.Subject(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appsmstemplate.FieldSubject,
		})
		_node.Subject = value
	}
	if value, ok := astc.mutation.Message(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appsmstemplate.FieldMessage,
		})
		_node.Message = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppSMSTemplate.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppSMSTemplateUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (astc *AppSMSTemplateCreate) OnConflict(opts ...sql.ConflictOption) *AppSMSTemplateUpsertOne {
	astc.conflict = opts
	return &AppSMSTemplateUpsertOne{
		create: astc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppSMSTemplate.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (astc *AppSMSTemplateCreate) OnConflictColumns(columns ...string) *AppSMSTemplateUpsertOne {
	astc.conflict = append(astc.conflict, sql.ConflictColumns(columns...))
	return &AppSMSTemplateUpsertOne{
		create: astc,
	}
}

type (
	// AppSMSTemplateUpsertOne is the builder for "upsert"-ing
	//  one AppSMSTemplate node.
	AppSMSTemplateUpsertOne struct {
		create *AppSMSTemplateCreate
	}

	// AppSMSTemplateUpsert is the "OnConflict" setter.
	AppSMSTemplateUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *AppSMSTemplateUpsert) SetCreatedAt(v uint32) *AppSMSTemplateUpsert {
	u.Set(appsmstemplate.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppSMSTemplateUpsert) UpdateCreatedAt() *AppSMSTemplateUpsert {
	u.SetExcluded(appsmstemplate.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppSMSTemplateUpsert) AddCreatedAt(v uint32) *AppSMSTemplateUpsert {
	u.Add(appsmstemplate.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppSMSTemplateUpsert) SetUpdatedAt(v uint32) *AppSMSTemplateUpsert {
	u.Set(appsmstemplate.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppSMSTemplateUpsert) UpdateUpdatedAt() *AppSMSTemplateUpsert {
	u.SetExcluded(appsmstemplate.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppSMSTemplateUpsert) AddUpdatedAt(v uint32) *AppSMSTemplateUpsert {
	u.Add(appsmstemplate.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppSMSTemplateUpsert) SetDeletedAt(v uint32) *AppSMSTemplateUpsert {
	u.Set(appsmstemplate.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppSMSTemplateUpsert) UpdateDeletedAt() *AppSMSTemplateUpsert {
	u.SetExcluded(appsmstemplate.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppSMSTemplateUpsert) AddDeletedAt(v uint32) *AppSMSTemplateUpsert {
	u.Add(appsmstemplate.FieldDeletedAt, v)
	return u
}

// SetAppID sets the "app_id" field.
func (u *AppSMSTemplateUpsert) SetAppID(v uuid.UUID) *AppSMSTemplateUpsert {
	u.Set(appsmstemplate.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppSMSTemplateUpsert) UpdateAppID() *AppSMSTemplateUpsert {
	u.SetExcluded(appsmstemplate.FieldAppID)
	return u
}

// SetLangID sets the "lang_id" field.
func (u *AppSMSTemplateUpsert) SetLangID(v uuid.UUID) *AppSMSTemplateUpsert {
	u.Set(appsmstemplate.FieldLangID, v)
	return u
}

// UpdateLangID sets the "lang_id" field to the value that was provided on create.
func (u *AppSMSTemplateUpsert) UpdateLangID() *AppSMSTemplateUpsert {
	u.SetExcluded(appsmstemplate.FieldLangID)
	return u
}

// SetUsedFor sets the "used_for" field.
func (u *AppSMSTemplateUpsert) SetUsedFor(v string) *AppSMSTemplateUpsert {
	u.Set(appsmstemplate.FieldUsedFor, v)
	return u
}

// UpdateUsedFor sets the "used_for" field to the value that was provided on create.
func (u *AppSMSTemplateUpsert) UpdateUsedFor() *AppSMSTemplateUpsert {
	u.SetExcluded(appsmstemplate.FieldUsedFor)
	return u
}

// SetSubject sets the "subject" field.
func (u *AppSMSTemplateUpsert) SetSubject(v string) *AppSMSTemplateUpsert {
	u.Set(appsmstemplate.FieldSubject, v)
	return u
}

// UpdateSubject sets the "subject" field to the value that was provided on create.
func (u *AppSMSTemplateUpsert) UpdateSubject() *AppSMSTemplateUpsert {
	u.SetExcluded(appsmstemplate.FieldSubject)
	return u
}

// SetMessage sets the "message" field.
func (u *AppSMSTemplateUpsert) SetMessage(v string) *AppSMSTemplateUpsert {
	u.Set(appsmstemplate.FieldMessage, v)
	return u
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *AppSMSTemplateUpsert) UpdateMessage() *AppSMSTemplateUpsert {
	u.SetExcluded(appsmstemplate.FieldMessage)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.AppSMSTemplate.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appsmstemplate.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppSMSTemplateUpsertOne) UpdateNewValues() *AppSMSTemplateUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(appsmstemplate.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.AppSMSTemplate.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *AppSMSTemplateUpsertOne) Ignore() *AppSMSTemplateUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppSMSTemplateUpsertOne) DoNothing() *AppSMSTemplateUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppSMSTemplateCreate.OnConflict
// documentation for more info.
func (u *AppSMSTemplateUpsertOne) Update(set func(*AppSMSTemplateUpsert)) *AppSMSTemplateUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppSMSTemplateUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AppSMSTemplateUpsertOne) SetCreatedAt(v uint32) *AppSMSTemplateUpsertOne {
	return u.Update(func(s *AppSMSTemplateUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppSMSTemplateUpsertOne) AddCreatedAt(v uint32) *AppSMSTemplateUpsertOne {
	return u.Update(func(s *AppSMSTemplateUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppSMSTemplateUpsertOne) UpdateCreatedAt() *AppSMSTemplateUpsertOne {
	return u.Update(func(s *AppSMSTemplateUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppSMSTemplateUpsertOne) SetUpdatedAt(v uint32) *AppSMSTemplateUpsertOne {
	return u.Update(func(s *AppSMSTemplateUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppSMSTemplateUpsertOne) AddUpdatedAt(v uint32) *AppSMSTemplateUpsertOne {
	return u.Update(func(s *AppSMSTemplateUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppSMSTemplateUpsertOne) UpdateUpdatedAt() *AppSMSTemplateUpsertOne {
	return u.Update(func(s *AppSMSTemplateUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppSMSTemplateUpsertOne) SetDeletedAt(v uint32) *AppSMSTemplateUpsertOne {
	return u.Update(func(s *AppSMSTemplateUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppSMSTemplateUpsertOne) AddDeletedAt(v uint32) *AppSMSTemplateUpsertOne {
	return u.Update(func(s *AppSMSTemplateUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppSMSTemplateUpsertOne) UpdateDeletedAt() *AppSMSTemplateUpsertOne {
	return u.Update(func(s *AppSMSTemplateUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetAppID sets the "app_id" field.
func (u *AppSMSTemplateUpsertOne) SetAppID(v uuid.UUID) *AppSMSTemplateUpsertOne {
	return u.Update(func(s *AppSMSTemplateUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppSMSTemplateUpsertOne) UpdateAppID() *AppSMSTemplateUpsertOne {
	return u.Update(func(s *AppSMSTemplateUpsert) {
		s.UpdateAppID()
	})
}

// SetLangID sets the "lang_id" field.
func (u *AppSMSTemplateUpsertOne) SetLangID(v uuid.UUID) *AppSMSTemplateUpsertOne {
	return u.Update(func(s *AppSMSTemplateUpsert) {
		s.SetLangID(v)
	})
}

// UpdateLangID sets the "lang_id" field to the value that was provided on create.
func (u *AppSMSTemplateUpsertOne) UpdateLangID() *AppSMSTemplateUpsertOne {
	return u.Update(func(s *AppSMSTemplateUpsert) {
		s.UpdateLangID()
	})
}

// SetUsedFor sets the "used_for" field.
func (u *AppSMSTemplateUpsertOne) SetUsedFor(v string) *AppSMSTemplateUpsertOne {
	return u.Update(func(s *AppSMSTemplateUpsert) {
		s.SetUsedFor(v)
	})
}

// UpdateUsedFor sets the "used_for" field to the value that was provided on create.
func (u *AppSMSTemplateUpsertOne) UpdateUsedFor() *AppSMSTemplateUpsertOne {
	return u.Update(func(s *AppSMSTemplateUpsert) {
		s.UpdateUsedFor()
	})
}

// SetSubject sets the "subject" field.
func (u *AppSMSTemplateUpsertOne) SetSubject(v string) *AppSMSTemplateUpsertOne {
	return u.Update(func(s *AppSMSTemplateUpsert) {
		s.SetSubject(v)
	})
}

// UpdateSubject sets the "subject" field to the value that was provided on create.
func (u *AppSMSTemplateUpsertOne) UpdateSubject() *AppSMSTemplateUpsertOne {
	return u.Update(func(s *AppSMSTemplateUpsert) {
		s.UpdateSubject()
	})
}

// SetMessage sets the "message" field.
func (u *AppSMSTemplateUpsertOne) SetMessage(v string) *AppSMSTemplateUpsertOne {
	return u.Update(func(s *AppSMSTemplateUpsert) {
		s.SetMessage(v)
	})
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *AppSMSTemplateUpsertOne) UpdateMessage() *AppSMSTemplateUpsertOne {
	return u.Update(func(s *AppSMSTemplateUpsert) {
		s.UpdateMessage()
	})
}

// Exec executes the query.
func (u *AppSMSTemplateUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppSMSTemplateCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppSMSTemplateUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AppSMSTemplateUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: AppSMSTemplateUpsertOne.ID is not supported by MySQL driver. Use AppSMSTemplateUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AppSMSTemplateUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AppSMSTemplateCreateBulk is the builder for creating many AppSMSTemplate entities in bulk.
type AppSMSTemplateCreateBulk struct {
	config
	builders []*AppSMSTemplateCreate
	conflict []sql.ConflictOption
}

// Save creates the AppSMSTemplate entities in the database.
func (astcb *AppSMSTemplateCreateBulk) Save(ctx context.Context) ([]*AppSMSTemplate, error) {
	specs := make([]*sqlgraph.CreateSpec, len(astcb.builders))
	nodes := make([]*AppSMSTemplate, len(astcb.builders))
	mutators := make([]Mutator, len(astcb.builders))
	for i := range astcb.builders {
		func(i int, root context.Context) {
			builder := astcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppSMSTemplateMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, astcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = astcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, astcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, astcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (astcb *AppSMSTemplateCreateBulk) SaveX(ctx context.Context) []*AppSMSTemplate {
	v, err := astcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (astcb *AppSMSTemplateCreateBulk) Exec(ctx context.Context) error {
	_, err := astcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (astcb *AppSMSTemplateCreateBulk) ExecX(ctx context.Context) {
	if err := astcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppSMSTemplate.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppSMSTemplateUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (astcb *AppSMSTemplateCreateBulk) OnConflict(opts ...sql.ConflictOption) *AppSMSTemplateUpsertBulk {
	astcb.conflict = opts
	return &AppSMSTemplateUpsertBulk{
		create: astcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppSMSTemplate.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (astcb *AppSMSTemplateCreateBulk) OnConflictColumns(columns ...string) *AppSMSTemplateUpsertBulk {
	astcb.conflict = append(astcb.conflict, sql.ConflictColumns(columns...))
	return &AppSMSTemplateUpsertBulk{
		create: astcb,
	}
}

// AppSMSTemplateUpsertBulk is the builder for "upsert"-ing
// a bulk of AppSMSTemplate nodes.
type AppSMSTemplateUpsertBulk struct {
	create *AppSMSTemplateCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AppSMSTemplate.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appsmstemplate.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppSMSTemplateUpsertBulk) UpdateNewValues() *AppSMSTemplateUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(appsmstemplate.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppSMSTemplate.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *AppSMSTemplateUpsertBulk) Ignore() *AppSMSTemplateUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppSMSTemplateUpsertBulk) DoNothing() *AppSMSTemplateUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppSMSTemplateCreateBulk.OnConflict
// documentation for more info.
func (u *AppSMSTemplateUpsertBulk) Update(set func(*AppSMSTemplateUpsert)) *AppSMSTemplateUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppSMSTemplateUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AppSMSTemplateUpsertBulk) SetCreatedAt(v uint32) *AppSMSTemplateUpsertBulk {
	return u.Update(func(s *AppSMSTemplateUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppSMSTemplateUpsertBulk) AddCreatedAt(v uint32) *AppSMSTemplateUpsertBulk {
	return u.Update(func(s *AppSMSTemplateUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppSMSTemplateUpsertBulk) UpdateCreatedAt() *AppSMSTemplateUpsertBulk {
	return u.Update(func(s *AppSMSTemplateUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppSMSTemplateUpsertBulk) SetUpdatedAt(v uint32) *AppSMSTemplateUpsertBulk {
	return u.Update(func(s *AppSMSTemplateUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppSMSTemplateUpsertBulk) AddUpdatedAt(v uint32) *AppSMSTemplateUpsertBulk {
	return u.Update(func(s *AppSMSTemplateUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppSMSTemplateUpsertBulk) UpdateUpdatedAt() *AppSMSTemplateUpsertBulk {
	return u.Update(func(s *AppSMSTemplateUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppSMSTemplateUpsertBulk) SetDeletedAt(v uint32) *AppSMSTemplateUpsertBulk {
	return u.Update(func(s *AppSMSTemplateUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppSMSTemplateUpsertBulk) AddDeletedAt(v uint32) *AppSMSTemplateUpsertBulk {
	return u.Update(func(s *AppSMSTemplateUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppSMSTemplateUpsertBulk) UpdateDeletedAt() *AppSMSTemplateUpsertBulk {
	return u.Update(func(s *AppSMSTemplateUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetAppID sets the "app_id" field.
func (u *AppSMSTemplateUpsertBulk) SetAppID(v uuid.UUID) *AppSMSTemplateUpsertBulk {
	return u.Update(func(s *AppSMSTemplateUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppSMSTemplateUpsertBulk) UpdateAppID() *AppSMSTemplateUpsertBulk {
	return u.Update(func(s *AppSMSTemplateUpsert) {
		s.UpdateAppID()
	})
}

// SetLangID sets the "lang_id" field.
func (u *AppSMSTemplateUpsertBulk) SetLangID(v uuid.UUID) *AppSMSTemplateUpsertBulk {
	return u.Update(func(s *AppSMSTemplateUpsert) {
		s.SetLangID(v)
	})
}

// UpdateLangID sets the "lang_id" field to the value that was provided on create.
func (u *AppSMSTemplateUpsertBulk) UpdateLangID() *AppSMSTemplateUpsertBulk {
	return u.Update(func(s *AppSMSTemplateUpsert) {
		s.UpdateLangID()
	})
}

// SetUsedFor sets the "used_for" field.
func (u *AppSMSTemplateUpsertBulk) SetUsedFor(v string) *AppSMSTemplateUpsertBulk {
	return u.Update(func(s *AppSMSTemplateUpsert) {
		s.SetUsedFor(v)
	})
}

// UpdateUsedFor sets the "used_for" field to the value that was provided on create.
func (u *AppSMSTemplateUpsertBulk) UpdateUsedFor() *AppSMSTemplateUpsertBulk {
	return u.Update(func(s *AppSMSTemplateUpsert) {
		s.UpdateUsedFor()
	})
}

// SetSubject sets the "subject" field.
func (u *AppSMSTemplateUpsertBulk) SetSubject(v string) *AppSMSTemplateUpsertBulk {
	return u.Update(func(s *AppSMSTemplateUpsert) {
		s.SetSubject(v)
	})
}

// UpdateSubject sets the "subject" field to the value that was provided on create.
func (u *AppSMSTemplateUpsertBulk) UpdateSubject() *AppSMSTemplateUpsertBulk {
	return u.Update(func(s *AppSMSTemplateUpsert) {
		s.UpdateSubject()
	})
}

// SetMessage sets the "message" field.
func (u *AppSMSTemplateUpsertBulk) SetMessage(v string) *AppSMSTemplateUpsertBulk {
	return u.Update(func(s *AppSMSTemplateUpsert) {
		s.SetMessage(v)
	})
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *AppSMSTemplateUpsertBulk) UpdateMessage() *AppSMSTemplateUpsertBulk {
	return u.Update(func(s *AppSMSTemplateUpsert) {
		s.UpdateMessage()
	})
}

// Exec executes the query.
func (u *AppSMSTemplateUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AppSMSTemplateCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppSMSTemplateCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppSMSTemplateUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
