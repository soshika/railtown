package app

import (
	"railtown/controllers/ping"
	"railtown/controllers/railtown"
)

func urlPatterns() {

	router.GET("/ping", ping.Pong)

	/////////////////////////////
	////Authentication request///
	////////////////////////////

	router.GET("rail-town/weather", railtown.RailController)

}
