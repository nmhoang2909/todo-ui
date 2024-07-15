package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nmhoang2909/todo-ui/internal"
)

func main() {
	router := gin.Default()

	//initialize config
	app := internal.Config{Router: router}

	//routes
	app.Routes()

	router.Run(":8080")
}
