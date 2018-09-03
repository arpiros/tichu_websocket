package models

import (
	"github.com/gorilla/websocket"
	"tichu/tichu_websocket/util"
	"github.com/Sirupsen/logrus"
)

const RoomCodeLength = 4

//TODO Mutex 처리
var RoomList = make(map[string]Room)

type Room struct {
	RoomCode  string
	Clients   map[*websocket.Conn]bool
	Broadcast chan InGameBroadCast
}

//func (t *Room) rLock() {
//	t.mutex.RLock()
//}
//
//func (t *Room) rUnlock() {
//	t.mutex.RUnlock()
//}
//
//func (t *Room) write_acquire() {
//	t.mutex.Lock()
//}
//
//func (t *Room) write_release() {
//	t.mutex.Unlock()
//}
//
////var defRoom *Room = &Room{
////	mutex: &sync.RWMutex{},
////}

type InGameBroadCast struct {
}

func CreateRoom(ws *websocket.Conn) *Room {
	// 5번 이상 돌지 않도록
	for i := 0; i < 5; i++ {
		roomCode := util.GenerateRandomString(RoomCodeLength)
		if _, ok := RoomList[roomCode]; !ok {
			room := Room{
				RoomCode:  roomCode,
				Clients:   make(map[*websocket.Conn]bool),
				Broadcast: make(chan InGameBroadCast),
			}

			room.Clients[ws] = true

			RoomList[roomCode] = room
			return &room
		}
	}
	return nil
}

func JoinRoom(ws *websocket.Conn, roomCode string) {
	//TODO user State Check

	if _, ok := RoomList[roomCode]; !ok {
		//TODO error
		return
	}

	room := RoomList[roomCode]

	room.Clients[ws] = true

	logrus.Infof("Join Room")
}

func LeaveRoom(ws *websocket.Conn, roomCode string) {
	if _, ok := RoomList[roomCode]; !ok {
		//TODO error
		return
	}

	room := RoomList[roomCode]

	room.Clients[ws] = false

	for _, value := range room.Clients {
		if value {
			return
		}
	}

	delete(RoomList, roomCode)
}