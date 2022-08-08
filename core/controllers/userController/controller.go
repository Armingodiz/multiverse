package userController

import (
	"multiverse/core/models"
	"multiverse/core/services/userService"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService userService.UserService
}

func (u *UserController) Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := u.UserService.CreateUser(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "created"})
	}
}
