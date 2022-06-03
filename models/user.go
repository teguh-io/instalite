package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	ID           int `gorm:"primarykey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Username     string        `gorm:"column:username;type:varchar;size:255;unique;not null" valid:"required~username is blank"`
	Email        string        `gorm:"column:email;type:varchar;size:255;unique;not null" valid:"required~email is blank,email~Invalid email format"`
	Password     string        `gorm:"column:password;not null" valid:"required~password field is blank,length(6|255)~Password has to have a minimum length of 6 characters"`
	Age          int           `gorm:"column:age;type:int;not null"`
	Photos       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Comments     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SocialMedias []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(u)
	return err
}
