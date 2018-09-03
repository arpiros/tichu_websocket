package protocol

type Base struct {
	ProtocolType int `json:"pt"`
}

type CreateRoomReq struct {
	Base
}

type JoinReq struct {
	Base

	RoomNumber int
}