package service

import (
	"douyin/model"
	"douyin/util"
	"fmt"
	"time"
)

type RegisterLoginResponse struct {
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

func Register(registerReq util.RegisterLoginRequest) util.RegisterLoginResponse {

	//检查用户名是否唯一
	if model.CheckUsernameExists(registerReq.Username) {
		return util.RegisterLoginResponse{
			StatusCode: -1,
			StatusMsg:  "用户名已存在",
		}
	}

	//生成唯一id-
	sf, err := util.NewSnowflake(1)
	if err != nil {
		panic(err)
	}
	id, err := sf.NextID()
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
	user_id := int64(id)

	//密码加密
	hashedPassword, err := util.HashPassword(registerReq.Password)
	if err != nil {
		return util.RegisterLoginResponse{
			StatusCode: -1,
			StatusMsg:  "密码加密失败",
		}
	}

	user := model.User{
		Username:    registerReq.Username,
		Password:    hashedPassword,
		UserID:      user_id,
		CreatedTime: time.Now(),
	}

	userinfo := model.Userinfo{
		Username:        registerReq.Username,
		UserID:          user_id,
		FollowCount:     0,
		FollowerCount:   0,
		IsFollow:        true,
		Avatar:          "",
		BackgroundImage: "",
		Signature:       "",
		TotalFavorited:  0,
		WorkCount:       0,
		FavoriteCount:   0,
	}

	//新增用户数据
	if err := model.CreateUser(&user); err != nil {
		return util.RegisterLoginResponse{
			StatusCode: -1,
			StatusMsg:  "新增用户数据失败",
		}
	}

	//新增用户信息
	if err := model.CreateUserInfo(&userinfo); err != nil {
		return util.RegisterLoginResponse{
			StatusCode: -1,
			StatusMsg:  "新增用户信息失败",
		}
	}

	//生成JWT token
	token, err := util.GenerateJWT(id)
	if err != nil {
		return util.RegisterLoginResponse{
			StatusCode: -1,
			StatusMsg:  "token生成失败",
		}
	}

	//返回用户id和token
	return util.RegisterLoginResponse{
		StatusCode: 0,
		StatusMsg:  "注册成功",
		User_id:    user_id,
		Token:      token,
	}
}

func Login(loginReq util.RegisterLoginRequest) util.RegisterLoginResponse {
	///查询是否有该用户
	user, err := model.FindUsr(loginReq.Username)
	if err != nil {
		return util.RegisterLoginResponse{
			StatusCode: -1,
			StatusMsg:  "用户不存在",
		}
	}

	// 校验密码
	if !util.CheckPasswordHash(loginReq.Password, user.Password) {
		return util.RegisterLoginResponse{
			StatusCode: -1,
			StatusMsg:  "密码错误",
		}
	}

	//生成JWT token
	token, err := util.GenerateJWT(user.UserID)
	if err != nil {
		return util.RegisterLoginResponse{
			StatusCode: -1,
			StatusMsg:  "token生成失败",
		}
	}

	//返回用户id和token
	return util.RegisterLoginResponse{
		StatusCode: 0,
		StatusMsg:  "注册成功",
		User_id:    user.UserID,
		Token:      token,
	}
}

func FindUserInfo(userinfo util.UserInfoRequest) util.UserInfoResponse {

	//根据user_id查询用户信息
	info, err := model.FindUserInfo(userinfo.UserId)
	if err != nil {
		return util.UserInfoResponse{
			StatusCode: -1,
			StatusMsg:  "用户信息查询失败",
		}
	}

	return util.UserInfoResponse{
		StatusCode: 0,
		StatusMsg:  "用户信息查询成功",
		Userinfo:   info,
	}

}
