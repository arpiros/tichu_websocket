package models

import (
	"github.com/gorilla/websocket"
	"github.com/Sirupsen/logrus"
)

var AllUserMap map[*websocket.Conn]User = make(map[*websocket.Conn]User)

const (
	UserState_None = iota
	UserState_Wait
	UserState_Playing
)

type User struct {
	//Clients map[*websocket.Conn]bool
	RoomCode string
	State int
}

func AddUser(ws *websocket.Conn) {
	if _, ok := AllUserMap[ws]; ok {
		return
	}

	AllUserMap[ws] = User{
		RoomCode: "",
		State: UserState_None,
	}

	logrus.Infof("Add User, Total User : %d", len(AllUserMap))
}

func DelUser(ws *websocket.Conn) {
	delete(AllUserMap, ws)
}