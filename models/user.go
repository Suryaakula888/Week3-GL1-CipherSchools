package models

import "github.com/jinzhu/gorm"

type user struct {
	gorm.Model
	EmailId    string
	Name      string
	Password string
	Age int
	
}
