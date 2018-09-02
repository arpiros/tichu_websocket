package main

import (
	"time"
	"math/rand"
	"tichu/tichu_server/local"
	"tichu/tichu_server/system"

	log "github.com/Sirupsen/logrus"
)

func main () {
	rand.Seed(time.Now().UnixNano())
	log.Infof("Tichu Server Start")

	system.InitEngine()

	startService()
}

func startService() {
	system.StartRouter(local.SetupPuzzleRouters)
}
