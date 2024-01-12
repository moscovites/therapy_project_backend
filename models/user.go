package models

import (
	"gorm.io/gorm"
	"core/database"
	"golang.org/x/crypto/bcrypt"
	"fmt"
)


type User struct {
	gorm.Model
	ID uint `gorm: "primaryKey"`
	UserType string `gorm: "not null;" json: "userType"`
	PhoneNumber string `gorm: "size: 255" json: "phoneNumber"`
	Email string `gorm: "size: 999; unique; not null" json: "email"`
	Password string `gorm: "size: 255; not null;" json: "-"`
	VerificationCode string `gorm: "size: 8;" json: "verification_code"`
	ResetPasswordCode string `gorm: "size: 8;" json: "passwordResetCode"`
	IsVerified bool `gorm: "default:false"`
}

func (user *User) BeforeSave(*gorm.DB) error {
	if user.ID == 0 {

		passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
}
	return nil
}

func (user *User) Save() (*User, error) {
	err := database.Database.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func FindUserByEmail(email string) (User, error) {
	var user User
	err := database.Database.Where("email", email).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func FindUserById(id uint) (User, error) {
	var user User
	err := database.Database.Where("ID=?", id).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (user *User) ValidatePassword(password string) error {
	fmt.Println("plain password", password)
	//fmt.Println("stored hashed pasword", user.Password)
	//hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	//if err != nil {
	//	return err
	//}
	//fmt.Println(bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)))
	//fmt.Println("Hashed Password Provided:", string(hashedPassword))
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func (user *User) UpdatePassword(input string) (*User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	
	if err := database.Database.Model(&user).Update("password", passwordHash); err != nil  {
		return &User{}, nil
	}
	
	return user, nil
}
