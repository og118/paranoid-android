package main

import (
	"paranoid_android/exporters"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	router := gin.Default()
	Init()

	router.GET("/issue/:issue_id", func(ctx *gin.Context) {
		exporters.ExportIssue(ctx, C)
	})

	router.GET("/issue/:issue_id/article/:article_id", func(ctx *gin.Context) {
		exporters.ExportIssueArticle(ctx, C)
	})

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":8080")
}
