package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"online_song/config"
	"online_song/logger"
	"online_song/models"
	"online_song/utils"
	"strings"
)

func CreateSongHandler(c *gin.Context) {
	var input models.CreateSong
	if err := c.ShouldBindJSON(&input); err != nil {
		logger.Logger.Warn("не парсится пришедшие данные для создания лекцции", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Не правильно введены данные"})
		return
	}

	//отправка в функцию связи с внешним api
	songDetails, err := utils.GetSongInfo(input.Group, input.Song)
	if err != nil {
		logger.Logger.Warn("Ошибка вызова внешнего API:", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Ошибка при получении дополнительной информации"})
		return
	}

	input.Group = strings.ToLower(input.Group)
	input.Song = strings.ToLower(input.Song)

	//создаем новый песню по пришедшим данным
	newSong := models.Songs{
		Muzgroup:    input.Group,
		Song:        input.Song,
		ReleaseDate: songDetails.ReleaseDate,
		Text:        songDetails.Text,
		Link:        songDetails.Link,
	}

	err = config.DB.Create(&newSong).Error
	if err != nil {
		logger.Logger.Error("Не сохраняются данные в бд, при создании песни", zap.Error(err))
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Песня сохранена",
	})
}
