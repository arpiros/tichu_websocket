package models

import (
	"errors"
	"github.com/Sirupsen/logrus"
	"github.com/gorilla/websocket"
)

var UserList = make(map[*websocket.Conn]*User)

const (
	UserState_None = iota
	UserState_InRoom
	UserState_Playing
)

type User struct {
	//Clients map[*websocket.Conn]bool
	RoomCode string
	State    int
}

func AddUser(ws *websocket.Conn) {
	if _, ok := UserList[ws]; ok {
		return
	}

	UserList[ws] = &User{
		RoomCode: "",
		State:    UserState_None,
	}

	logrus.Infof("Add User, Total User : %d", len(UserList))
}

func DelUser(ws *websocket.Conn) {
	delete(UserList, ws)
	ws.Close()
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
