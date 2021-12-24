package models

import (
	"time"
)

type UserToken struct {
	ID        uint
	UserID    uint
	Token     string
	CreatedAt time.Time
}
