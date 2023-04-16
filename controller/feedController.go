package controller

import (
	"douyin/config"
	"douyin/model"
	"douyin/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 视频流接口
func Feed(c *gin.Context) {
	var videos []model.Video
	if err := config.DB.Preload("Author").Find(&videos).Error; err != nil {
		log.Println("截取视频数据出错：", err.Error())
		c.JSON(http.StatusOK, util.VideoListResponse{
			Response: util.Response{
				StatusCode: 0,
				StatusMsg:  "获取发布列表失败",
			},
		})
		return
	}

	//这里是要改进的，就是错误处理方面
	for i, video := range videos {
		userinfo := &model.Userinfo{}
		config.DB.Model(&model.Video{}).
			Joins("left join videouserinfo on videouserinfo.userinfo_user_id = userinfo.user_id").
			Where("videouserinfo.video_video_id = ?", video.VideoID).
			Find(&userinfo)
		videos[i].Author = userinfo
	}

	c.JSON(http.StatusOK, util.VideoListResponse{
		Response: util.Response{
			StatusCode: 0,
			StatusMsg:  "获取发布列表成功",
		},
		VideoList: videos,
	})
}
