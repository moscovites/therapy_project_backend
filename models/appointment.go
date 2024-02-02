package models

import (
	//"gorm.io/gorm"
	"core/database"
	// "time"
	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	PatientID uint
	TherapistID uint
	Patient User `gorm: "foreignKey:PatientID" json:patient`
	Therapist User `gorm: "foreignKey:TherapistID" json:therapist`
	TherapyType string `json: "therapyType"`
	StartTime   string `json: "startTime"`
	EndTime     string `json: "endTime"`
	Location    string `json: "location"`
}

func FindAppointmentById(id uint) (Appointment, error) {
	var appointment Appointment
	err := database.Database.Where("ID=?", id).Find(&appointment).Error
	if err != nil {
		return Appointment{}, err
	}
	return appointment, nil
}

func (appointment * Appointment) Save() (*Appointment, error) {

	err := database.Database.Create(&appointment).Error

	if err != nil {
		return nil, err
	}
	return appointment, nil
}

func (appointment * Appointment) Update(input *Appointment) (*Appointment, error) {
	err := database.Database.Model(&appointment).Updates(input).Error
	if err != nil {
		return nil, err
	}
	return appointment, nil
}