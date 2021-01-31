package main


import (
	"fmt"
	"github.com/bug-hh/websocket"
	"net/http"
	"runtime"

	//"code.google.com/p/go.net/websocket"
)

func main() {
	// 游戏服务器开发如何利用 cpu 多核
	fmt.Println("本机几核：", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
	http.Handle("/GolangLtd", websocket.Handler(wwwGolangLtd))

	if err := http.ListenAndServe(":8888", nil); err != nil {
		fmt.Println("网络错误", err)
		return
	}

}