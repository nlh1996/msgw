package api

import (
	"go-gateway/controller/ws"

	"github.com/gin-gonic/gin"
)

func init() {
	router := gin.Default()
	router.GET("/", ws.Ping)
	router.Run(":3000")
}
