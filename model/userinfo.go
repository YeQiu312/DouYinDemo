package model

import (
	"douyin/config"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Userinfo struct {
	UserID          int64    `gorm:"primary_key;not null;column:user_id" json:"id"`             // 用户id
	Username        string   `gorm:"not null;column:username" json:"name"`                      // 用户名称
	FollowCount     int64    `gorm:"not null;column:follow_count"  json:"follow_count"`         // 关注总数
	FollowerCount   int64    `gorm:"not null;column:follower_count"  json:"follower_count"`     // 粉丝总数
	IsFollow        bool     `gorm:"not null;column:is_follow"  json:"is_follow"`               // true-已关注，false-未关注
	Avatar          string   `gorm:"not null;column:avatar"  json:"avatar"`                     // 用户头像
	BackgroundImage string   `gorm:"not null;column:background_image"  json:"background_image"` // 用户个人页顶部大图
	Signature       string   `gorm:"not null;column:signature"  json:"signature"`               // 个人简介
	TotalFavorited  int64    `gorm:"not null;column:total_favorited"  json:"total_favorited"`   // 获赞数量
	WorkCount       int64    `gorm:"not null;column:work_count" json:"work_count"`              // 作品数
	FavoriteCount   int64    `gorm:"not null;column:favorite_count" json:"favorite_count"`      // 喜欢数
	Videos          []*Video `gorm:"many2many:videouserinfo;" json:"-"`
}

func (userinfo *Userinfo) TableName() string {
	return "userinfo"
}

func CreateUserInfo(userinfo *Userinfo) error {
	if err := config.DB.Create(userinfo).Error; err != nil {
		return err
	}
	return nil
}

func FindUserInfo(userid int64) (*Userinfo, error) {
	var userinfo Userinfo
	result := config.DB.Where("user_id = ?", userid).First(&userinfo)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("未找到用户信息")
	}
	return &userinfo, result.Error
}
