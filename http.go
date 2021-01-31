package main

import (
	"fmt"
	"github.com/bug-hh/websocket"
)


func wwwGolangLtd(ws *websocket.Conn) {
	fmt.Println("Golang 社区欢迎您！！！")
	data := ws.Request().URL.Query().Get("data")
	fmt.Println("data: ", data)
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
