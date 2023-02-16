package main

import (
	"properties/router"
	"properties/utils/cache"
	"time"

	"properties/utils/db"

	cors "github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"

	"fmt"
)

var (
	gin_router *gin.Engine
)

func init() {
	gin_router = gin.Default()
	gin_router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}

func main() {
	gin_router = gin.Default()
	router.MapUrls(gin_router)
	err := db.Init_db()
	defer db.Disconect_db()
	cache.Init_cache()

	if err != nil {
		fmt.Println("Cannot init db")
		fmt.Println(err)
		return
	}
	fmt.Println("Starting server")
	gin_router.Run(":8090")
}
