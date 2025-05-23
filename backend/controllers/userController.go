package controllers

import (
	"gin_app/config"
	"gin_app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	var users []models.User
	config.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   users,
	})
}
