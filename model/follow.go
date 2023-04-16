package model

import "douyin/config"

type Follow struct {
	ID          int64    `gorm:"primary_key"`
	FollowerID  int64    `gorm:"column:user_id"`         // 关注者ID
	FollowingID int64    `gorm:"column:to_user_id"`      // 被关注者ID
	Following   Userinfo `gorm:"foreignKey:FollowingID"` // 关联的 Following 用户
	Follower    Userinfo `gorm:"foreignKey:FollowerID"`  // 关联的 Follower 用户
}

func (f *Follow) TableName() string {
	return "follow"
}

// 新增关注
// 业务漏洞，关于查看用户信息时应该再查follow表看看有没有关注，这一点业务逻辑未实现
func FollowUser(followerID int64, followingID int64) error {
	follow := &Follow{
		FollowerID:  followerID,
		FollowingID: followingID,
	}
	return config.DB.Create(follow).Error
}

// 取消关注
func UnfollowUser(followerID, followingID int64) error {
	return config.DB.Delete(&Follow{}, "follower_id = ? AND following_id = ?", followerID, followingID).Error
}
