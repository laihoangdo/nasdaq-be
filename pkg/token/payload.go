package token

import (
	"errors"
	"time"
)

var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token is invalid")
)

type Payload struct {
	TokenID   string    `json:"jti"`
	ID        string    `json:"id"`
	UserUUID  string    `json:"user_id"`
	Profile   Profile   `json:"profile"`
	Type      string    `json:"type"`
	IssueAt   time.Time `json:"issue_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

type Profile struct {
	UUID     string `json:"uuid"`
	RoleSlug string `json:"role_slug"`
}

func NewPayload(tokenID string, sessionID, userID, pType string, duration time.Duration) Payload {
	return Payload{
		TokenID:   tokenID,
		ID:        sessionID,
		UserUUID:  userID,
		IssueAt:   time.Now(),
		ExpiredAt: time.Now().Add(duration),
		Type:      pType,
	}
}

func (p Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
