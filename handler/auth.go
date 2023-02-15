package handler

import (
	"log"

	"github.com/Suryaakula888/database"
	"github.com/Suryaakula888/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(in *gin.Context) {
	user := models.Book{}
	in.BindJSON(&user)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	err = database.SignUp(h.DB, &user)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
}
token , err :=generateToken (user)
if err != nil {
	log.Println(err)
	in.JSON(http.StatusInternalServerError, err)
}
in.Json(http.StatusOk,token)

func generateToken(user models.Users)(string error){
	token:=jwt.New(jwt.signMethodHS256)
	claims:=token.Claims(jwt.MapClaims)
	claims["authorized"]=true
	claims["emailid"]=user.EmilId
	claims["exp"]=time.Now.Add(time.Minute * 30).Unix()
	tokenString ,err := token.SignedString("rahul")
	if err != nil {
		log.Println(err)
		return "",err
	}
	return tokenString,nil
}