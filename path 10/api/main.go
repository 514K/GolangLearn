package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.POST("/sotrudnik")
	router.GET("/sotrudnik/:id")
	router.PUT("/sotrudnik/:id")
	router.DELETE("/sotrudnik/:id")

	router.Run()
}
