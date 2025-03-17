package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"online_song/config"
	"online_song/models"
	"strconv"
	"strings"
)

func SongPage(c *gin.Context) {
	//параметры фильтрации
	filterGrope := strings.ToLower(c.Query("grope"))
	filterReleaseDate := strings.ToLower(c.Query("date"))
	filterSong := strings.ToLower(c.Query("song"))
	filterText := strings.ToLower(c.Query("text"))
	filterLink := strings.ToLower(c.Query("link"))

	//пагинация по-умолчанию
	limit := 10
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

	//фомирование запроса в бд
	query := config.DB.Model(&models.Songs{})
	if filterGrope != "" {
		query = query.Where("Muzgroup = ?", filterGrope)
	}
	if filterReleaseDate != "" {
		query = query.Where("Release_Date = ?", filterReleaseDate)
	}
	if filterSong != "" {
		query = query.Where("Song = ?", filterSong)
	}
	if filterText != "" {
		query = query.Where("Text = ?", filterText)
	}
	if filterLink != "" {
		query = query.Where("Link = ?", filterLink)
	}

	//запрос в бд
	var filtered_songs []models.Songs
	if err := query.Find(&filtered_songs).Error; err != nil {
		log.Println("Ошибка поиска данных в бд", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Сервер не нашел данных"})
		return
	}
	//пагинация
	begin := offset
	end := offset + limit
	songListLength := len(filtered_songs)
	if begin > songListLength {
		begin = songListLength
	}
	if end > songListLength {
		end = songListLength
	}
	data := filtered_songs[begin:end]

	//клиенту
	c.JSON(http.StatusOK, models.ResponseSongs{
		Data:   data,
		Length: songListLength,
	})
}
