package models

import (
	"time"
)

type User struct {
	ID          uint
	Username    string
	Name        string
	Email       string
	Password    string
	ActivatedAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
