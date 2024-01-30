package models 

import (
	"core/database"
	"gorm.io/gorm"

)

type TherapistProfile struct {
	gorm.Model
	UserID uint
	User User `gorm: "foreignKey:UserID" json:user`
	FirstName string `gorm: "not null;" json: "firstName"`
	LastName string `gorm: "not null;" json: "lastName"`
	YearsOfExperience int `gorm: "not null;" json: "yearsOfExperience"`
	Specialty string `gorm: "not null;" json: "specialty"`
	Gender string `gorm: "not null;" json: "gender"`
	Sexuality string `json: "sexuality"`
	Religious bool `json: "religious`
	ReligiousDenomination string `json: "religiousDenomination`
}

func FindTherapistProfileById(id uint) (TherapistProfile, error) {
	var therapistProfile TherapistProfile
	err := database.Database.Where("ID=?", id).Find(&therapistProfile).Error
	if err != nil {
		return TherapistProfile{}, err
	}
	return therapistProfile, nil
}

func (therapistProfile *TherapistProfile) Save() (*TherapistProfile, error) {
	err := database.Database.Create(&therapistProfile).Error
	if err != nil {
		return nil, err
	}
	return therapistProfile, nil
} 


func (therapistProfile *TherapistProfile) Update(input *TherapistProfile) (*TherapistProfile, error) {
	err := database.Database.Model(&therapistProfile).Updates(input).Error
	if err != nil {
		return nil, err
	}
	return therapistProfile, nil
}