package model

import (
	"douyin/config"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	UserID      int64     `gorm:"primary_key;column:user_id"`
	Password    string    `gorm:"column:password"`
	Username    string    `gorm:"uniqueIndex;column:username"`
	CreatedTime time.Time `gorm:"column:created_time"`
	Userinfo    Userinfo
}

func (u *User) TableName() string {
	return "user"
}

// 查询有无用户名
func CheckUsernameExists(username string) bool {
	var count int64
	if err := config.DB.Model(&User{}).Where("username = ?", username).Count(&count).Error; err != nil {
		return true
	}
	return count > 0
}

// 新增用户
func CreateUser(user *User) error {
	if err := config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// 查找有无该用户
func FindUsr(username string) (*User, error) {
	var user User
	result := config.DB.Where("username = ? ", username).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("未找到用户")
	}
	return &user, result.Error
}
