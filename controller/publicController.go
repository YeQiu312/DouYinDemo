package controller

import (
	"douyin/config"
	"douyin/model"
	"douyin/service"
	"douyin/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path"
	"path/filepath"
)

// 发布列表
func List(c *gin.Context) {
	userid, _ := c.Get("userID")

	var videoList []model.Video
	err := config.DB.Model(&model.Video{}).
		Joins("left join videouserinfo on videouserinfo.video_video_id = video.video_id").
		Where("videouserinfo.userinfo_user_id = ?", userid.(int64)).
		Find(&videoList).Error

	userinfo := &model.Userinfo{}
	config.DB.First(userinfo, userid.(int64))
	// 将上传者的信息添加到视频的 Authors 列表中
	for i, _ := range videoList {
		videoList[i].Author = userinfo
	}

	// 错误处理
	if err != nil {
		c.JSON(http.StatusOK, util.VideoListResponse{
			Response: util.Response{
				StatusCode: 0,
				StatusMsg:  "获取发布列表失败",
			},
		})
		return
	}

	c.JSON(http.StatusOK, util.VideoListResponse{
		Response: util.Response{
			StatusCode: 0,
			StatusMsg:  "获取发布列表成功",
		},
		VideoList: videoList,
	})
}

// 投稿接口
func Action(c *gin.Context) {
	//获取在中间件中获得的userid
	userid, _ := c.Get("userID")
	title := c.PostForm("title")

	//处理发送的视频文件
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, util.Response{
			StatusCode: -1,
			StatusMsg:  "用户文件上传失败" + err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", userid, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, util.Response{
			StatusCode: -1,
			StatusMsg:  "用户文件保存失败" + err.Error(),
		})
		return
	}

	// 截取封面
	PlayURL := "./public/" + finalName
	fileExt := path.Ext(filename)                       // 获取文件后缀，返回 ".mp4"
	filePrefix := filename[:len(filename)-len(fileExt)] // 获取文件名前缀，返回 "example"
	coverPrefix := fmt.Sprintf("%d_%s", userid, filePrefix)
	coverName := coverPrefix + ".jpg"
	CoverURL := "./cover/" + coverName
	err = util.ExtractVideoCover(PlayURL, CoverURL)
	if err != nil {
		log.Fatal(err)
	}

	//后续使用云端存储这里记得改
	RealPlayURL := "http://localhost:8080/public/" + finalName
	RealCoverURL := "http://localhost:8080/cover/" + coverName
	resp := service.SaveUserVideo(userid.(int64), title, RealPlayURL, RealCoverURL)

	if resp.StatusCode != 0 {
		c.JSON(http.StatusOK, resp)
	}
	c.JSON(http.StatusOK, util.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " 上传成功",
	})
}
