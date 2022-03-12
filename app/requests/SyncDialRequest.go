package requests

type SyncDialRequest struct {
	Dials []interface{} `form:"dials" json:"dials" binding:"required"`
}
