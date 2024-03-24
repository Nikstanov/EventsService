package main

import (
	"BookingService/internal/db"
	"BookingService/internal/routes"
	"BookingService/internal/utills"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

const ServerPort = "SERVER_PORT"

func main() {
	server := gin.Default()

	utills.InitJWT()
	err := db.InitDB()
	if err != nil {
		panic(fmt.Sprintf("error with database: %s", err))
	}

	routes.RegisterRoutes(server)

	port, exists := os.LookupEnv(ServerPort)
	if !exists {
		panic("The secret key is not set")
	}
	err = server.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		panic("Server didn't start")
	}
}
