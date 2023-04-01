package main

import "github.com/gin-gonic/gin"

func main() {
	memStorage := NewMemoryStorage()
	handler := NewHandler(memStorage)

	router := gin.Default()

	router.POST("/sotrudnik", handler.CreateSotrudnik)
	router.GET("/sotrudnik/:id", handler.GetSotr)
	router.PUT("/sotrudnik/:id", handler.UpdateSotr)
	router.DELETE("/sotrudnik/:id", handler.DeleteSotr)

	router.Run(":80")
}
