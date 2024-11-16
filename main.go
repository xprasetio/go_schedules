package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xprasetio/go_schedules.git/modules/station"
)

func main() {
	InitStationRoutes()
}

func InitStationRoutes() {
	var (
		router = gin.Default()
		api    = router.Group("/v1/api")
	)
	station.InitStationRoutes(api)
	router.Run(":8089")
}
