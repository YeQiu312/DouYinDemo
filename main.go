package main

import (
	"douyin/config"
	"douyin/model"
	"douyin/routers"
	"log"
)

func main() {
	//初始化数据库连接
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("连接数据库失败：%v", err)
	}
	//全局禁用复数表名
	db.SingularTable(true)
	//开启数据自动迁移
	db.AutoMigrate(&model.User{}, &model.Userinfo{}, &model.Relation{}, &model.Message{}, &model.Follow{}, &model.Video{}, &model.VideoUserInfo{})
	//关闭数据库连接
	defer db.Close()

	//初始化路由和中间件
	r := routers.InitRouter()

	//注册路由

	//启动服务
	r.Run()
}
