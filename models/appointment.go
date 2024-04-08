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
	// TherapyType string `json: "therapyType"`
	Date   string `json: "date"`
	StartTime   string `json: "startTime"`
	EndTime     string `json: "endTime"`
	Location    string `json: "location"`
	AddressOrId string `json: "addressOrId"`
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

func AllAppointments() ([]Appointment, error) {
    var appointments []Appointment
    if err := database.Database.Preload("Patient").Preload("Therapist").Find(&appointments).Error; err != nil {
        return nil, err
    }
    return appointments, nil
}

