package main

import (
	"github.com/gin-gonic/gin"
	"log"
	routers "sonicfinger/routers"
)

// router 객체 라우터를 만드는 기본적인 방법
var (
	router = gin.Default()
)

func main() {
	router.Static("/static", "./static")
	router.LoadHTMLGlob("view/*") //템플릿을 불러오는 방법
	router.GET("/", routers.Home)
	router.GET("/auth", routers.Auth)

	log.Fatal(router.Run(":8080"))
}
