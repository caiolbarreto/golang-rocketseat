package controllers

import "github.com/gin-gonic/gin"

type LoanController struct {
}

func NewLoanController() *LoanController {
	return &LoanController{}
}

func (c *LoanController) RegisterRoutes(router *gin.Engine) {
	loan := router.Group("/loans")
	{
		loan.POST("/", c.CreateLoan)
		loan.GET("/", c.GetAllLoans)
		loan.GET("/:id", c.GetLoan)
		loan.PUT("/:id", c.UpdateLoan)
		loan.DELETE("/:id", c.DeleteLoan)
	}
}

func (c *LoanController) CreateLoan(ctx *gin.Context)  {}
func (c *LoanController) GetLoan(ctx *gin.Context)     {}
func (c *LoanController) GetAllLoans(ctx *gin.Context) {}
func (c *LoanController) UpdateLoan(ctx *gin.Context)  {}
func (c *LoanController) DeleteLoan(ctx *gin.Context)  {}
