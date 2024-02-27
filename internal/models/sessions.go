package models

import (
	"fmt"
	"time"
)

const (
	blacklistSessionPrefix = "api-blacklist-session"
)

type Session struct {
	ID        int64      `json:"id" gorm:"column:id"`
	UUID      string     `json:"uuid" gorm:"column:uuid"`
	UserUUID  string     `json:"user_uuid" gorm:"column:user_uuid"`
	UserAgent string     `json:"user_agent,omitempty" gorm:"column:user_agent"`
	ClientIP  string     `json:"client_ip,omitempty" gorm:"column:client_ip"`
	ExpiredAt time.Time  `json:"expired_at" gorm:"column:expired_at"`
	CreatedAt time.Time  `json:"created_at,omitempty" gorm:"column:created_at;default:now()"`
	UpdatedAt time.Time  `json:"updated_at,omitempty" gorm:"column:updated_at;default:now()"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"column:deleted_at"`
}

func (s *Session) TableName() string {
	return "sessions"
}

func (s Session) IsValid() bool {
	if s.UUID == "" || s.UserUUID == "" || s.ExpiredAt.IsZero() || s.DeletedAt != nil {
		return false
	}
	return true
}

func GetSessionKeyWithPrefix(sessionId string) string {
	return fmt.Sprintf("%s:%s", blacklistSessionPrefix, sessionId)
}
