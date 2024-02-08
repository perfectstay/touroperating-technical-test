package main

import(
	"github.com/gin-gonic/gin"
	"net/http"
)

func handlerQuote(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "hello world"})    
  }