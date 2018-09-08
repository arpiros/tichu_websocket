package protocol

import "tichu/tichu_websocket/models"

type CreateRoomResp struct {
	RoomCode string
}

type JoinRoomResp struct {
	UserCount int
}

type RoomInitResp struct {
	Team   *models.Team
	Player *models.Player
}

type CallLargeTichuResp struct {
	CallTichu map[int]int
}

type DistributeAllCardResp struct {
	Player *models.Player
}

type StartGameResp struct {
	Player              *models.Player
	CurrentActivePlayer int
}

type MoveTurnResp struct {
	Message string
}
