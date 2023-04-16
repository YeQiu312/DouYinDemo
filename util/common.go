package util

import (
	"douyin/model"
)

//返回信息工具类
// 状态码，0-成功，其他值-失败

type Response struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type RegisterLoginResponse struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
	User_id    int64  `json:"user_id,omitempty"`
	Token      string `json:"token,omitempty"`
}

type RegisterLoginRequest struct {
	Username string `form:"username" binding:"required,max=32"`
	Password string `form:"password" binding:"required,max=32"`
}

type UserInfoRequest struct {
	UserId int64  `form:"user_id" binding:"required"`
	Token  string `form:"token" binding:"required"`
}

type UserInfoResponse struct {
	StatusCode int64             `json:"status_code"`
	StatusMsg  string            `json:"status_msg,omitempty"`
	Userinfo   *(model.Userinfo) `json:"userinfo"`
}

type VideoListResponse struct {
	Response
	VideoList []model.Video `json:"video_list"`
}

type FeedResponse struct {
	Response
	VideoList []model.Video `json:"video_list,omitempty"`
	NextTime  int64         `json:"next_time,omitempty"`
}

type FollowRequest struct {
	ToUserId   string `form:"to_user_id" binding:"required"`
	ActionType string `form:"action_type" binding:"required"`
}

type FollowListResponse struct {
	Response
	UserList []model.Userinfo `json:"user_list,omitempty"`
}

type SendMessageRequest struct {
	ToUserId   int64  `form:"to_user_id" binding:"required"`
	ActionType string `form:"action_type" binding:"required"`
	Content    string `form:"content" binding:"required"`
}

type ChatMessageRequest struct {
	ToUserId int64 `form:"to_user_id" binding:"required"`
}

type ChatMessageResonse struct {
	Response
	MessageList []model.Message `json:"message_list,omitempty"`
}
