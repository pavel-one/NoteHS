package models

import (
	"app/Services/Scrapper"
	"app/base"
	"app/requests"
	"gopkg.in/guregu/null.v4"
	"log"
	"os"
	"strconv"
	"time"
)

type Dial struct {
	ID          uint
	UserID      uint
	Name        null.String
	Description null.String
	Screen      null.String
	Url         string
	Final       bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (dial *Dial) FillWithRequest(db *base.DB, request requests.CreateDialRequest) {
	dial.Url = request.Url
	dial.Description = request.Description
	dial.Name = request.Description
}

func (dial *Dial) CreateOrUpdateInfo(db *base.DB) {

	defer dial.SetProcessEnd(db)

	url, err := Scrapper.GetUrlInfo(dial.Url, strconv.Itoa(int(dial.ID)))
	if err != nil {
		log.Println(err)
		return
	}

	if !dial.Name.Valid {
		dial.Name = null.StringFrom(url.Title)
	}

	if !dial.Screen.Valid {
		dial.Screen = null.StringFrom(url.Screen)
	}
}

func (dial *Dial) SetProcess(db *base.DB) {
	if dial.Final == false {
		return
	}

	dial.Final = false
	db.Save(dial)
}

func (dial *Dial) SetProcessEnd(db *base.DB) {
	if dial.Final == true {
		return
	}

	dial.Final = true
	db.Save(dial)
}

func (dial *Dial) DropDialWithFiles(db *base.DB) {
	if dial.Screen.Valid {
		_ = os.Remove(dial.Screen.String)
	}

	db.Delete(dial)
}
