package controller

import (
	"douyin/config"
	"douyin/model"
	"douyin/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 发送消息
func SendMessage(c *gin.Context) {
	userid, _ := c.Get("userID")

	var req util.SendMessageRequest

	//参数检验
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, util.Response{
			StatusCode: -1,
			StatusMsg:  "请求失败",
		})
		return
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
	messageId := int64(id)

	//生成当前时间
	now := time.Now()
	currentTime := now.Format("2006-01-02 15:04:05")

	newMessage := model.Message{
		MessageId:   messageId,
		ToUser:      req.ToUserId,
		FromUser:    userid.(int64),
		Content:     req.Content,
		CreatedTime: currentTime,
	}

	err = model.CreeateMessage(newMessage)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.Response{
			StatusCode: -1,
			StatusMsg:  "发送消息失败",
		})
		return
	} else {
		c.JSON(http.StatusOK, util.Response{
			StatusCode: 0,
			StatusMsg:  "发送消息成功",
		})
	}

}

// 聊天记录
func ChatMessage(c *gin.Context) {
	userid, _ := c.Get("userID")

	var req util.ChatMessageRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, util.Response{
			StatusCode: -1,
			StatusMsg:  "请求失败",
		})
		return
	}
	to_user_id := req.ToUserId

	var message []model.Message
	err := config.DB.Where("from_user = ? AND to_user = ?", userid.(int64), to_user_id).Order("created_time DESC").Find(&message).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, util.ChatMessageResonse{
			Response: util.Response{
				StatusCode: -1,
				StatusMsg:  "查找聊天记录失败",
			},
		})
	} else {
		c.JSON(http.StatusOK, util.ChatMessageResonse{
			Response: util.Response{
				StatusCode: -1,
				StatusMsg:  "查找聊天记录成功",
			},
			MessageList: message,
		})
	}
}
