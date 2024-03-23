package routes

import (
	"BookingService/internal/models"
	"BookingService/internal/utills"
	"github.com/gin-gonic/gin"
	"net/http"
)

func signUp(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request data"})
		return
	}

	err = user.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch data, please try again later"})
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "user created", "event": user})
}

func login(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request data"})
		return
	}
	id, err := user.ValidateCredentials()
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	token, err := utills.GenerateToken(user.Email, int64(id))
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successful", "token": token})
}
