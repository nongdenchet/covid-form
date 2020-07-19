package model

import "github.com/jinzhu/gorm"

type Visit struct {
	gorm.Model
	Name  string
	Phone string
	Email *string
}
