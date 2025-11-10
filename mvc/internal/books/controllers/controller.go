package controllers

import "github.com/gin-gonic/gin"

type BookController struct {
}

func NewBookController() *BookController {
	return &BookController{}
}

func (c *BookController) RegisterRoutes(router *gin.Engine) {
	book := router.Group("/books")
	{
		book.POST("/", c.CreateBook)
		book.GET("/", c.GetAllBooks)
		book.GET("/:id", c.GetBook)
		book.PUT("/:id", c.UpdateBook)
		book.DELETE("/:id", c.DeleteBook)
	}
}

func (c *BookController) CreateBook(ctx *gin.Context)  {}
func (c *BookController) GetBook(ctx *gin.Context)     {}
func (c *BookController) GetAllBooks(ctx *gin.Context) {}
func (c *BookController) UpdateBook(ctx *gin.Context)  {}
func (c *BookController) DeleteBook(ctx *gin.Context)  {}
