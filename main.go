package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"account-service/src/config"
	"account-service/src/core"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}

	core.InitDB()
	defer core.GetDB().Close()

	core.InitRedis()

	router := gin.Default()

	router.NoRoute(func(c *gin.Context) {
		r := fmt.Sprintf("%s%s Not Found", c.Request.Host, c.Request.URL)
		c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "error": r, "message": http.StatusText(http.StatusNotFound)})
	})

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"apps":    "Account Service",
			"version": config.Version,
		})
	})

	port := os.Getenv("PORT")
	router.Run(port)
}
