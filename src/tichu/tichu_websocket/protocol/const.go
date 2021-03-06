package protocol

// Request
const (
	ReqCreateRoom = iota
	ReqJoinRoom
	ReqCallLargeTichu
	ReqChangeCard
	ReqCallTichu
	ReqSubmitCard
	ReqTurnPass
	ReqBoom
	ReqMoveTurn
)

// Response
const (
	RespCreateRoom = iota
	RespJoinRoom
	RespRoomInit
	RespCallLargeTichu
	RespDistributeAllCard
	RespStartGame
	RespCallTichu
	RespSubmitCard
	RespTurnPass
	RespNextGame
)
