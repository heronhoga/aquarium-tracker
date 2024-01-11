package routes

import (
	"github.com/gin-gonic/gin"
	"go-restapi-gin-aquarium/controllers/unitcontroller"
)

func Route() {
	// TODO: Add routes
	r := gin.Default()

	r.GET("/aq/units", unitcontroller.Index)
	r.GET("/aq/unit/:id", unitcontroller.IndexOne)
	r.POST("aq/unit", unitcontroller.Create)
	r.PUT("aq/unit/:id", unitcontroller.Update)
	r.DELETE("aq/unit", unitcontroller.Delete)

	r.Run()
}