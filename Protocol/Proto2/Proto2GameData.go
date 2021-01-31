package Proto2

// GameData_Proto 的子协议
const (
	INIT_PROTO2 = iota
	C2S_PlayerLoginProto2 // 为 1，用户登录协议
	S2C_PlayerloginProto2 // 为 2，用户登录协议

	C2S_ChooseRoomProto2 // 为 3，选择房间
	S2C_ChooseRoomProto2 // 为 4，选择房间
)


type PlayerSt struct {
	UID int
	PlayerName string
	OpenID string
}

// 功能结构

type Head_Proto struct {
	Proto int // 主协议
	Proto2 int // 子协议
}
type C2S_PlayerLogin struct {
	Head_Proto
	Code string // 微信授权 CODE
}

type S2C_PlayerLogin struct {
	Head_Proto
	PlayerData *PlayerSt // 玩家的结构
}