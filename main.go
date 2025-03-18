package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"online_song/config"
	"online_song/controllers"
	"online_song/logger"
	"online_song/models"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}
}
func main() {
	//Подключение логера
	logger.InitLogger()
	defer logger.Logger.Sync()
	logger.Logger.Info("Приложение запущено")

	//подключение бд
	config.InitDB()
	config.DB.AutoMigrate(&models.Songs{})

	r := gin.Default()
	//унес бы в routers,но пока мало и так сойдет
	r.GET("/", controllers.SongPage)
	r.GET("/verse/:id", controllers.VersePage)
	r.POST("/", controllers.CreateSongHandler)
	r.PUT("/", controllers.ChangeSongHandler)
	r.DELETE("/", controllers.DeleteSongHandler)
	r.Run(":8080")
}
