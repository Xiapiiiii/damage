package model

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint           `gorm:"primarykey"`
	CreatedAt time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP(3)" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type BaseModel struct {
	db        *gorm.DB
	tableName string
}

func (b *BaseModel) BeginTx() *gorm.DB {
	return b.db.Begin()
}

func (b *BaseModel) getDB(db []*gorm.DB) *gorm.DB {
	if len(db) > 0 {
		return db[0]
	}
	return b.db
}
