package main

import (
	"github.com/Sirupsen/logrus"
	"math/rand"
	"tichu/tichu_websocket/local"
	"tichu/tichu_websocket/system"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	logrus.Infof("Tichu Server Start")

	system.InitEngine()

	StartService()
}

func StartService() {
	local.StartRouter()
}
