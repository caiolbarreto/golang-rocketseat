package controllers

import (
	"library-mvc/internal/books/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	bookService models.BookService
}

func NewBookController(bookService models.BookService) *BookController {
	return &BookController{
		bookService: bookService,
	}
}

func (b *BookController) RegisterRoutes(router *gin.Engine) {
	book := router.Group("/books")
	{
		book.POST("/", b.CreateBook)
		book.GET("/", b.GetAllBooks)
		book.GET("/:id", b.GetBook)
		book.PUT("/:id", b.UpdateBook)
		book.DELETE("/:id", b.DeleteBook)
	}
}

func (b *BookController) CreateBook(ctx *gin.Context) {
	var book models.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if err := b.bookService.CreateBook(&book); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, book)
}

func (b *BookController) GetBook(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid book id"})
		return
	}

	book, err := b.bookService.GetBook(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (b *BookController) GetAllBooks(ctx *gin.Context) {
	books, err := b.bookService.GetAllBooks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func (b *BookController) UpdateBook(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid book id"})
		return
	}

	var book models.Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	err = b.bookService.UpdateBook(id, &book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (b *BookController) DeleteBook(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid book id"})
		return
	}

	err = b.bookService.DeleteBook(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
