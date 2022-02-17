package resources

import (
	"app/models"
	"gopkg.in/guregu/null.v4"
)

type dialResource struct {
	Id          uint        `json:"id"`
	Name        null.String `json:"name"`
	Description null.String `json:"description"`
	Screen      null.String `json:"screen"`
	Url         string      `json:"url"`
	Final       bool        `json:"final"`
}

func DialResource(dial *models.Dial) *dialResource {
	resource := dialResource{
		Id:          dial.ID,
		Name:        dial.Name,
		Description: dial.Description,
		Screen:      dial.Screen,
		Url:         dial.Url,
		Final:       dial.Final,
	}

	return &resource
}

func DialResources(dials []models.Dial) []*dialResource {
	var out []*dialResource
	for _, dial := range dials {
		out = append(out, DialResource(&dial))
	}

	return out
}
