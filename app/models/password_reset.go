package models

import (
	"echo_shop/pkg/utils"
	"time"

	"github.com/jinzhu/gorm"
)

const (
	// PasswordResetTableName : password_reset table name
	PasswordResetTableName = "password_resets"
)

// PasswordReset 密码重置 model
type PasswordReset struct {
	Email     string `sql:"not null; index"`
	Token     string `sql:"not null; index"`
	CreatedAt time.Time
}

// TableName 表名
func (PasswordReset) TableName() string {
	return PasswordResetTableName
}

// BeforeCreate - hook
func (p *PasswordReset) BeforeCreate(scope *gorm.Scope) (err error) {
	token := string(utils.RandomCreateBytes(30))
	scope.SetColumn("token", token)
	return nil
}
