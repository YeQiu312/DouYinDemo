package model

import "time"

type Relation struct {
	RelationId  uint      `gorm:"primary_key;column:relation_id"`
	FromUser    uint      `gorm:"column:from_user"`
	ToUser      uint      `gorm:"column:to_user"`
	Relation    uint      `gorm:"column:relation"`
	CreatedTime time.Time `gorm:"column:created_time"`
}

func (r *Relation) TableName() string {
	return "relation"
}
