package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	FirstName string `gorm:"type:varchar(60);column:first_name;not null" json:"first_name"`
	LastName string `gorm:"type:varchar(60);column:last_name;not null" json:"first_name"`
	Email string `gorm:"type:varchar(100);column:email;unique;not null" json:"email"`
}

func (user *User) FullName() string {
	return user.FirstName + " " + user.LastName
}

func (User) TableName() string {
	return "users"
}