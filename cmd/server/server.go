package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

// Router is an alias for gin.Engine
type Router = gin.Engine

func NewRouter() *Router {
	router := gin.New()
	router.Use(gin.Recovery())
	router.SetTrustedProxies(nil)
	return router
}

func main() {

	app := fx.New(
		fx.Provide(NewRouter),
		// Regiter http routes here
		fx.Provide(NewUserRouter),
		fx.Provide(NewAdminRouter),
		fx.Provide(ProvideRoutes),
		// Register middlewares here
		fx.Provide(NewCorsMiddleware),
		fx.Provide(ProvideMiddlewares),
		fx.Invoke(registerHooks),
	)

	ctx := context.Background()

	app.Start(ctx)

	defer app.Stop(ctx)

}

func registerHooks(
	lf fx.Lifecycle,
	router *gin.Engine,
	routes Routes,
	middlewares Middlewares,
) {
	server := http.Server{
		Addr:    ":4242",
		Handler: router,
	}

	lf.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			router.Use(gin.Logger())
			router.Use(gin.Recovery())
			// Register middlewares
			middlewares.Registers()
			// Register routes
			routes.Registers()
			// server.Handler = router
			fmt.Println("Starting server")
			go server.ListenAndServe()

			return nil

		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Stopping server")
			server.Shutdown(ctx)
			return nil
		},
	})

}

type Params struct {
	fx.In
	UserRouter  *UserRouter
	AdminRouter *AdminRouter
}

// Register routes here
func ProvideRoutes(
	params Params,
) Routes {
	return Routes{
		params.UserRouter,
		params.AdminRouter,
	}
}

// routes
type Routes []Route

func (r Routes) Registers() {
	for _, route := range r {
		route.Register()
	}
}

// middlewares

type MiddlewareParams struct {
	fx.In
	CorsMiddleware *CorsMiddleware
}

type Middlewares []Middleware

func ProvideMiddlewares(params MiddlewareParams) Middlewares {
	return Middlewares{
		params.CorsMiddleware,
	}
}

func (m Middlewares) Registers() {
	for _, middleware := range m {
		middleware.Register()
	}
}

type Middleware interface {
	Register()
}

type Route interface {
	Register()
}

type UserRouter struct {
	router *Router
}

func NewUserRouter(router *gin.Engine) *UserRouter {
	return &UserRouter{
		router: router,
	}
}

func (u *UserRouter) Register() {
	u.router.GET("/user", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

}

type AdminRouter struct {
	router *Router
}

func NewAdminRouter(router *gin.Engine) *AdminRouter {
	return &AdminRouter{
		router: router,
	}
}

func (a *AdminRouter) Register() {
	a.router.GET("/admin", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

type CorsMiddleware struct {
	gin *gin.Engine
}

func NewCorsMiddleware(gin *gin.Engine) *CorsMiddleware {
	return &CorsMiddleware{
		gin: gin,
	}
}

func (c *CorsMiddleware) Register() {
	c.gin.Use(func(ctx *gin.Context) {
		fmt.Println("cors middleware")
	})
}
