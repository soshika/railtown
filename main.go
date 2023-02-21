package main

import (
	"github.com/gin-gonic/gin"
	"railtown/app"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	app.StartApplication()
}
