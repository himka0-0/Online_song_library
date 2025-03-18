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
	"strconv"
	"strings"
)

func VersePage(c *gin.Context) {
	//id песни
	urlID := c.Param("id")

	//проверка наличия id в бд
	var existingSong models.Songs
	if err := config.DB.First(&existingSong, "ID = ?", urlID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Logger.Warn("Запись с указанным ID не найдена", zap.Error(err))
			c.JSON(http.StatusNotFound, gin.H{"msg": "Запись с указанным ID не найдена"})
		} else {
			logger.Logger.Warn("Ошибка при поиске записи", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"msg": "Ошибка при поиске записи"})
		}
		return
	}
	//значени пагинации по-умолчанию
	limit := 2
	offset := 0

	if l := c.Query("limit"); l != "" {
		if lVal, err := strconv.Atoi(l); err == nil {
			limit = lVal
		}
	}
	if o := c.Query("offset"); o != "" {
		if oVal, err := strconv.Atoi(o); err == nil {
			offset = oVal
		}
	}

	var songText string
	err := config.DB.Model(&models.Songs{}).Select("Text").Where("ID=?", urlID).Find(&songText).Error
	if err != nil {
		logger.Logger.Error("Не получилось вытащить текст песни", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Текста песни пока нет"})
		return
	}
	lines := strings.Split(songText, "\n")
	var nonEmptyLines []string
	for _, el := range lines {
		if el != "" {
			nonEmptyLines = append(nonEmptyLines, el)
		}
	}
	//количество страниц
	var pages int
	if len(nonEmptyLines)%4 != 0 {
		pages = (len(nonEmptyLines) + 1) / 4
	} else {
		pages = len(nonEmptyLines) / 4
	}

	//формирование куплетов(4 строки)
	var verse [][]string
	for i := 0; i < len(nonEmptyLines); i += 4 {
		end := i + 4
		if end > len(nonEmptyLines) {
			end = len(nonEmptyLines)
		}
		verse = append(verse, nonEmptyLines[i:end])
	}

	//пагинация
	start := offset
	end := offset + limit
	if start > len(verse) {
		start = len(verse)
	}
	if end > len(verse) {
		end = len(verse)
	}
	RequiredVerse := verse[start:end]
	c.JSON(http.StatusOK, gin.H{
		"verse": RequiredVerse,
		"page":  pages,
	})
}
