package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"online_song/config"
	"online_song/models"
)

func ChangeSongHandler(c *gin.Context) {
	var input models.Songs
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("Не правильно введены данные", err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Не правильно введены данные"})
		return
	}
	var existingSong models.Songs
	if err := config.DB.First(&existingSong, "ID = ?", input.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Println("Запись с указанным ID не найдена", err)
			c.JSON(http.StatusNotFound, gin.H{"msg": "Запись с указанным ID не найдена"})
		} else {
			log.Println("Ошибка при поиске записи", err)
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
		log.Println("Ошибка обновлении данных песни", err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Данные не сохранены"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Данные успешно обновлены"})
}
