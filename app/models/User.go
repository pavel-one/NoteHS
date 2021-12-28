package models

import (
	"app/base"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
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
	Token       UserToken
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (u *User) Save(db *base.DB) (bool, error) {
	if !u.isUnique(db) {
		return false, errors.New("такой пользователь уже существует")
	}
	password := u.Password
	u.Password = "hashing..."

	db.Create(&u)
	u.CreateToken(db)
	u.setActualToken(db)
	go u.hashPassword(db, password)

	return true, nil
}

func (u *User) CreateToken(db *base.DB) {
	token := UserToken{Token: u.generateToken()}
	u.Tokens = []UserToken{token}
	db.Save(&u)
}

func (u *User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) setActualToken(db *base.DB) {
	token := UserToken{}
	ttl, _ := strconv.Atoi(os.Getenv("TOKEN_TTL_DAY"))
	dateWithTtl := time.Now().AddDate(0, 0, -ttl).Format("2006-01-02 15:04:05")

	//TODO: Неправильная выборка актуального токена
	db.Model(&token).Order("created_at desc").Where("created_at >= ? and user_id = ?", dateWithTtl, u.ID).First(&token)

	u.Token = token
}

func (u *User) generateToken() string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(u.Email), bcrypt.DefaultCost)
	hasher := md5.New()
	hasher.Write(hash)
	return hex.EncodeToString(hasher.Sum(nil))
}

func (u *User) isUnique(db *base.DB) bool {
	var find int64
	db.Where("email = ?", u.Email).Table("users").Count(&find)

	return !(find > 0)
}

func (u *User) hashPassword(db *base.DB, password string) {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	u.Password = string(bytes)
	db.Save(&u)

	return
}
