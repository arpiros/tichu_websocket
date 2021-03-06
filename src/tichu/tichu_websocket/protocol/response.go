package protocol

import "tichu/tichu_websocket/models"

type BaseResp struct {
	ResponseType int `json:"resp_t"`
}

func NewBaseResp(respType int) BaseResp {
	return BaseResp{
		ResponseType: respType,
	}
}

type CreateRoomResp struct {
	BaseResp

	RoomCode string
}

type JoinRoomResp struct {
	BaseResp

	UserCount int
}

type RoomInitResp struct {
	BaseResp

	Team   *models.Team
	Player *models.Player
}

type CallLargeTichuResp struct {
	BaseResp

	CallTichu map[int]int
}

type DistributeAllCardResp struct {
	BaseResp

	Player    *models.Player
	CallTichu map[int]int
}

type StartGameResp struct {
	BaseResp

	GainCard            models.CardList
	CurrentActivePlayer int
}

type CallTichuResp struct {
	BaseResp

	CallTichu map[int]int
}

type SubmitCardResp struct {
	BaseResp

	Player              *models.Player
	CurrentActivePlayer int
}

type TurnPassResp struct {
	BaseResp

	Player              *models.Player
	CurrentActivePlayer int
}

type NextGameResp struct {
	BaseResp

	Player *models.Player
}

type MoveTurnResp struct {
	BaseResp

	Message string
}
