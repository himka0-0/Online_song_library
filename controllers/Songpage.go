package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"online_song/config"
	"online_song/logger"
	"online_song/models"
	"strconv"
	"strings"
)

// SongPage godoc
// @Summary      Получение списка песен с фильтрацией и пагинацией
// @Description  Возвращает отфильтрованный список песен. Параметры запроса позволяют фильтровать по музыкальной группе, дате релиза, названию, тексту и ссылке. Также поддерживается пагинация через параметры limit и offset.
// @Tags         Songs
// @Accept       json
// @Produce      json
// @Param        grope  query    string  false  "Музыкальная группа"
// @Param        date   query    string  false  "Дата релиза"
// @Param        song   query    string  false  "Название песни"
// @Param        text   query    string  false  "Текст песни"
// @Param        link   query    string  false  "Ссылка на песню"
// @Param        limit  query    int     false  "Количество записей на страницу (по умолчанию 10)"
// @Param        offset query    int     false  "Смещение (по умолчанию 0)"
// @Success      200    {object} models.ResponseSongs "Успешный ответ с данными и длиной списка"
// @Failure      500    {object} map[string]string    "Ошибка поиска данных в БД"
// @Router       / [get]
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
		logger.Logger.Error("Ошибка поиска данных в бд", zap.Error(err))
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
