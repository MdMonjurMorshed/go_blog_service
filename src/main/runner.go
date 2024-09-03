package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
    router.GET("/", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, "Blog service is oks") })
	router.Run(":3000")

}
