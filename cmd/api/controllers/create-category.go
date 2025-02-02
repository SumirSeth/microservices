package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sumirseth/microservices/internal/repositories"
	use_cases "github.com/sumirseth/microservices/internal/use-cases"
)

type createCategoryInput struct {
	Name string `json:"name" binding:"required"`
}


func CreateCategory(ctx *gin.Context, repository repositories.ICategoryRepository) {

	var body createCategoryInput

	if err:= ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"success": false,
				"error": err.Error(),
			})
		
		return
	}

	useCase := use_cases.NewCreateCategoryUseCase(repository)
	err := useCase.Execute(body.Name)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"success": false,
				"error": err.Error(),
			})
		
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
	})
}