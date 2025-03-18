package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"online_song/config"
	"online_song/logger"
	"online_song/models"
)

func DeleteSongHandler(c *gin.Context) {
	var input models.DeleteSong
	err := c.ShouldBindJSON(&input)
	if err != nil {
		logger.Logger.Warn("Ошибка парсинга ID при удалении", zap.Error(err))
	}
	err = config.DB.Where("ID=?", input.ID).Delete(models.Songs{}).Error
	if err != nil {
		logger.Logger.Error("Ошибка удаления пользователя из бд", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении записи"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Запись успешно удалена"})
}
