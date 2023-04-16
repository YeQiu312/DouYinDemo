package routers

import (
	"douyin/controller"
	"douyin/middleware"
	"github.com/gin-gonic/gin"
)

// 路由统一管理
func InitRouter() *gin.Engine {
	r := gin.Default()

	//设置本机访问视频文件路由，使用云服务时记得删除
	r.Static("/public", "./public")
	r.Static("/cover", "./cover")

	douyinGroup := r.Group("/douyin")
	{

		//用户路由分组
		userGroup := douyinGroup.Group("user")
		{
			userGroup.POST("/register/", controller.Register)
			userGroup.POST("/login/", controller.Login)
			userGroup.GET("/", controller.FindUserInfo).Use(middleware.AuthMiddleware())
		}

		//视频流路由分组
		feedGroup := douyinGroup.Group("/feed")
		{
			feedGroup.GET("/", controller.Feed)
		}

		//投稿、发布接口路由分组
		publicGroup := douyinGroup.Group("publish").Use(middleware.AuthMiddleware())
		{
			publicGroup.GET("/list/", controller.List)
			publicGroup.POST("/action/", controller.Action)
		}

		//关注、粉丝、好友接口
		RelationGroup := douyinGroup.Group("relation").Use(middleware.AuthMiddleware())
		//RelationGroup := douyinGroup.Group("/relation")
		{
			RelationGroup.POST("/action/", controller.Follow)
			RelationGroup.GET("/follow/list/", controller.FollowList)
			RelationGroup.GET("/follower/list/", controller.FollowerList)
			RelationGroup.GET("/friend/list/", controller.FriendList)
		}

		//聊天接口
		MessageGroup := douyinGroup.Group("message").Use(middleware.AuthMiddleware())
		{
			MessageGroup.POST("/action/", controller.SendMessage)
			MessageGroup.GET("/chat/", controller.ChatMessage)
		}

		////喜欢、赞操作
		//FavoriteGroup := douyinGroup.Group("favorite").Use(middleware.AuthMiddleware())
		//{
		//	FavoriteGroup.POST("/action/")
		//	FavoriteGroup.GET("/list/")
		//}
		//
		////评论操作、评论列表
		//CommmentGroup := douyinGroup.Group("comment")
		//{
		//	CommmentGroup.POST("/action/")
		//	CommmentGroup.GET("/list/")
		//}

	}
	return r
}
