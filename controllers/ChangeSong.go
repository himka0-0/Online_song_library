package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
	"online_song/config"
	"online_song/logger"
	"online_song/models"
)

func ChangeSongHandler(c *gin.Context) {
	var input models.Songs
	if err := c.ShouldBindJSON(&input); err != nil {
		logger.Logger.Warn("Не правильно введены данные", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Не правильно введены данные"})
		return
	}
	var existingSong models.Songs
	if err := config.DB.First(&existingSong, "ID = ?", input.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Logger.Warn("Запись с указанным ID не найдена", zap.Error(err))
			c.JSON(http.StatusNotFound, gin.H{"msg": "Запись с указанным ID не найдена"})
		} else {
			logger.Logger.Warn("Ошибка при поиске записи", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "Ошибка при поиске записи"})
		}
		return
	}
	err := config.DB.Model(&models.Songs{}).Where("ID=?", input.ID).Updates(models.Songs{
		Muzgroup:    input.Muzgroup,
		Song:        input.Song,
		ReleaseDate: input.ReleaseDate,
		Text:        input.Text,
		Link:        input.Link,
	}).Error
	if err != nil {
		logger.Logger.Error("Ошибка обновлении данных песни", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Данные не сохранены"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Данные успешно обновлены"})
}
