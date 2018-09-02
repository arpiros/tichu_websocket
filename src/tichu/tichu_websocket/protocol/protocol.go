package protocol

const (
	CREATE_ROOM = iota
	JOIN
)

type Base struct {
	ProtocolType int
}

type CreateRoom struct {
	Base


}

type Join struct {
	Base

	RoomNumber int
}