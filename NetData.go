package main

import (
	"encoding/json"
	"fmt"
	"game/Protocol"
	"game/Protocol/Proto2"
	"github.com/bug-hh/websocket"
	"github.com/golang/glog"
)


// 网络数据结构的保存
// 1 websocket 的网络链接
// 2 StrMD5 房间加密信息

type NetDataConn struct {
	Connection *websocket.Conn
	StrMD5 string
}

// 结构体数据类型
type Requestbody struct {
	req string
}

// json 转化为 map数据的处理
func (r *Requestbody) Json2map() (s map[string]interface{}, err error) {
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(r.req), &result); err != nil {
		glog.Info("Json2map:", err.Error())
		return nil, err
	}
	return result, nil
}
// NetDataConn 结构体的方法 - 接受者是指针类型
func (this *NetDataConn) PullFromClient()  {
	// 网络层处理数据
	// 1 针对服务器而言 一直等待消息
	fmt.Println("PullFromClient")
	for {
		fmt.Println("进入循环")
		var content string
		if err := websocket.Message.Receive(this.Connection, &content); err != nil {
			break
		}
		if len(content) == 0 {
			break
		}
		// 正常处理消息
		// go 并发编程使用
		go this.SyncMessageFunc(content)
	}
	fmt.Println("结束循环")
	return
}

func (this *NetDataConn) SyncMessageFunc(content string) {
	// 1 字符串 转 其他格式 必须高效 （大量并发情况下，依然不影响性能，游戏服务器 计算密集型）
	fmt.Println(content)
	// 2 已经通过第一步转化成我们所有的格式了，实现格式的处理函数：主协议、子协议、struct
	var r Requestbody
	r.req = content

	if ProtocolData, err := r.Json2map(); err == nil {
		// 处理我们的函数
		this.HandleCltProtocol(ProtocolData["proto"], ProtocolData["proto2"], ProtocolData)
	} else {
		fmt.Println("解析失败: ", err.Error())
	}
}

func (this *NetDataConn) HandleCltProtocol(protocol interface{}, protocol2 interface{}, ProtocolData map[string]interface{}) {
	// 分发处理   ---  首先判断主协议存在，再判断子协议是否存在
	switch protocol {
	case Protocol.GameData_Proto:
		{
			// 子协议处理

		}
	case Protocol.GameDataDB_Proto:
		{}
	default:
		panic("主协议：不存在!!!")
	}
	return


}

// 子协议处理
func (this *NetDataConn) HandleCltProtocol2(protocol2 interface{}, ProtocolData map[string]interface{}) {
	switch protocol2 {
	case Proto2.C2S_PlayerLoginProto2:
		{
			// 功能函数处理 -- 用户登录协议
		}
	default:
		panic("子协议：不存在!!!")
	}
	return

}