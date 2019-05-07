package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/myml/src/api/controllers/myml"
)

const (
	port = ":8080"
)

var (
	router = gin.Default() // Engine
)

func StartApp() {
	router.GET(
		"myml/:userID", myml.GetMyML)
	router.Run(port)
}
