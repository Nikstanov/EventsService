package middlewares

import (
	"BookingService/internal/utills"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}
	id, err := utills.VerifyToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized", "error": err})
		return
	}

	ctx.Set("userID", id)
	ctx.Next()
}
