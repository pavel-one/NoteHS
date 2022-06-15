package models

import (
	"app/base"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/guregu/null.v4"
	"time"
)

type User struct {
	ID        uint
	Username  null.String
	Name      null.String
	Email     null.String
	Password  null.String
	GoogleID  null.String
	Tokens    []UserToken `gorm:"foreignKey:UserID;references:ID"`
	Token     UserToken
	Settings  *UserSetting `gorm:"foreignKey:UserId;references:ID"`
	Dials     []Dial
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) Save(db *base.DB) (bool, error) {
	if !u.isUnique(db) {
		return false, errors.New("такой пользователь уже существует")
	}
	password := u.Password.String
	u.Password = null.StringFrom("hashing...")

	u.Settings = &UserSetting{
		Component: "NotePage",
		PostId:    "0",
	}

	db.Create(&u)
	go u.hashPassword(db, password)

	return true, nil
}

func (u *User) CreateToken(db *base.DB) *User {
	token := UserToken{Token: null.StringFrom(u.generateToken())}
	u.Tokens = []UserToken{token}
	db.Save(u)

	return u
}

func (u *User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password.String), []byte(password))
	return err == nil
}

func (u *User) SetActualToken(db *base.DB) *User {
	token := UserToken{}
	db.Model(&token).Order("created_at desc").First(&token)
	u.Token = token
	return u
}

func (u *User) generateToken() string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(u.Email.String), bcrypt.DefaultCost)
	hasher := md5.New()
	hasher.Write(hash)
	return hex.EncodeToString(hasher.Sum(nil))
}

func (u *User) isUnique(db *base.DB) bool {
	var find int64
	db.Where("email = ?", u.Email.String).Table("users").Count(&find)

	return !(find > 0)
}

func (u *User) hashPassword(db *base.DB, password string) {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	u.Password = null.StringFrom(string(bytes))
	db.Save(&u)

	return
}
