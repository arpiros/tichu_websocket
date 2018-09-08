package models

import (
	"github.com/gorilla/websocket"
	"tichu/tichu_websocket/util"
	"errors"
	"github.com/Sirupsen/logrus"
)

const RoomCodeLength = 4
const RoomMemberLimit = 4
const TeamCount = 2

const (
	CallTichuNone  = iota
	CallTichuSmail
	CallTichuLarge
)

//TODO Mutex 처리
var RoomList = make(map[string]*Room)

type Room struct {
	RoomCode string

	Clients  map[*websocket.Conn]*Player `json:"-"`
	Players  []*Player
	Teams    []*Team
	CardDeck []*Card                     `json:"-"`

	CallTichu map[int]int

	CurrentActivePlayer int
}

type Player struct {
	Index      int
	TeamNumber int
	CardList   []*Card
	IsConnect  bool
	IsMyTurn   bool

	GainCard []*Card `json:"-"`
}

type Team struct {
	TeamNumber int
	Player     []*Player `json:"-"`
	TotalScore int
}

func NewTeam(teamNumber int) *Team {
	return &Team{
		TeamNumber: teamNumber,
	}
}

func CreateRoom(ws *websocket.Conn) *Room {
	// 5번 이상 돌지 않도록
	for i := 0; i < 5; i++ {
		roomCode := util.GenerateRandomString(RoomCodeLength)
		if _, ok := RoomList[roomCode]; !ok {
			room := &Room{
				RoomCode:  roomCode,
				Clients:   make(map[*websocket.Conn]*Player),
				Teams:     make([]*Team, TeamCount),
				CallTichu: make(map[int]int),
				CurrentActivePlayer: -1, // 아직 아무도 게임 진행중이지 않음
			}

			for key := range room.Teams {
				room.Teams[key] = NewTeam(key)
			}

			RoomList[roomCode] = room
			return JoinRoom(ws, roomCode)
		}
	}
	return nil
}

func JoinRoom(ws *websocket.Conn, roomCode string) *Room {
	// TODO user State Check

	if _, ok := RoomList[roomCode]; !ok {
		// TODO error
		return nil
	}

	room := RoomList[roomCode]
	if len(room.Players) >= RoomMemberLimit {
		// TODO room member full error
		return nil
	}

	newPlayer := &Player{
		Index:      len(room.Players) + 1,
		TeamNumber: len(room.Players) % TeamCount,
		IsConnect:  true,
	}

	room.Clients[ws] = newPlayer
	room.Players = append(room.Players, newPlayer)
	team := room.Teams[newPlayer.TeamNumber]
	team.TeamNumber = newPlayer.TeamNumber
	team.Player = append(team.Player, newPlayer)

	logrus.Infof("Join Room")

	return room
}

func GetRoom(roomCode string) (*Room, error) {
	room, ok := RoomList[roomCode]
	if !ok {
		return nil, errors.New("Not Found User")
	}

	return room, nil
}

func LeaveRoom(ws *websocket.Conn, roomCode string) {
	if _, ok := RoomList[roomCode]; !ok {
		//TODO error
		return
	}

	room := RoomList[roomCode]

	room.Clients[ws].IsConnect = false

	for _, value := range room.Clients {
		if value.IsConnect {
			return
		}
	}

	delete(RoomList, roomCode)
}
