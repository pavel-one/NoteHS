package requests

type SetSettingRequest struct {
	*BaseRequest
	Component string `form:"component" json:"component" binding:"omitempty,max=255"`
	PostId    string `form:"post_id" json:"post_id" binding:"omitempty,max=255"`
}
