package models

import "time"

type Dial struct {
	ID          uint
	UserID      uint
	Name        string
	Description string
	Screen      string
	Url         string
	Final       bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
