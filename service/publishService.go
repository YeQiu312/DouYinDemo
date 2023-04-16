package service

import (
	"douyin/config"
	"douyin/model"
	"douyin/util"
	"fmt"
)

func SaveUserVideo(userid int64, title string, playURL string, coverURL string) util.Response {

	//生成唯一video_id
	sf, err := util.NewSnowflake(1)
	if err != nil {
		panic(err)
	}
	id, err := sf.NextID()
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
	video_id := int64(id)

	video := &model.Video{
		VideoID:       video_id,
		PlayURL:       playURL,
		CoverURL:      coverURL,
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         title,
	}

	// 从数据库中查询上传者的信息
	userinfo := &model.Userinfo{}
	config.DB.First(userinfo, userid)
	// 将上传者的信息添加到视频的 Authors 列表中
	video.Author = userinfo

	if err := model.CreateVideoUserInfo(video); err != nil {
		return util.Response{
			StatusCode: -1,
			StatusMsg:  "投稿视频失败",
		}
	}
	return util.Response{
		StatusCode: 0,
		StatusMsg:  "投稿视频成功",
	}
}
