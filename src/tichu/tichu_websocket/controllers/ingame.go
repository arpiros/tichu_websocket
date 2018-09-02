package controllers

import (
	"sync"
	"tichu/tichu_websocket/models"
	"tichu/tichu_websocket/protocol"
	"github.com/gorilla/websocket"
	"log"
)

type room struct {
	roomNumber int
	rooms map[string]models.Hub
	mutex *sync.RWMutex
}

func (t *room) rLock() {
	t.mutex.RLock()
}

func (t *room) rUnlock() {
	t.mutex.RUnlock()
}

func (t *room) write_acquire() {
	t.mutex.Lock()
}

func (t *room) write_release() {
	t.mutex.Unlock()
}

var defRoom *room = &room{
	mutex: &sync.RWMutex{},
}

func GetAllRoomHub() map[string]models.Hub {
	defRoom.rLock()
	defer defRoom.rUnlock()

	return defRoom.rooms
}

func CreateHub(roomCode string) map[string]models.Hub {
	defRoom.write_acquire()
	defer defRoom.write_release()

	_, ok := defRoom.rooms[roomCode]
	if !ok {
		defRoom.rooms = make(map[string]models.Hub)
	}

	defRoom.rooms[roomCode] = models.NewHub()

	return defRoom.rooms
}

func GetRoomNumber() int {
	defRoom.rLock()
	defer defRoom.rUnlock()

	ret := defRoom.roomNumber
	defRoom.roomNumber += 1

	return ret
}

func CreateRoom(data protocol.CreateRoom, c *websocket.Conn) {
	roomList := GetAllRoomHub()
	roomNumber := GetRoomNumber()
	if len(roomList) == 0 {
		roomList = CreateHub(roomNumber)
	}
	room, ok := roomList[roomNumber]
	if !ok {
		//roomList =
	}

	room.Run()
}

func Join(data protocol.Join, c *websocket.Conn) {
	roomList := GetAllRoomHub()
	_, ok := roomList[data.RoomNumber]
	if !ok {

	}
}

func Response(c *websocket.Conn, v interface{}) {
	err := c.WriteJSON(v)
	if err != nil {
		log.Println("WriteJson Error : ", err)
	}
}