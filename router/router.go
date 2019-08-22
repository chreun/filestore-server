package router

import (
	"filestore-server/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.StaticFS("/", http.Dir("./static"))

	r.POST("/login", handler.DoLoginHandler)
	r.POST("/register", handler.DoRegisterHandler)
	return r
}