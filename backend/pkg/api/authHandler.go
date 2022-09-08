package api

import (
	"bhavdeep.me/weight_logger/pkg/db"
	"github.com/gin-gonic/gin"
)

func handleLogin(c *gin.Context) {
	validator := authQuery{}
	if err := c.ShouldBindJSON(&validator); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	tempUser := db.User{
		Username: validator.Username,
		Password: validator.Password,
	}
	if userExists, err := db.FindUser(tempUser); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	} else if userExists.Password == "" {
		c.JSON(400, gin.H{"msg": "User does not exist"})
		return
	} else if userExists.Password == "-1" {
		c.JSON(401, gin.H{"msg": "Unauthorized"})
		return
	} else {
		// TODO Do JWT stuff here
		c.JSON(200, userExists)
	}

}

func handleRegister(c *gin.Context) {
	validator := authQuery{}
	if err := c.ShouldBindJSON(&validator); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	tempUser := db.User{
		Username: validator.Username,
		Password: validator.Password,
	}
	if userExists, err := db.FindUser(tempUser); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	} else if userExists.Password != "" {
		c.JSON(400, gin.H{"msg": "User already exists"})
		return
	}
	if user, err := db.RegisterUser(tempUser); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	} else {
		c.JSON(200, user)
	}
}
