package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:message`
}

type Handler struct {
	storage Storage
}

func NewHandler(storage Storage) *Handler {
	return &Handler{storage: storage}
}

func (h *Handler) CreateSotrudnik(c *gin.Context) {
	var sotrudnik Sotrudnik

	if err := c.BindJSON(&sotrudnik); err != nil {
		fmt.Printf("failed to bind sotr: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	h.storage.Insert(&sotrudnik)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": sotrudnik.Id,
	})
}

func (h *Handler) UpdateSotr(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("Failed to convert id param to int: %v\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	var sotrudnik Sotrudnik

	if err := c.BindJSON(&sotrudnik); err != nil {
		fmt.Printf("failed to bind sotr: %v\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	h.storage.Update(id, sotrudnik)

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": sotrudnik.Id,
	})
}

func (h *Handler) GetSotr(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	sotrudnik, err := h.storage.Get(id)
	if err != nil {
		fmt.Printf("failed to get sotr %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, sotrudnik)
}

func (h *Handler) DeleteSotr(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("failed to convert id param to int: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	h.storage.Delete(id)

	c.String(http.StatusOK, "sotrudnik deleted")
}
