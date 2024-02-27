package models

import "time"

type UserStatus string

func (us UserStatus) String() string {
	return string(us)
}

type User struct {
	ID        int       `gorm:"column:id" json:"id"`
	Sku       int       `gorm:"column:sku" json:"sku"`
	UserName  string    `gorm:"column:username" json:"username"`
	Password  string    `gorm:"column:password" json:"password"`
	Name      string    `gorm:"column:name" json:"name"`
	Phone     string    `gorm:"column:phone" json:"phone"`
	Email     string    `gorm:"column:email" json:"email"`
	Status    bool      `gorm:"column:status" json:"status"`
	Balance   int       `gorm:"column:balance" json:"balance"`
	Lock      bool      `gorm:"column:lock" json:"lock"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type AdminUserCredentials struct {
	AccessToken           string
	AccessTokenExpiresAt  time.Time
	RefreshToken          string
	RefreshTokenExpiresAt time.Time
	AdminInfo             User
}

func (*User) TableName() string {
	return "users"
}
