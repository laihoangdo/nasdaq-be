package models

import (
	"time"
)

const (
	RoleSlugPrimarySchoolStudent   RoleSlug = "primary-school-student"
	RoleSlugSecondarySchoolStudent RoleSlug = "secondary-school-student"
	RoleSlugHighSchoolStudent      RoleSlug = "high-school-student"
	RoleSlugParent                 RoleSlug = "parent"
	RoleSlugTeacher                RoleSlug = "teacher"
	RoleSlugOther                  RoleSlug = "other"
)

// RoleSlug represents for role slug string
type RoleSlug string

// IsTeacher checks if role slug is teacher
func (rslug RoleSlug) IsTeacher() bool {
	return rslug == RoleSlugTeacher
}

// IsStudent checks if role slug is student
func (rslug RoleSlug) IsStudent() bool {
	return rslug == RoleSlugPrimarySchoolStudent ||
		rslug == RoleSlugSecondarySchoolStudent ||
		rslug == RoleSlugHighSchoolStudent
}

// IsValid checks if role slug is valid
func (rslug RoleSlug) IsValid() bool {
	return rslug == RoleSlugPrimarySchoolStudent ||
		rslug == RoleSlugSecondarySchoolStudent ||
		rslug == RoleSlugHighSchoolStudent ||
		rslug == RoleSlugParent ||
		rslug == RoleSlugTeacher ||
		rslug == RoleSlugOther
}

// IsTeacher checks if role slug is teacher
func (rslug RoleSlug) IsParent() bool {
	return rslug == RoleSlugParent
}

// IsTeacher checks if role slug is teacher
func (rslug RoleSlug) ToString() string {
	return string(rslug)
}

// Role represents for role model
type Role struct {
	ID              int       `gorm:"column:id"`
	UUID            string    `gorm:"column:uuid"`
	Slug            RoleSlug  `gorm:"column:slug"`
	Name            string    `gorm:"column:name"`
	Icon            string    `gorm:"column:icon"`
	BackgroundColor string    `gorm:"column:background_color"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
	CreatedAt       time.Time `gorm:"column:created_at"`
}

func GetRoleBySlug(roleList RoleList, roleSlug RoleSlug) (Role, error) {
	for _, role := range roleList.Role {

		if role.Slug == roleSlug {
			return role, nil
		}
	}
	return Role{}, nil
}

// TableName returns table name
func (Role) TableName() string {
	return "roles"
}

// All News response
type RoleList struct {
	Role []Role
}
