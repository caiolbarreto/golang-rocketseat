package controllers

import (
	"library-mvc/internal/loans/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LoanController struct {
	loanService models.LoanService
}

func NewLoanController(loanService models.LoanService) *LoanController {
	return &LoanController{
		loanService: loanService,
	}
}

func (l *LoanController) RegisterRoutes(router *gin.Engine) {
	loans := router.Group("/loans")
	{
		loans.POST("/", l.CreateLoan)
		loans.GET("/", l.GetAllLoans)
		loans.GET("/:id", l.GetLoan)
		loans.PUT("/:id/return", l.ReturnBook)
	}

	users := router.Group("/users")
	{
		users.GET("/:userId/loans", l.GetUserLoans)
	}
}

func (l *LoanController) CreateLoan(ctx *gin.Context) {
	var request struct {
		BookID int64 `json:"bookId"`
		UserID int64 `json:"userId"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	loan, err := l.loanService.CreateLoan(request.BookID, request.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, loan)
}

func (l *LoanController) GetLoan(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid loan id"})
		return
	}

	loan, err := l.loanService.GetLoan(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, loan)
}

func (l *LoanController) GetAllLoans(ctx *gin.Context) {
	loans, err := l.loanService.GetAllLoans()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, loans)
}

func (l *LoanController) GetUserLoans(ctx *gin.Context) {
	userId, err := strconv.ParseInt(ctx.Param("userId"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	loans, err := l.loanService.GetUserLoans(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, loans)
}

func (l *LoanController) ReturnBook(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid loan id"})
		return
	}

	err = l.loanService.ReturnBook(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
