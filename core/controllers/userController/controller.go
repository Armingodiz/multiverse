package userController

import (
	"multiverse/core/models"
	"multiverse/core/services/brokerService"
	"multiverse/core/services/calculatorService"
	"multiverse/core/services/userService"
	"multiverse/core/services/welcomerService"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService      userService.UserService
	WelcomerService  welcomerService.WelcomerService
	CalculatorServie calculatorService.CalculatorService
	BrokerService    brokerService.BrokerService
}

func (u *UserController) Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user.RegistrationDate = time.Now().String()
		err := u.UserService.CreateUser(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		message, err := u.WelcomerService.GetWelcomeMessage(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		err = u.BrokerService.Publish(models.Task{
			Target: user.Email,
			Text:   message.GetResult(),
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "created"})
	}
}

func (u *UserController) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Param("email")
		user, err := u.UserService.GetUser(email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"user": user})
	}
}

func (u *UserController) DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Param("email")
		err := u.UserService.DeleteUser(email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "deleted"})
	}
}

func (u *UserController) Calculate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var calculation models.Calculation
		if err := c.ShouldBindJSON(&calculation); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, err := u.CalculatorServie.Calculate(calculation)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"result": res})
	}
}
