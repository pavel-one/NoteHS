package requests

import (
	"database/sql"
	"encoding/json"
)

type NullString sql.NullString

type Auth struct {
	Email    NullString `form:"email,string" json:"email,string" binding:"email"`
	Password NullString `form:"password,string" json:"password,string" binding:"min=6"`
	GoogleID NullString `form:"google_id,string" json:"google_id,string" binding:"uuid"`
}

func (ns NullString) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.String)
	}
	return json.Marshal(nil)
}

func (ns *NullString) UnmarshalJSON(data []byte) error {
	var b *string
	if err := json.Unmarshal(data, &b); err != nil {
		return err
	}
	if b != nil {
		ns.Valid = true
		ns.String = *b
	} else {
		ns.Valid = false
	}
	return nil
}
