package railtown

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"railtown/domain/railtown"
	"railtown/logger"
	"railtown/services"
	"railtown/utills/responses"
)

func RailController(c *gin.Context) {
	logger.Info("Enter to RailController successfully")

	req := railtown.RailTownREQ{
		Query: c.Query("q"),
	}

	data, serviceErr := services.RailTownsService.GETTemperatureService(&req)
	if serviceErr != nil {
		c.JSON(serviceErr.Status(), serviceErr)
		return
	}

	ok := responses.NewRequestSuccessOk("Temperature data sent successfully", "", data)
	c.JSON(http.StatusOK, ok)
	logger.Info("Close from RailController successfully")
}
