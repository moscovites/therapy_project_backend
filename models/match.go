package models

import (
	//"gorm.io/gorm"
	"core/database"
	// "time"
	"gorm.io/gorm"
)

type Match struct {
	gorm.Model
	PatientID uint
	TherapistID uint
	Patient PatientProfile `gorm: "foreignKey:PatientID" json:patient`
	Therapist TherapistProfile `gorm: "foreignKey:TherapistID" json:therapist`
}

func FindMatchById(id uint) (Match, error) {
	var match Match
	err := database.Database.Where("ID=?", id).Find(&match).Error
	if err != nil {
		return Match{}, err
	}
	return match, nil
}

func (match * Match) Save() (*Match, error) {

	err := database.Database.Create(&match).Error

	if err != nil {
		return nil, err
	}
	return match, nil
}