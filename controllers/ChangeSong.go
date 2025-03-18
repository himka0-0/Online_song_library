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

// ChangeSongHandler обновляет данные существующей песни
//
// @Summary      Обновление песни
// @Description  Обновляет данные песни в базе данных по переданному JSON. В JSON должны быть указаны все нужные поля, включая ID для идентификации записи.
// @Tags         Songs
// @Accept       json
// @Produce      json
// @Param        song  body      models.Songs true "Объект песни с новыми данными"
// @Success      200   {object}  map[string]string "Данные успешно обновлены"
// @Failure      400   {object}  map[string]string "Неправильно введены данные или данные не сохранены"
// @Failure      404   {object}  map[string]string "Запись с указанным ID не найдена"
// @Failure      500   {object}  map[string]string "Ошибка при поиске записи"
// @Router       / [put]
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
