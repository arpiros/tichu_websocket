package protocol

type RequestBase struct {
	RequestType int `json:"req_t"`
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

	IsCall int
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

type SubmitCardReq struct {
	RequestBase

	Cards []int
}

type PassTurnReq struct {
	RequestBase
}

type MoveTurnReq struct {
	RequestBase

	Message string `json:"msg"`
}
