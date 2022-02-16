package resources

import "app/models"

type dialResource struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Screen      string `json:"screen"`
	Url         string `json:"url"`
	Final       bool   `json:"final"`
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
