package models

import (
	"app/Services/Scrapper"
	"app/base"
	"app/requests"
	"log"
	"strconv"
	"time"
)

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

	if dial.Name == "" {
		dial.Name = url.Title
	}

	if dial.Screen == "" {
		dial.Screen = url.Screen
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
