// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/third-manager/pkg/db/ent/appcontact"
	"github.com/NpoolPlatform/third-manager/pkg/db/ent/appemailtemplate"
	"github.com/NpoolPlatform/third-manager/pkg/db/ent/appsmstemplate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 3)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   appcontact.Table,
			Columns: appcontact.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appcontact.FieldID,
			},
		},
		Type: "AppContact",
		Fields: map[string]*sqlgraph.FieldSpec{
			appcontact.FieldCreatedAt:   {Type: field.TypeUint32, Column: appcontact.FieldCreatedAt},
			appcontact.FieldUpdatedAt:   {Type: field.TypeUint32, Column: appcontact.FieldUpdatedAt},
			appcontact.FieldDeletedAt:   {Type: field.TypeUint32, Column: appcontact.FieldDeletedAt},
			appcontact.FieldAppID:       {Type: field.TypeUUID, Column: appcontact.FieldAppID},
			appcontact.FieldUsedFor:     {Type: field.TypeString, Column: appcontact.FieldUsedFor},
			appcontact.FieldSender:      {Type: field.TypeString, Column: appcontact.FieldSender},
			appcontact.FieldAccount:     {Type: field.TypeString, Column: appcontact.FieldAccount},
			appcontact.FieldAccountType: {Type: field.TypeString, Column: appcontact.FieldAccountType},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   appemailtemplate.Table,
			Columns: appemailtemplate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appemailtemplate.FieldID,
			},
		},
		Type: "AppEmailTemplate",
		Fields: map[string]*sqlgraph.FieldSpec{
			appemailtemplate.FieldCreatedAt:         {Type: field.TypeUint32, Column: appemailtemplate.FieldCreatedAt},
			appemailtemplate.FieldUpdatedAt:         {Type: field.TypeUint32, Column: appemailtemplate.FieldUpdatedAt},
			appemailtemplate.FieldDeletedAt:         {Type: field.TypeUint32, Column: appemailtemplate.FieldDeletedAt},
			appemailtemplate.FieldAppID:             {Type: field.TypeUUID, Column: appemailtemplate.FieldAppID},
			appemailtemplate.FieldLangID:            {Type: field.TypeUUID, Column: appemailtemplate.FieldLangID},
			appemailtemplate.FieldDefaultToUsername: {Type: field.TypeString, Column: appemailtemplate.FieldDefaultToUsername},
			appemailtemplate.FieldUsedFor:           {Type: field.TypeString, Column: appemailtemplate.FieldUsedFor},
			appemailtemplate.FieldSender:            {Type: field.TypeString, Column: appemailtemplate.FieldSender},
			appemailtemplate.FieldReplyTos:          {Type: field.TypeJSON, Column: appemailtemplate.FieldReplyTos},
			appemailtemplate.FieldCcTos:             {Type: field.TypeJSON, Column: appemailtemplate.FieldCcTos},
			appemailtemplate.FieldSubject:           {Type: field.TypeString, Column: appemailtemplate.FieldSubject},
			appemailtemplate.FieldBody:              {Type: field.TypeString, Column: appemailtemplate.FieldBody},
		},
	}
	graph.Nodes[2] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   appsmstemplate.Table,
			Columns: appsmstemplate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appsmstemplate.FieldID,
			},
		},
		Type: "AppSMSTemplate",
		Fields: map[string]*sqlgraph.FieldSpec{
			appsmstemplate.FieldCreatedAt: {Type: field.TypeUint32, Column: appsmstemplate.FieldCreatedAt},
			appsmstemplate.FieldUpdatedAt: {Type: field.TypeUint32, Column: appsmstemplate.FieldUpdatedAt},
			appsmstemplate.FieldDeletedAt: {Type: field.TypeUint32, Column: appsmstemplate.FieldDeletedAt},
			appsmstemplate.FieldAppID:     {Type: field.TypeUUID, Column: appsmstemplate.FieldAppID},
			appsmstemplate.FieldLangID:    {Type: field.TypeUUID, Column: appsmstemplate.FieldLangID},
			appsmstemplate.FieldUsedFor:   {Type: field.TypeString, Column: appsmstemplate.FieldUsedFor},
			appsmstemplate.FieldSubject:   {Type: field.TypeString, Column: appsmstemplate.FieldSubject},
			appsmstemplate.FieldMessage:   {Type: field.TypeString, Column: appsmstemplate.FieldMessage},
		},
	}
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (acq *AppContactQuery) addPredicate(pred func(s *sql.Selector)) {
	acq.predicates = append(acq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the AppContactQuery builder.
func (acq *AppContactQuery) Filter() *AppContactFilter {
	return &AppContactFilter{config: acq.config, predicateAdder: acq}
}

// addPredicate implements the predicateAdder interface.
func (m *AppContactMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the AppContactMutation builder.
func (m *AppContactMutation) Filter() *AppContactFilter {
	return &AppContactFilter{config: m.config, predicateAdder: m}
}

// AppContactFilter provides a generic filtering capability at runtime for AppContactQuery.
type AppContactFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *AppContactFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *AppContactFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(appcontact.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *AppContactFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(appcontact.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *AppContactFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(appcontact.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *AppContactFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(appcontact.FieldDeletedAt))
}

// WhereAppID applies the entql [16]byte predicate on the app_id field.
func (f *AppContactFilter) WhereAppID(p entql.ValueP) {
	f.Where(p.Field(appcontact.FieldAppID))
}

// WhereUsedFor applies the entql string predicate on the used_for field.
func (f *AppContactFilter) WhereUsedFor(p entql.StringP) {
	f.Where(p.Field(appcontact.FieldUsedFor))
}

// WhereSender applies the entql string predicate on the sender field.
func (f *AppContactFilter) WhereSender(p entql.StringP) {
	f.Where(p.Field(appcontact.FieldSender))
}

// WhereAccount applies the entql string predicate on the account field.
func (f *AppContactFilter) WhereAccount(p entql.StringP) {
	f.Where(p.Field(appcontact.FieldAccount))
}

// WhereAccountType applies the entql string predicate on the account_type field.
func (f *AppContactFilter) WhereAccountType(p entql.StringP) {
	f.Where(p.Field(appcontact.FieldAccountType))
}

// addPredicate implements the predicateAdder interface.
func (aetq *AppEmailTemplateQuery) addPredicate(pred func(s *sql.Selector)) {
	aetq.predicates = append(aetq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the AppEmailTemplateQuery builder.
func (aetq *AppEmailTemplateQuery) Filter() *AppEmailTemplateFilter {
	return &AppEmailTemplateFilter{config: aetq.config, predicateAdder: aetq}
}

// addPredicate implements the predicateAdder interface.
func (m *AppEmailTemplateMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the AppEmailTemplateMutation builder.
func (m *AppEmailTemplateMutation) Filter() *AppEmailTemplateFilter {
	return &AppEmailTemplateFilter{config: m.config, predicateAdder: m}
}

// AppEmailTemplateFilter provides a generic filtering capability at runtime for AppEmailTemplateQuery.
type AppEmailTemplateFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *AppEmailTemplateFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *AppEmailTemplateFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(appemailtemplate.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *AppEmailTemplateFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(appemailtemplate.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *AppEmailTemplateFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(appemailtemplate.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *AppEmailTemplateFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(appemailtemplate.FieldDeletedAt))
}

// WhereAppID applies the entql [16]byte predicate on the app_id field.
func (f *AppEmailTemplateFilter) WhereAppID(p entql.ValueP) {
	f.Where(p.Field(appemailtemplate.FieldAppID))
}

// WhereLangID applies the entql [16]byte predicate on the lang_id field.
func (f *AppEmailTemplateFilter) WhereLangID(p entql.ValueP) {
	f.Where(p.Field(appemailtemplate.FieldLangID))
}

// WhereDefaultToUsername applies the entql string predicate on the default_to_username field.
func (f *AppEmailTemplateFilter) WhereDefaultToUsername(p entql.StringP) {
	f.Where(p.Field(appemailtemplate.FieldDefaultToUsername))
}

// WhereUsedFor applies the entql string predicate on the used_for field.
func (f *AppEmailTemplateFilter) WhereUsedFor(p entql.StringP) {
	f.Where(p.Field(appemailtemplate.FieldUsedFor))
}

// WhereSender applies the entql string predicate on the sender field.
func (f *AppEmailTemplateFilter) WhereSender(p entql.StringP) {
	f.Where(p.Field(appemailtemplate.FieldSender))
}

// WhereReplyTos applies the entql json.RawMessage predicate on the reply_tos field.
func (f *AppEmailTemplateFilter) WhereReplyTos(p entql.BytesP) {
	f.Where(p.Field(appemailtemplate.FieldReplyTos))
}

// WhereCcTos applies the entql json.RawMessage predicate on the cc_tos field.
func (f *AppEmailTemplateFilter) WhereCcTos(p entql.BytesP) {
	f.Where(p.Field(appemailtemplate.FieldCcTos))
}

// WhereSubject applies the entql string predicate on the subject field.
func (f *AppEmailTemplateFilter) WhereSubject(p entql.StringP) {
	f.Where(p.Field(appemailtemplate.FieldSubject))
}

// WhereBody applies the entql string predicate on the body field.
func (f *AppEmailTemplateFilter) WhereBody(p entql.StringP) {
	f.Where(p.Field(appemailtemplate.FieldBody))
}

// addPredicate implements the predicateAdder interface.
func (astq *AppSMSTemplateQuery) addPredicate(pred func(s *sql.Selector)) {
	astq.predicates = append(astq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the AppSMSTemplateQuery builder.
func (astq *AppSMSTemplateQuery) Filter() *AppSMSTemplateFilter {
	return &AppSMSTemplateFilter{config: astq.config, predicateAdder: astq}
}

// addPredicate implements the predicateAdder interface.
func (m *AppSMSTemplateMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the AppSMSTemplateMutation builder.
func (m *AppSMSTemplateMutation) Filter() *AppSMSTemplateFilter {
	return &AppSMSTemplateFilter{config: m.config, predicateAdder: m}
}

// AppSMSTemplateFilter provides a generic filtering capability at runtime for AppSMSTemplateQuery.
type AppSMSTemplateFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *AppSMSTemplateFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[2].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *AppSMSTemplateFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(appsmstemplate.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *AppSMSTemplateFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(appsmstemplate.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *AppSMSTemplateFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(appsmstemplate.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *AppSMSTemplateFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(appsmstemplate.FieldDeletedAt))
}

// WhereAppID applies the entql [16]byte predicate on the app_id field.
func (f *AppSMSTemplateFilter) WhereAppID(p entql.ValueP) {
	f.Where(p.Field(appsmstemplate.FieldAppID))
}

// WhereLangID applies the entql [16]byte predicate on the lang_id field.
func (f *AppSMSTemplateFilter) WhereLangID(p entql.ValueP) {
	f.Where(p.Field(appsmstemplate.FieldLangID))
}

// WhereUsedFor applies the entql string predicate on the used_for field.
func (f *AppSMSTemplateFilter) WhereUsedFor(p entql.StringP) {
	f.Where(p.Field(appsmstemplate.FieldUsedFor))
}

// WhereSubject applies the entql string predicate on the subject field.
func (f *AppSMSTemplateFilter) WhereSubject(p entql.StringP) {
	f.Where(p.Field(appsmstemplate.FieldSubject))
}

// WhereMessage applies the entql string predicate on the message field.
func (f *AppSMSTemplateFilter) WhereMessage(p entql.StringP) {
	f.Where(p.Field(appsmstemplate.FieldMessage))
}
