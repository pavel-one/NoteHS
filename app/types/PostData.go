package types

import "encoding/json"

type PostData struct {
	Time    int64                    `json:"time"  binding:"-"`
	Version string                   `json:"version"  binding:"-"`
	Blocks  []map[string]interface{} `json:"blocks"`
}

func (d *PostData) ToString() string {
	m, _ := json.Marshal(d)

	return string(m)
}
