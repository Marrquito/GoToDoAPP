package models

import "time"

type ToDo struct {
	ID          uint      `json:"id" form:"id" gorm:"primaryKey;autoIncrement"`
	Title       string    `json:"title" form:"title" gorm:"not null"`
	Description string    `json:"description" form:"description" gorm:"not null"`
	Done        bool      `json:"done" form:"done" gorm:"default:false;"`
	CreatedAt   time.Time `json:"create_at" form:"create_at" gorm:"autoUpdateTime"`
	UpdatedAt   time.Time `json:"updated_at" form:"updated_at" gorm:"autoUpdateTime"`
}

func (ToDo) TableName() string {
	return "todo"
}
