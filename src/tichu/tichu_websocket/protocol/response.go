package protocol

import "tichu/tichu_websocket/models"

type CreateRoomResp struct {
	RoomCode string
}

type JoinRoomResp struct {
	User models.User
}

type MoveTurnResp struct {
	Message string
}