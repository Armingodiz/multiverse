package app

import (
	"context"
	"multiverse/core/config"
	"multiverse/core/controllers/health"
	"multiverse/core/controllers/userController"
	"multiverse/core/middlewares"
	"multiverse/core/services/brokerService"
	"multiverse/core/services/calculatorService"
	"multiverse/core/services/userService"
	"multiverse/core/services/welcomerService"
	"multiverse/core/store"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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
	Store := store.NewMongoStore(getMongoDbCollection())
	UserService := userService.NewUserService(Store)
	WelcomerService := welcomerService.NewWelcomerService()
	CaclulatorService := calculatorService.NewCalculatorService()
	BrokerService := brokerService.NewBrokerService()
	UserController := userController.UserController{UserService: UserService, WelcomerService: WelcomerService, CalculatorServie: CaclulatorService, BrokerService: BrokerService}
	healthCheckController := health.NewHealthCheckController()
	//unprotected routes
	r.GET("/health", healthCheckController.GetStatus())
	r.POST("/user/signup", UserController.Signup())
	r.POST("/user/login", UserController.Login())
	r.GET("/user/:email", UserController.GetUser())
	r.DELETE("/user/:email", UserController.DeleteUser())
	r.POST("/calculator", UserController.Calculate())

	//Protected routes
	r.Use(middlewares.JwtAuthorizationMiddleware())
}

func getMongoDbCollection() *mongo.Collection {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.Configs.Database.Url))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	return client.Database(config.Configs.Database.DbName).Collection(config.Configs.Database.CollectionName)
}
