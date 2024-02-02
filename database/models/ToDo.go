package models

type ToDo struct {
	ID          uint   `json:"id" form:"id" gorm:"primaryKey;autoIncrement"`
	Title       string `json:"title" form:"title" gorm:"not null"`
	Description string `json:"description" form:"description" gorm:"not null"`
	Done        bool   `json:"done" form:"done" gorm:"default:false; not null"`
}

func (ToDo) TableName() string {
	return "todo"
}
