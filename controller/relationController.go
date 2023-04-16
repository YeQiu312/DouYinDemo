package controller

import (
	"douyin/model"
	"douyin/service"
	"douyin/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 关注操作
func Follow(c *gin.Context) {
	userid, _ := c.Get("userID")

	//获取 to_user_id 参数
	//touserid := c.Request.FormValue("to_user_id")

	// 获取 action_type 参数
	//actionType := c.Request.FormValue("action_type")

	var req util.FollowRequest
	if err := c.BindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, util.Response{
			StatusCode: -1,
			StatusMsg:  "请求失败",
		})
		return
	}

	touserid, err := strconv.ParseInt(req.ToUserId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.Response{
			StatusCode: -1,
			StatusMsg:  "关注失败",
		})
		return
	}

	actionType := req.ActionType
	if actionType == "1" {
		if err := model.FollowUser(userid.(int64), touserid); err != nil {
			c.JSON(http.StatusBadRequest, util.Response{
				StatusCode: -1,
				StatusMsg:  "关注失败",
			})
			return
		} else {
			c.JSON(http.StatusOK, util.Response{
				StatusCode: 0,
				StatusMsg:  "关注成功",
			})
			return
		}
	}

	if actionType == "2" {
		if err := model.UnfollowUser(userid.(int64), touserid); err != nil {
			c.JSON(http.StatusBadRequest, util.Response{
				StatusCode: -1,
				StatusMsg:  "取关失败",
			})
			return
		} else {
			c.JSON(http.StatusOK, util.Response{
				StatusCode: 0,
				StatusMsg:  "取关成功",
			})
			return
		}
	}

}

// 关注列表
func FollowList(c *gin.Context) {
	//content获取userid
	userid, _ := c.Get("userID")

	resp := service.FollowList(userid.(int64))
	if resp.StatusCode != 0 {
		c.JSON(http.StatusBadRequest, resp)
		return
	} else {
		c.JSON(http.StatusOK, resp)
	}

}

// 粉丝列表
func FollowerList(c *gin.Context) {
	//content获取userid
	userid, _ := c.Get("userID")

	resp := service.FollowerList(userid.(int64))
	if resp.StatusCode != 0 {
		c.JSON(http.StatusBadRequest, resp)
		return
	} else {
		c.JSON(http.StatusOK, resp)
	}

}

// 好友列表
func FriendList(c *gin.Context) {
	//content获取userid
	userid, _ := c.Get("userID")

	resp := service.FollowList(userid.(int64))
	if resp.StatusCode != 0 {
		c.JSON(http.StatusBadRequest, resp)
		return
	} else {
		c.JSON(http.StatusOK, resp)
	}
}
