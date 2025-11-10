package controllers

import "github.com/gin-gonic/gin"

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) RegisterRoutes(router *gin.Engine) {
	user := router.Group("/users")
	{
		user.POST("/", c.CreateUser)
		user.GET("/", c.GetAllUsers)
		user.GET("/:id", c.GetUser)
		user.PUT("/:id", c.UpdateUser)
		user.DELETE("/:id", c.DeleteUser)
	}
}

func (c *UserController) CreateUser(ctx *gin.Context)  {}
func (c *UserController) GetUser(ctx *gin.Context)     {}
func (c *UserController) GetAllUsers(ctx *gin.Context) {}
func (c *UserController) UpdateUser(ctx *gin.Context)  {}
func (c *UserController) DeleteUser(ctx *gin.Context)  {}
