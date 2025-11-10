package main

import (
	"log"

	bookController "library-mvc/internal/books/controllers"
	loansController "library-mvc/internal/loans/controllers"
	userController "library-mvc/internal/users/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	loansController := loansController.NewLoanController()
	loansController.RegisterRoutes(router)

	booksController := bookController.NewBookController()
	booksController.RegisterRoutes(router)

	usersController := userController.NewUserController()
	usersController.RegisterRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
