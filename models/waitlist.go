package models 

import (
	"gorm.io/gorm"
	"core/database"
)

type Waitlist struct {
	gorm.Model
	FullName string `gorm: "size:255" json: "fullName"`
	Email string `gorm: "size:255" json: "email"`
}

func (waitlistMember *Waitlist) Save() (*Waitlist, error) {
	err := database.Database.Create(&waitlistMember).Error
	if err != nil {
		return &Waitlist{}, err
	}
	return waitlistMember, nil
}