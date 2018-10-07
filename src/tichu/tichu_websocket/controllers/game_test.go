package controllers

import (
	"encoding/json"
	"math/rand"
	"sort"
	"testing"
	"tichu/tichu_websocket/protocol"
)

func TestShuffleArray(t *testing.T) {
	list := rand.Perm(25)
	for key, value := range list {
		println(key, value)
	}
}

func TestSortArray(t *testing.T) {
	var numbers []int
	for i := 0; i < 5; i++ {
		numbers = append(numbers, rand.Intn(13)+1)
	}

	sort.Ints(numbers)

	for _, value := range numbers {
		println(value)
	}
}

func TestMarshalJson(t *testing.T) {
	req := protocol.ChangeCardReq{
		RequestBase: protocol.RequestBase{
			RequestType: protocol.ReqChangeCard,
		},
		Change: map[int]int{
			1: 0,
			2: 1,
			3: 2,
		},
	}

	bytes, _ := json.Marshal(req)
	println(string(bytes))
}

func TestTurnChange(t *testing.T) {
	value := 0
	for i := 0; i < 10000; i++ {
		value = (value + 1) % 4

		println(value)
	}
}
