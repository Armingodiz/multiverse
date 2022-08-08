package app

import (
	"multiverse/core/controllers/health"
	"multiverse/core/controllers/userController"
	"multiverse/core/middlewares"
	"multiverse/core/services/userService"
	"multiverse/core/store"

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
	Store := store.NewMongoStore()
	UserService := userService.NewUserService(Store)
	UserController := userController.UserController{UserService: UserService}
	healthCheckController := health.NewHealthCheckController()
	//unprotected routes
	r.GET("/health", healthCheckController.GetStatus())
	r.POST("/user/signup", UserController.Signup())
	r.GET("/user/:email", UserController.GetUser())
	r.DELETE("/user/:email", UserController.DeleteUser())

	//Protected routes
	r.Use(middlewares.JwtAuthorizationMiddleware())
}
