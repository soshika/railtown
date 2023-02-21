package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func StartApplication() {
	urlPatterns()
	// logger.Info("about to start the application V 1.0.2 !")

	router.Run(":9093")
}
