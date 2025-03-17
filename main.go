package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"online_song/config"
	"online_song/controllers"
	"online_song/models"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}
}
func main() {
	config.InitDB()
	config.DB.AutoMigrate(&models.Songs{})
	r := gin.Default()
	r.GET("/", controllers.SongPage)
	r.POST("/", controllers.CreateSongHandler)
	r.DELETE("/", controllers.DeleteSongHandler)
	r.Run(":8080")
}
