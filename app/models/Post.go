package models

import (
	"gopkg.in/guregu/null.v4"
	"time"
)

type Post struct {
	Uuid        string
	PostData    string
	UserId      uint
	Name        string
	Description null.String
	Public      bool
	UpdatedAt   time.Time
}
