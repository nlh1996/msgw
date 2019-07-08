package api

import (
	"go-gateway/controller/http"

	"github.com/gin-gonic/gin"
)

func init() {
	router := gin.Default()
	router.GET("/", http.GetName)
	router.Run(":3100")
}
