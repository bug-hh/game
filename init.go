package main

import (
	"flag"
	"github.com/golang/glog"
)

// 1 网络层处理  -- Main -- > 消息处理 玩家绑定 uid - go
// 2 定时器
// 3 逻辑处理

// 游戏服务器初始化
func init()  {
	flag.Set("alsologtostderr", "true")   // 日志文件写入的同时，输出到 stderr
	flag.Set("log_dir", "./log")  // 日志文件保存目录
	flag.Set("v", "3")  // 日志等级
	flag.Parse()
	Go_func()
}

func Go_func() {
	glog.Info("Golang语言社区")
	//fmt.Println("Go_func")
	return
}
