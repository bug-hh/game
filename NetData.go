package main

import (
	"encoding/json"
	"fmt"
	"game/Protocol"
	"game/Protocol/Proto2"
	"github.com/bug-hh/websocket"
	"github.com/golang/glog"
	"reflect"
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
		glog.Error("Json2map:", err.Error())
		return nil, err
	}
	return result, nil
}
// NetDataConn 结构体的方法 - 接受者是指针类型
func (this *NetDataConn) PullFromClient()  {
	// 网络层处理数据
	// 1 针对服务器而言 一直等待消息
	glog.Info("PullFromClient")
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
	return
}

func (this *NetDataConn) SyncMessageFunc(content string) {
	// 1 字符串 转 其他格式 必须高效 （大量并发情况下，依然不影响性能，游戏服务器 计算密集型）
	// 2 已经通过第一步转化成我们所有的格式了，实现格式的处理函数：主协议、子协议、struct
	var r Requestbody
	r.req = content

	if ProtocolData, err := r.Json2map(); err == nil {
		// 处理我们的函数
		this.HandleCltProtocol(ProtocolData["Protocol"], ProtocolData["Protocol2"], ProtocolData)
	} else {
		glog.Error("解析失败: ", err.Error())
	}
}

func (this *NetDataConn) HandleCltProtocol(protocol interface{}, protocol2 interface{}, ProtocolData map[string]interface{}) {
	// 分发处理   ---  首先判断主协议存在，再判断子协议是否存在
	//fmt.Println(protocol)
	//fmt.Println(Protocol.GameData_Proto)
	//fmt.Println(typeof(protocol))
	//fmt.Println(typeof(Protocol.GameData_Proto))
	switch protocol {
	case float64(Protocol.GameData_Proto):
		{
			// 子协议处理
			this.HandleCltProtocol2(protocol, ProtocolData)
		}
	case float64(Protocol.GameDataDB_Proto):
		{}
	default:
		panic("主协议：不存在!!!")
	}
	return


}

func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}

// 子协议处理
func (this *NetDataConn) HandleCltProtocol2(protocol2 interface{}, ProtocolData map[string]interface{}) {
	switch protocol2 {
	case float64(Proto2.C2S_PlayerLoginProto2):
		{
			// 功能函数处理 -- 用户登录协议
			this.PlayerLogin(ProtocolData)
		}
	default:
		panic("子协议：不存在!!!")
	}
	return

}

// 用户登录协议
func (this *NetDataConn) PlayerLogin(ProtocolData map[string]interface{}) {
	// 服务器的逻辑处理
	// 获取用户发送来的信息
	if ProtocolData["StrLoginName"] == nil ||
		ProtocolData["StrLoginPW"] == nil ||
		ProtocolData["StrLoginEmail"] == nil {
		panic("登录功能数据错误")
		return
	}

	StrLoginName := ProtocolData["StrLoginName"].(string)
	StrLoginPW := ProtocolData["StrLoginPW"].(string)
	StrLoginEmail := ProtocolData["StrLoginEmail"].(string)

	glog.Info(StrLoginName, StrLoginPW, StrLoginEmail)
	// 数据库保存 -- 发给 DBServer
	// 服务器返给客户端
	head_data := Proto2.Head_Proto {
		Protocol1: Protocol.GameData_Proto,
		Protocol2: Proto2.S2C_PlayerloginProto2,
	}
	data := &Proto2.S2C_PlayerLogin {
		Head_Proto: head_data,
		PlayerData: nil,
	}
	// 发送数据给客户端
	this.PlayerSendMessage(data)

	return
}

// 发送给客户端的数据信息函数
func (this *NetDataConn) PlayerSendMessage(senddata interface{}) {
	// 消息序列化，interface —-> json
	b, errjson := json.Marshal(senddata)
	if errjson != nil {
		glog.Error(errjson.Error())
		return
	}

	//数据转换 json 的 byte 数组 ---> string
	data := "data:" + string(b[0:len(b)])
	fmt.Println(data)
	err := websocket.JSON.Send(this.Connection, senddata)
	if err != nil {
		glog.Error(err.Error())
		return
	}
	return
}

