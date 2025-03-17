package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"online_song/config"
	"online_song/models"
)

func DeleteSongHandler(c *gin.Context) {
	var input models.DeleteSong
	err := c.ShouldBindJSON(&input)
	if err != nil {
		log.Println("Ошибка парсинга ID при удалении", err)
	}
	err = config.DB.Where("ID=?", input.ID).Delete(models.Songs{}).Error
	if err != nil {
		log.Println("Ошибка удаления пользователя из бд", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении записи"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Запись успешно удалена"})
}
