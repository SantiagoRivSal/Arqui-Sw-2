package main

import (
	"properties/router"
	"properties/utils/cache"

	"properties/utils/db"

	"github.com/gin-gonic/gin"

	"fmt"
)

var (
	gin_router *gin.Engine
)

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
