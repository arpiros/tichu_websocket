package protocol

type Base struct {
	ProtocolType int `json:"pt"`
}

type CreateRoomReq struct {
	Base
}

type JoinRoomReq struct {
	Base

	RoomCode string
}