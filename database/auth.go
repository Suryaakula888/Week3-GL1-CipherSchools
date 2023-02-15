package database

import (
	"github.com/Suryaakula888/models"
	"github.com/jinzhu/gorm"
)

func SignUp(db *gorm.DB, user *models.users) error {
	err := db.Save(user).Error
	if err != nil {
		return err
	}
	return nil
}

func SignIn(db *gorm.DB, user models.User) error {
	getUser := models.users{}
	err := db.Select("users.*").Group("users.email_id").Where("users.email_id=?", user.EmailId).First(&book).Error
	if err != nil {
		return err
	}
	log.Println(getUser)
	if user.Password != getUser.Password {
		return err.New("Pasword Incorrect")
	}
	return  nil
}
