package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sumirseth/microservices/cmd/api/controllers"
	"github.com/sumirseth/microservices/internal/repositories"
)

func CategoryRoutes(router *gin.Engine) {
	// v1 := router.Group("/categories")
	// v1.POST("/login", loginEndpoint)
	// v1.POST("/submit", submitEndpoint)
	// v1.POST("/read", readEndpoint)

	categoryRoutes := router.Group("/categories")
	inMemoryCategoryRepository := repositories.NewInMemoryCategoryRepository()
	categoryRoutes.POST("/", func (ctx *gin.Context) {
		controllers.CreateCategory(ctx, inMemoryCategoryRepository)
	})

	categoryRoutes.GET("/", func (ctx *gin.Context) {
		controllers.ListCategory(ctx, inMemoryCategoryRepository)
	})
}