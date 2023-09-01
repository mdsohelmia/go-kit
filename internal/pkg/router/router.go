package router

import (
	"github.com/gin-gonic/gin"
)

// Router is an alias for gin.Engine
type Router = gin.Engine

func NewRouter() *Router {
	router := gin.New()
	router.Use(gin.Recovery())
	router.SetTrustedProxies(nil)
	return router
}
