package route

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func NewServer() *gin.Engine {
	router := gin.New()
	return router
}

var Module = fx.Module("route", fx.Provide(NewServer))
