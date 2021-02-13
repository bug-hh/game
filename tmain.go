package main

import (
	"github.com/bug-hh/websocket"
	"github.com/golang/glog"
	"net/http"
	"runtime"

	//"code.google.com/p/go.net/websocket"
)

//func init() {
	//flag.Set("alsologtostderr", "true")   // 日志文件写入的同时，输出到 stderr
	//flag.Set("log_dir", "./log")  // 日志文件保存目录
	//flag.Set("v", "3")  // 日志等级
	//flag.Parse()
//}

func main() {
	// 游戏服务器开发如何利用 cpu 多核
	glog.Info("本机几核：", runtime.NumCPU())
	glog.Flush()
	runtime.GOMAXPROCS(runtime.NumCPU())
	http.Handle("/GolangLtd", websocket.Handler(wwwGolangLtd))

	if err := http.ListenAndServe(":8888", nil); err != nil {
		glog.Info("网络错误", err)
		return
	}



}