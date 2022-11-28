package main

import (
	"ChiesePalOnlineServer/api"
	"github.com/aceld/zinx/znet"
)

func main() {
	//创建服务器句柄
	s := znet.NewServer()

	// 连接创建和销毁HOOK钩子函数
	s.SetOnConnStart(OnConnectionAdd)

	//注册一些路由服务
	s.AddRouter(2, &api.WorldChatApi{})

	//启动服务
	s.Serve()
}
