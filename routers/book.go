package routers

import (
	handler "github.com/Suryaakula888/Handler"
	"github.com/Suryaakula888/database"
	"github.com/gin-gonic/gin"
)

func BookRouter(*gin.Engine api handler.Handler) {
	//router := gin.Default()
	//api := handler.Handler{
	// 	DB: database.GetDB(),

	router.GET("/Books", api.GetBooks)
	router.POST("/Books", api.SaveBook)
	return router
}
