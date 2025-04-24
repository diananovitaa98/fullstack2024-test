package handler

import (
	"fullstacktest/entity"
	"fullstacktest/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ClientHandler struct {
	uc usecase.ClientUsecaseItf
}

func NewClientHandler(uc usecase.ClientUsecaseItf) ClientHandler {
	return ClientHandler{
		uc: uc,
	}
}

func (h ClientHandler) GetClients(ctx *gin.Context) {
	clients, err := h.uc.SelectAllClient(ctx)
	if err != nil {
		log.Println("GetClients error:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch clients"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success get clients",
		"data":    clients,
	})
}

func (h ClientHandler) InsertClient(ctx *gin.Context) {
	var req entity.MyClient
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid input", "error": err.Error()})
		return
	}

	id, err := h.uc.InsertClient(ctx, &req)
	if err != nil {
		log.Println("InsertClient error:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to insert client"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success create client",
		"id":      id,
	})
}

func (h ClientHandler) UpdateClient(ctx *gin.Context) {
	var req entity.MyClient
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid input", "error": err.Error()})
		return
	}

	err := h.uc.UpdateClient(ctx, &req)
	if err != nil {
		log.Println("UpdateClient error:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to update client"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success update client",
	})
}

func (h ClientHandler) DeleteClient(ctx *gin.Context) {
	slug := ctx.Param("slug")
	if slug == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "slug is required"})
		return
	}

	err := h.uc.DeleteClient(ctx, slug)
	if err != nil {
		log.Println("DeleteClient error:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to delete client"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success delete client",
	})
}
