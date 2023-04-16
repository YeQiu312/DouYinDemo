package service

import (
	"douyin/config"
	"douyin/model"
	"douyin/util"
)

func FollowList(userid int64) util.FollowListResponse {
	var followers []model.Follow
	var userinfos []model.Userinfo

	// 查询所有关注记录，并预加载关联的Following用户信息
	if err := config.DB.Preload("Following").Where("user_id = ?", userid).Find(&followers).Error; err != nil {
		return util.FollowListResponse{
			Response: util.Response{
				StatusCode: -1,
				StatusMsg:  "查询关注列表失败",
			},
		}
	}

	// 遍历所有Follow记录，获取关联的Following用户信息
	for _, f := range followers {
		userinfos = append(userinfos, f.Following)
	}

	return util.FollowListResponse{
		Response: util.Response{
			StatusCode: 0,
			StatusMsg:  "查询关注列表成功",
		},
		UserList: userinfos,
	}
}

func FollowerList(userid int64) util.FollowListResponse {
	var followings []model.Follow
	var userinfos []model.Userinfo

	// 根据FollowerID查询所有Follow记录
	if err := config.DB.Preload("Follower").Where("to_user_id = ?", userid).Find(&followings).Error; err != nil {
		return util.FollowListResponse{
			Response: util.Response{
				StatusCode: -1,
				StatusMsg:  "查询粉丝列表失败",
			},
		}
	}

	// 遍历所有Follow记录，获取关联的Following用户信息
	for _, f := range followings {
		userinfos = append(userinfos, f.Follower)
	}
	return util.FollowListResponse{
		Response: util.Response{
			StatusCode: 0,
			StatusMsg:  "查询粉丝列表成功",
		},
		UserList: userinfos,
	}

}
