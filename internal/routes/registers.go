package routes

import (
	"BookingService/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func register(ctx *gin.Context) {
	userID := ctx.GetInt("userID")
	eventID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request data, invalid id"})
		return
	}
	event, err := models.GetEventById(int(eventID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Cant fetch event"})
		return
	}
	err = event.Registration(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register event"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "registered"})

}
func cancelRegistration(ctx *gin.Context) {
	userID := ctx.GetInt("userID")
	eventID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request data, invalid id"})
		return
	}
	event, err := models.GetEventById(int(eventID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Cant fetch event"})
		return
	}
	err = event.DeleteRegistration(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register event"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "canceled"})
}
