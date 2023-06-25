package model

import (
	"time"
)

type UserInfo struct {
	ID          int        `gorm:"column:id;not null;auto_increment;primary_key" `
	UserName    string     `gorm:"column:user_name;not null;type:varchar(128);comment:'用户昵称'" `
	UserId      string     `gorm:"column:user_id;not null;type:varchar(50);comment:'用户id'" `
	PassWord    string     `gorm:"column:pass_word;not null;type:varchar(128);comment:'密码'" `
	Avatar      string     `gorm:"column:avatar;not null;type:varchar(256);comment:'用户头像地址'" `
	Email       string     `gorm:"column:email;not null;type:varchar(128);comment:'用户邮箱'" `
	Level       int8       `gorm:"column:level;not null;comment:'用户等级'" `
	LevelName   string     `gorm:"column:level_name;not null;type:varchar(50);comment:'用户等级名'" `
	Phone       int64      `gorm:"column:phone;not null;comment:'手机号'" `
	LastLoginAt time.Time  `gorm:"column:last_login_at" `
	LastLoginIp string     `gorm:"column:last_login_ip;not null;type:varchar(50);comment:'最后登录ip'" `
	CreatedIp   string     `gorm:"column:created_ip;not null;type:varchar(50);comment:'创建地ip'" `
	CreatedAt   time.Time  `gorm:"column:created_at;default:CURRENT_TIMESTAMP" `
	UpdatedAt   time.Time  `gorm:"column:updated_at;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" `
	DeletedAt   *time.Time `sql:"index" json:"deleted_at"`
}

type UserProfile struct {
}
