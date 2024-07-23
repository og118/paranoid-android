package main

import (
	"paranoid_android/exporters"

	"github.com/gin-gonic/gin"
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

	router.Run(":8080")
}
