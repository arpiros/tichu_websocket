package main

import (
	"math/rand"
	"time"
	log "github.com/Sirupsen/logrus"
	"tichu/tichu_websocket/system"
	"tichu/tichu_websocket/local"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	log.Infof("Tichu Server Start")

	system.InitEngine()

	StartService()
}

func StartService() {
	local.StartRouter()
}