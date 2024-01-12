package models

import (
	//"gorm.io/gorm"
	"core/database"

	"gorm.io/gorm"
)

type PatientProfile struct {
	gorm.Model
	UserID uint
	User User `gorm: "foreignKey:UserID" json:user`
	TherapyType string `json: "therapyType"`
	Name string `gorm: "size:255" json: "firstName"`
	Age int `json: "age"`
	Gender string `json: "gender"`
	Sexuality string `json: "sexuality"`
	RelationshipStatus string `json: "relationshipStatus"`
	Religious bool `json: "religious`
	ReligiousDenomination bool `json: "religiousDenomination`
	BeenInTherapyBefore bool `json: "beenInTherapyBefore"`	
}

func (patientProfile * PatientProfile) Save() (*PatientProfile, error) {

	err := database.Database.Create(&patientProfile).Error

	if err != nil {
		return nil, err
	}
	return patientProfile, nil
}