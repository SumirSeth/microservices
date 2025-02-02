package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sumirseth/microservices/internal/repositories"
	use_cases "github.com/sumirseth/microservices/internal/use-cases"
)



func ListCategory(ctx *gin.Context, repository repositories.ICategoryRepository) {
	useCase := use_cases.NewlistCategoriesUseCase(repository)
	categories, err := useCase.Execute()

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
		"categories": categories,
	})
}