package models

import (
	"github.com/gorilla/websocket"
	"github.com/Sirupsen/logrus"
	"errors"
)

var UserList = make(map[*websocket.Conn]*User)

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
	if _, ok := UserList[ws]; ok {
		return
	}

	UserList[ws] = &User{
		RoomCode: "",
		State: UserState_None,
	}

	logrus.Infof("Add User, Total User : %d", len(UserList))
}

func DelUser(ws *websocket.Conn) {
	delete(UserList, ws)
}

func GetUser(ws *websocket.Conn) (*User, error) {
	user, ok := UserList[ws]
	if !ok {
		return nil, errors.New("Not Found User")
	}

	return user, nil
}
//
//func JoinRoom(ws *websocket.Conn, roomCode string) error {
//	user, ok := UserList[ws]
//	if !ok {
//		return errors.New("Not Found User")
//	}
//
//	user.RoomCode = roomCode
//	UserList[ws] = user
//
//	return nil
//}