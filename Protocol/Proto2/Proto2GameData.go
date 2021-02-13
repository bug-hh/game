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
	Protocol1 int // 主协议
	Protocol2 int // 子协议
}
type C2S_PlayerLogin struct {
	Head_Proto
	Itype int // 1 登录，2 注册
	Code string // 微信授权 CODE
	StrLoginName string
	StrLoginPW string
	StrLoginEmail string
}

type S2C_PlayerLogin struct {
	Head_Proto
	PlayerData *PlayerSt // 玩家的结构
}