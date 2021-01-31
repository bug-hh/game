package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}
func wwwGolangLtd(w http.ResponseWriter, r *http.Request) {
	_, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("网络错误", err)
		return
	}

	fmt.Fprintln(w, "Golang 社区欢迎您！！！")
}
