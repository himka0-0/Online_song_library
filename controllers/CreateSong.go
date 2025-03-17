package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"online_song/config"
	"online_song/models"
	"strings"
)

func CreateSongHandler(c *gin.Context) {
	var input models.CreateSong
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("не парсится пришедшие данные для создания лекцции", err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Не правильно введены данные"})
	}
	input.Group = strings.ToLower(input.Group)
	input.Song = strings.ToLower(input.Song)
	err := config.DB.Create(&models.Songs{Muzgroup: input.Group, Song: input.Song}).Error
	if err != nil {
		log.Println("Не сохраняются данные в бд, при создании песни", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Песня сохранена",
	})
}
