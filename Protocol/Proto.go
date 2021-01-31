package Protocol

const (
	INIT_PROTO = iota
	GameData_Proto  // 为 1，游戏的主协议
	GameDataDB_Proto // 为 2，游戏的 DB 的主协议
)