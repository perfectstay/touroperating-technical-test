package main

import (
	"technical-test/cmd/handlers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	router.POST("/quote", handlers.RetrieveStayQuoteHandler)
}
