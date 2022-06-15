package models

import (
	"app/base"
	"gopkg.in/guregu/null.v4"
	"time"
)

type UserToken struct {
	ID        uint
	UserID    uint
	Token     null.String
	LatestUse time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *UserToken) GetUser(db *base.DB) User {
	var user User

	db.Model(&user).Where("id = ?", t.UserID).Preload("Settings").First(&user)

	return user
}

func (t *UserToken) UpdateUse(db *base.DB) {
	t.LatestUse = time.Now()
	db.Save(&t)
}
