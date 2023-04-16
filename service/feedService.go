package service

//func GetVideos() (*util.FeedResponse, error) {
//	// 获取public目录下所有视频文件
//	files, err := ioutil.ReadDir("public")
//	if err != nil {
//		return nil, err
//	}
//	// 遍历所有视频文件
//	var videos []model.Video
//	for _, f := range files {
//		// 构造Video对象
//		video := model.Video{
//			VideoID:  f.Name(),
//			PlayURL:  "http://example.com/" + f.Name(),          // 假设视频播放地址为 http://example.com/{视频文件名}
//			CoverURL: "http://example.com/" + f.Name() + ".jpg", // 假设视频封面地址为 http://example.com/{视频文件名}.jpg
//			// 其他字段根据需要自行构造
//		}
//		videos = append(videos, video)
//	}
//	// 将所有Video对象存放在FeedResponse结构体中返回
//	return &FeedResponse{
//		VideoList: videos,
//	}, nil
//}
