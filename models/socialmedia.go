package models

type SocialMedia struct {
	ID             int    `gorm:"primarykey"`
	Name           string `gorm:"column:name;type:varchar;size:255"`
	SocialMediaUrl string `gorm:"column:social_media_url;type:varchar"`
	UserID         int    `gorm:"column:user_id;type:int"`
}
