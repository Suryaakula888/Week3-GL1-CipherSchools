package main

import (
	"log"

	"github.com/Suryaakula888/database"
	"github.com/Suryaakula888/routers"
	"github.com/Suryaakula888/handler"

)

func init(){
	database.Setup()
}
func init(){
	fmt.Print(1)
}
func init(){
	fmt.Print(2)
}

func main() {
	engine:= gin.Default()
	api := handler.Handler{
		DB: database.GetDB(),
	database.Setup()
	engine := engine.Router()
	routers.BookRouter(engine,api)
	routers.AuthRouter(engine,api)
	err := engine.Run("127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	}
}
