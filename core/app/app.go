package app

import (
	"multiverse/core/controllers/health"
	"multiverse/core/middlewares"

	"github.com/gin-gonic/gin"
)

type App struct {
	route *gin.Engine
}

func NewApp() *App {
	r := gin.Default()
	routing(r)
	return &App{
		route: r,
	}
}

func (a *App) Start(addr string) error {
	return a.route.Run(addr)
}

func routing(r *gin.Engine) {
	r.Use(middlewares.CORSMiddleware())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	healthCheckController := health.NewHealthCheckController()
	//unprotected routes
	r.GET("/health", healthCheckController.GetStatus())

	//Protected routes
	r.Use(middlewares.JwtAuthorizationMiddleware())
}
