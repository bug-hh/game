package main


import (
	"flag"
	"fmt"
	"net/http"
	"runtime"

	//"code.google.com/p/go.net/websocket"
)

var addr = flag.String("addr", "localhost:8888", "http service address")

// 游戏服务器初始化
func init()  {
	
}

func main() {
	// 游戏服务器开发如何利用 cpu 多核
	fmt.Println("本机几核：", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU() * 2)
	http.HandleFunc("/", wwwGolangLtd)
	http.ListenAndServe(*addr, nil)

}