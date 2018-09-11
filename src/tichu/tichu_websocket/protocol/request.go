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

type CallLargeTichuReq struct {
	RequestBase

	IsCall bool
}

type ChangeCardReq struct {
	RequestBase

	Change map[int]int
}

type CallTichuReq struct {
	RequestBase
}

type UseBoomReq struct {
	RequestBase

	Cards []int
}

type MoveTurnReq struct {
	RequestBase

	Message string `json:"msg"`
}
