package controller

import (
	"douyin/service"
	"douyin/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 注册功能
func Register(c *gin.Context) {
	var registerReq util.RegisterLoginRequest
	//校验参数正确性
	if err := c.ShouldBind(&registerReq); err != nil {
		c.JSON(http.StatusBadRequest, util.RegisterLoginResponse{
			StatusCode: -1,
			StatusMsg:  "用户名密码格式错误，请重新输入",
		})
		return
	}

	//创建用户，返回user_id和token
	resp := service.Register(registerReq)

	//根据返回结果判断注册是否成功
	if resp.StatusCode != 0 {
		c.JSON(http.StatusOK, resp)
	} else {
		userID := resp.User_id
		token := resp.Token
		c.JSON(http.StatusOK, util.RegisterLoginResponse{
			StatusCode: 0,
			StatusMsg:  "注册成功",
			User_id:    userID,
			Token:      token,
		})
	}
}

func Login(c *gin.Context) {
	var loginReq util.RegisterLoginRequest

	//校验参数正确性
	if err := c.ShouldBind(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, util.Response{
			StatusCode: -1,
			StatusMsg:  "用户名密码格式错误，请重新输入",
		})
		return
	}

	resp := service.Login(loginReq)

	//根据返回结果判断注册是否成功
	if resp.StatusCode != 0 {
		c.JSON(http.StatusOK, resp)
	} else {
		userID := resp.User_id
		token := resp.Token
		c.JSON(http.StatusOK, util.RegisterLoginResponse{
			StatusCode: 0,
			StatusMsg:  "登录成功",
			User_id:    userID,
			Token:      token,
		})
	}
}

// 查询用户信息
func FindUserInfo(c *gin.Context) {
	var req util.UserInfoRequest

	//校验参数正确性
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, util.Response{
			StatusCode: -1,
			StatusMsg:  "请求错误",
		})
		return
	}

	//// 校验 token 是否正确
	//claims, err := util.ParseToken(req.Token)
	//if err != nil {
	//	c.JSON(http.StatusUnauthorized, util.UserInfoResponse{
	//		StatusCode: -1,
	//		StatusMsg:  "token错误",
	//	})
	//}
	//
	//// 判断 userID 是否匹配
	//if claims.UserID != req.UserId {
	//	c.JSON(http.StatusUnauthorized, util.UserInfoResponse{
	//		StatusCode: -1,
	//		StatusMsg:  "验证失败，请求非法",
	//	})
	//}

	resp := service.FindUserInfo(req)

	if resp.StatusCode != 0 {
		c.JSON(http.StatusBadRequest, resp)
	} else {
		c.JSON(http.StatusOK, resp)
	}

}
