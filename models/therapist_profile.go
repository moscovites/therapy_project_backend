package models 

import (
	"core/database"

)

type TherapistProfile struct {
	UserID uint
	User User `gorm: "foreignKey:UserID" json:user`
	FirstName string `gorm: "not null;" json: "firstName"`
	LastName string `gorm: "not null;" json: "lastName"`
	YearsOfExperience int `gorm: "not null;" json: "yearsOfExperience"`
	Specialty string `gorm: "not null;" json: "specialty"`
	Gender string `gorm: "not null;" json: "gender"`
	Sexuality string `json: "sexuality"`
	Religious bool `json: "religious`
	ReligiousDenomination bool `json: "religiousDenomination`
}

func (therapistProfile *TherapistProfile) Save() (*TherapistProfile, error) {
	err := database.Database.Create(&therapistProfile).Error
	if err != nil {
		return nil, err
	}
	return therapistProfile, nil
} 
