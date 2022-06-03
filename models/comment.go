package models

import "time"

type Comment struct {
	ID            int `gorm:"primarykey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	UserID        int    `gorm:"column:user_id;type:int"`
	PhotoID       int    `gorm:"column:photo_id;type:int"`
	Message       string `gorm:"column:message;type:varchar"`
	PhotosID      int    `gorm:"-:migration;->;"`
	PhotosTitle   string `gorm:"-:migration;->;"`
	PhotosCaption string `gorm:"-:migration;->;"`
	PhotosUrl     string `gorm:"-:migration;->;"`
	PhotosUserID  int    `gorm:"-:migration;->;"`
	UsersID       int    `gorm:"-:migration;->;"`
	UsersEmail    string `gorm:"-:migration;->;"`
	UsersName     string `gorm:"-:migration;->;"`
}
