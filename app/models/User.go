package models

import (
	"app/base"
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID          uint
	Username    string
	Name        string
	Email       string
	Password    string
	ActivatedAt sql.NullTime `gorm:"type:TIMESTAMP NULL"`
	Tokens      []UserToken  `gorm:"foreignKey:UserID;references:ID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (u *User) Save(db *base.DB) (bool, error) {
	if !u.isUnique(db) {
		return false, errors.New("такой пользователь уже существует")
	}

	hash, err := u.hashPassword()
	u.Password = hash
	if err != nil {
		return false, errors.New("ошибка хэширования пароля")
	}

	u.Tokens = []UserToken{{Token: "test1"}, {Token: "test2"}}

	db.Create(&u)

	return true, nil
}

func (u *User) isUnique(db *base.DB) bool {
	var find int64
	db.Where("email = ?", u.Email).Table("users").Count(&find)

	return !(find > 0)
}

func (u *User) hashPassword() (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)

	return string(bytes), err
}

func (u *User) checkPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
