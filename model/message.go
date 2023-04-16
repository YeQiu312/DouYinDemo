package model

import (
	"douyin/config"
)

type Message struct {
	MessageId   int64  `gorm:"primary_key;column:message_id" json:"id"`
	ToUser      int64  `gorm:"column:to_user" json:"from_user_id"`
	FromUser    int64  `gorm:"column:from_user" json:"to_user_id"`
	Content     string `gorm:"column:content" json:"content"`
	CreatedTime string `gorm:"column:created_time" json:"create_time"`
}

func (m *Message) TableName() string {
	return "message"
}

func CreeateMessage(message Message) error {
	err := config.DB.Create(&message).Error
	return err

}
