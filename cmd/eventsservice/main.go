package main

import (
	"BookingService/internal/db"
	"BookingService/internal/routes"
	"BookingService/internal/utills"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	utills.InitJWT()
	err := db.InitDB()
	if err != nil {
		panic(fmt.Sprintf("error with database: %s", err))
	}

	routes.RegisterRoutes(server)

	err = server.Run(":8080")
	if err != nil {
		panic("Server didn't start")
	}
}
