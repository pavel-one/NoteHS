package resources

import "app/models"

type SettingResource struct {
	Component *string `json:"component"`
	PostId    *string `json:"post"`
}

func GetSettingResource(setting *models.UserSetting) SettingResource {
	resource := SettingResource{
		Component: &setting.Component,
		PostId:    &setting.PostId,
	}

	return resource
}
