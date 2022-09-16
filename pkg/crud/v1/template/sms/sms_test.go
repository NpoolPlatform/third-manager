package sms

import (
	"context"
	"fmt"

	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/third-manager/pkg/db/ent"
	"github.com/NpoolPlatform/third-manager/pkg/testinit"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/message/npool"
	"github.com/NpoolPlatform/message/npool/third/mgr/v1/template/sms"
	"github.com/NpoolPlatform/message/npool/third/mgr/v1/usedfor"

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var entAppSMSTemplate = ent.AppSMSTemplate{
	ID:      uuid.New(),
	AppID:   uuid.New(),
	LangID:  uuid.New(),
	UsedFor: usedfor.UsedFor_Signin.String(),
	Subject: uuid.NewString(),
	Message: uuid.NewString(),
}

var (
	id             = entAppSMSTemplate.ID.String()
	appID          = entAppSMSTemplate.AppID.String()
	langID         = entAppSMSTemplate.LangID.String()
	usedFor        = usedfor.UsedFor_Signin
	appSMSTemplate = sms.SMSTemplateReq{
		ID:      &id,
		AppID:   &appID,
		LangID:  &langID,
		UsedFor: &usedFor,
		Subject: &entAppSMSTemplate.Subject,
		Message: &entAppSMSTemplate.Message,
	}
)

var info *ent.AppSMSTemplate

func rowToObject(row *ent.AppSMSTemplate) *ent.AppSMSTemplate {
	return &ent.AppSMSTemplate{
		ID:        row.ID,
		CreatedAt: row.CreatedAt,
		AppID:     row.AppID,
		LangID:    row.LangID,
		UsedFor:   row.UsedFor,
		Subject:   row.Subject,
		Message:   row.Message,
	}
}

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &appSMSTemplate)
	if assert.Nil(t, err) {
		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
			entAppSMSTemplate.ID = info.ID
			entAppSMSTemplate.CreatedAt = info.CreatedAt
		}
		assert.Equal(t, rowToObject(info), &entAppSMSTemplate)
	}
}

func createBulk(t *testing.T) {
	entApp := []ent.AppSMSTemplate{
		{
			ID:      uuid.New(),
			AppID:   uuid.New(),
			LangID:  uuid.New(),
			UsedFor: usedfor.UsedFor_Signin.String(),
			Subject: uuid.NewString(),
			Message: uuid.NewString(),
		},
		{
			ID:      uuid.New(),
			AppID:   uuid.New(),
			LangID:  uuid.New(),
			UsedFor: usedfor.UsedFor_Signin.String(),
			Subject: uuid.NewString(),
			Message: uuid.NewString(),
		},
	}

	apps := []*sms.SMSTemplateReq{}
	for key := range entApp {
		id := entApp[key].ID.String()
		appID = entApp[key].AppID.String()
		langID := entApp[key].LangID.String()
		usedFor = usedfor.UsedFor_Signin
		apps = append(apps, &sms.SMSTemplateReq{
			ID:      &id,
			AppID:   &appID,
			LangID:  &langID,
			UsedFor: &usedFor,
			Subject: &entApp[key].Subject,
			Message: &entApp[key].Message,
		})
	}
	infos, err := CreateBulk(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
		assert.NotEqual(t, infos[0].ID, uuid.UUID{}.String())
		assert.NotEqual(t, infos[1].ID, uuid.UUID{}.String())
	}
}

func update(t *testing.T) {
	var err error
	info, err = Update(context.Background(), &appSMSTemplate)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entAppSMSTemplate)
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entAppSMSTemplate)
	}
}

func rows(t *testing.T) {
	infos, total, err := Rows(context.Background(),
		&sms.Conds{
			ID: &npool.StringVal{
				Value: info.ID.String(),
				Op:    cruder.EQ,
			},
		}, 0, 0)
	if assert.Nil(t, err) {
		assert.Equal(t, total, 1)
		assert.Equal(t, rowToObject(infos[0]), &entAppSMSTemplate)
	}
}

func rowOnly(t *testing.T) {
	var err error
	info, err = RowOnly(context.Background(),
		&sms.Conds{
			ID: &npool.StringVal{
				Value: info.ID.String(),
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entAppSMSTemplate)
	}
}

func count(t *testing.T) {
	count, err := Count(context.Background(),
		&sms.Conds{
			ID: &npool.StringVal{
				Value: info.ID.String(),
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, count, count)
	}
}

func exist(t *testing.T) {
	exist, err := Exist(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existConds(t *testing.T) {
	exist, err := ExistConds(context.Background(),
		&sms.Conds{
			ID: &npool.StringVal{
				Value: info.ID.String(),
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteT(t *testing.T) {
	info, err := Delete(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entAppSMSTemplate)
	}
}

func TestMainOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("create", create)
	t.Run("createBulk", createBulk)
	t.Run("row", row)
	t.Run("rows", rows)
	t.Run("rowOnly", rowOnly)
	t.Run("update", update)
	t.Run("exist", exist)
	t.Run("existConds", existConds)
	t.Run("delete", deleteT)
	t.Run("count", count)
}
