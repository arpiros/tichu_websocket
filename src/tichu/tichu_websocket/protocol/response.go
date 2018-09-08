package protocol

import "tichu/tichu_websocket/models"

type CreateRoomResp struct {
	RoomCode string
}

type JoinRoomResp struct {
	UserCount int
}

type MoveTurnResp struct {
	Message string
}

type StartGameResp struct {
	Team *models.Team
	Player *models.Player
}