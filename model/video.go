package model

import "douyin/config"

type Video struct {
	VideoID       int64     `gorm:"primary_key;column:video_id" json:"id"`
	Author        *Userinfo `gorm:"many2many:videouserinfo;ForeignKey:user_id" json:"author"`
	PlayURL       string    `gorm:"column:play_url" json:"play_url"`
	CoverURL      string    `gorm:"column:cover_url" json:"cover_url"`
	FavoriteCount int64     `gorm:"column:favorite_count" json:"favorite_count"`
	CommentCount  int64     `gorm:"column:comment_count" json:"comment_count"`
	IsFavorite    bool      `gorm:"column:is_favorite" json:"is_favorite"`
	Title         string    `gorm:"column:title" json:"title"`
}

func (v *Video) TableName() string {
	return "video"
}

type VideoUserInfo struct {
	UserInfoUserID int64 `gorm:"primaryKey;column:userinfo_user_id"`
	VideoVideoID   int64 `gorm:"primaryKey;column:video_video_id"`
}

func (vu *VideoUserInfo) TableName() string {
	return "videouserinfo"
}

func CreateVideoUserInfo(video *Video) error {
	// 先将 Author 设为 nil
	author := video.Author
	video.Author = nil

	// 保存 video 到 video 表中
	if err := config.DB.Save(video).Error; err != nil {
		return err
	}

	// 恢复 Author
	video.Author = author

	// 创建或更新 videouserinfo 表中的记录
	videouserinfo := &VideoUserInfo{
		UserInfoUserID: video.Author.UserID,
		VideoVideoID:   video.VideoID,
	}
	if err := config.DB.Save(videouserinfo).Error; err != nil {
		return err
	}

	return nil
}
