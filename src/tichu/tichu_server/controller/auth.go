package controller

import "github.com/gin-gonic/gin"

func Login(c *gin.Context) {
	tick := c.PostForm("tick")

	c.String(200, tick)
}
