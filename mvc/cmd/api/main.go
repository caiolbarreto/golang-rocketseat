package main

import (
	"log"

	"library-mvc/internal/users/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	usersController := controllers.NewUserController()
	usersController.RegisterRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
