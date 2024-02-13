package main

import (
	"github.com/gin-gonic/gin"
)

type Application struct {
	router *gin.Engine
}

func main() {

	app := Application{
		router: gin.Default(),
	}

	Routes(app.router)

	app.router.Run()

}
