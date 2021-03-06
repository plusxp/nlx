package database

import "time"

type Role struct {
	ID          uint `gorm:"primarykey;column:role_id;"`
	Name        string
	Code        string
	Permissions []Permission `gorm:"-"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (s *Role) TableName() string {
	return "nlx_management.roles"
}
