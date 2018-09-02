package local

import (
	"github.com/gin-gonic/gin"
	"tichu/tichu_server/controller"
)

func SetupPuzzleRouters(r *gin.Engine) {
	r.POST("/about", func(c *gin.Context) {
		c.String(200, "Tichu Server")
	})

	r.POST("/auth/login", controller.Login)
	//r.POST("/room/create", controller.CreateRoom)
	//r.POST("/room/join", controller.JoinRoom)
	//r.POST("/room/list", controller.RoomList)
	//r.POST("/room/change_team", controller.ChangeTeam)
	//r.POST("/room/start_game", controller.StartGame)
}