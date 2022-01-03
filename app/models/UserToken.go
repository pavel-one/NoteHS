package models

import (
	"app/base"
	"time"
)

type UserToken struct {
	ID        uint
	UserID    uint
	Token     string
	CreatedAt time.Time
}

func (t UserToken) GetUser(token string, db *base.DB) User {
	var user User

	db.Model(&user).Where("id = ?", user.ID).First(&user)

	return user
}
