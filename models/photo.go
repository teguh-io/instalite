package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	ID        int `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string    `gorm:"column:title;type:varchar" valid:"required~title is blank"`
	Caption   string    `gorm:"column:caption;type:varchar"`
	PhotoUrl  string    `gorm:"column:photo_url;type:varchar" valid:"required~photo url is blank"`
	UserID    int       `gorm:"column:user_id;type:int"`
	Comments  []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	User      *User     `gorm:"-:migration;->;embedded;"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(p)
	return err
}
