package utils

// import (
// 	"core/models"
// 	"core/database"
// 	"gorm.io/gorm"
// )


// func (userId int) CreateMatch(*gorm.DB) error {
//     var patientProfiles []PatientProfile
//     var therapistProfiles []TherapistProfile

//     // Find patient profiles with a specific religious value
//     if err := database.Database.Where("religious = ?", true).Find(&patientProfiles).Error; err != nil {
//         return err
//     }

//     // Find therapist profiles with the same religious value
//     if err := database.Database.Where("religious = ?", true).Find(&therapistProfiles).Error; err != nil {
//         return err
//     }

//     // Iterate through matching profiles and create matches
//     for _, patientProfile := range patientProfiles {
//         for _, therapistProfile := range therapistProfiles {
//             // Check if the patients and therapists have the same religious value
//             if patientProfile.Religious == therapistProfile.Religious {
//                 match := Match{
//                     PatientID:   patientProfile.UserID,
//                     TherapistID: therapistProfile.UserID,
//                     Patient:     patientProfile,
//                     Therapist:   therapistProfile,
//                 }

//                 // Create the match
//                 if err := database.Database.Create(&match).Error; err != nil {
//                     return err
//                 }
//             }
//         }
//     }

//     return nil
// }