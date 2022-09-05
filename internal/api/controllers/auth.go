package controllers

import (
	"net/http"
	"sum/internal/api"
	"sum/internal/models"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Input struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Auth(tokenTTL time.Duration, tokenSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input Input

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid json"})
			return
		}

		user := models.User{}
		user.Username = input.Username
		user.Password = input.Password

		err := user.CheckCredentials()
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "incorrect username or password"})
			return
		}

		token, err := api.GenerateToken(user.Username, tokenSecret, tokenTTL)
		if err != nil {
			log.Errorf("internal server error %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "internal problem with the request"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
