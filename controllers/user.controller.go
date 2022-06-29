package controllers

import (
	"net/http"

	"github.com/TomasHansut/gin_api/models"
	"github.com/TomasHansut/gin_api/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

// Constructor
func New(userservice services.UserService) UserController {
	return UserController{
		UserService: userservice,
	}
}

// Checks for error from user.service.impl if no error send 200 gin.Context hold information about request that we are gona send
func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	// Stores the request body into the context, and reuse when it is called again.
	if err := ctx.ShouldBindJSON(&user); err != nil {
		// ctx.json will send response
		ctx.JSON(http.StatusBadRequest, gin.H{"message: ": err.Error()})
		return
	}
	err := uc.UserService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message: ": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message: ": "success"})
}

// Get user by name
func (uc *UserController) GetUser(ctx *gin.Context) {
	username := ctx.Param("name")
	user, err := uc.UserService.GetUser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message: ": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// Get all users
func (uc *UserController) GetAll(ctx *gin.Context) {
	users, err := uc.UserService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message: ": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// Update user
func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message: ": err.Error()})
		return
	}
	err := uc.UserService.UpdateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message: ": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message: ": "success"})
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	username := ctx.Param("name")
	err := uc.UserService.DeleteUser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message: ": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message: ": "success"})
}

// grouping all routes under one name=user
func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/user")
	userroute.POST("/create", uc.CreateUser)
	userroute.GET("/get/:name", uc.GetUser)
	userroute.GET("/getall", uc.GetAll)
	userroute.PATCH("/update", uc.UpdateUser)
	userroute.DELETE("/delete/:name", uc.DeleteUser)
}
