package main

import (
	"github.com/bug-hh/websocket"
	"github.com/golang/glog"
)


func wwwGolangLtd(ws *websocket.Conn) {
	glog.Info("Golang 社区欢迎您！！！")
	data := ws.Request().URL.Query().Get("data")
	glog.Info("data: ", data)
	// 消息处理
	// 先处理 主协议
	// 再处理 子协议
	NetDataConnTmp := &NetDataConn{
		Connection: ws,
		StrMD5: "",
	}

	// 从客户端取消息
	NetDataConnTmp.PullFromClient()

}
