package protocol

type RequestBase struct {
	ProtocolType int `json:"pt"`
}

type CreateRoomReq struct {
	RequestBase
}

type JoinRoomReq struct {
	RequestBase

	RoomCode string
}

type MoveTurnReq struct {
	RequestBase

	Message string `json:"msg"`
}