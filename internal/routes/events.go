package routes

import (
	"BookingService/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getEvents(ctx *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch data, please try again later"})
	}
	ctx.JSON(http.StatusOK, events)
}

func createEvent(ctx *gin.Context) {

	var event models.Event
	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request data"})
		return
	}
	event.UserID = ctx.GetInt("userID")
	err = event.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch data, please try again later"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})
}

func getEvent(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request data, invalid id"})
		return
	}
	result, err := models.GetEventById(int(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request data, unknown id"})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func deleteEvent(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request data, invalid id"})
		return
	}
	event, err := models.GetEventById(int(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request data, invalid id"})
		return
	}
	if event.UserID != ctx.GetInt("userId") {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "You are not the owner"})
		return
	}

	err = models.DeleteEvent(int(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch data, please try again later"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Event was deleted"})
}

func updateEvent(ctx *gin.Context) {
	var event models.Event
	err := ctx.ShouldBindJSON(&event)

	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request data, invalid id"})
		return
	}
	event.ID = int(id)

	oldEvent, err := models.GetEventById(event.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request data, invalid id"})
		return
	}
	if oldEvent.UserID != ctx.GetInt("userId") {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "You are not the owner"})
		return
	}

	err = models.UpdateEventById(event)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch data, please try again later"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Event was updated"})
}
