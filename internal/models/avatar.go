package models

import "time"

// Avatar model
type Avatar struct {
	ID        int64     `gorm:"column:id" json:"id"`
	UUID      string    `gorm:"column:uuid" json:"uuid"`
	Gender    int       `gorm:"column:gender" json:"gender"`
	Url       string    `gorm:"column:url" json:"url"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// Name table avatar
func (Avatar) TableName() string {
	return "avatars"
}

// Filter avatar
type AvatarFilter struct {
	Gender string `json:"gender"`
}

// All News response
type AvatarList struct {
	Avatar []Avatar
}
