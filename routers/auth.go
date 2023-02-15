package routers

func AuthRouter(*gin.Engine api handler.Handler) {

	router.POST("/signUp", api.SignUp)
	router.POST("/signin", api.Signip)
}
