package requests

import "gopkg.in/guregu/null.v4"

type PostRequest struct {
	*BaseRequest
	Id          null.String `json:"id" binding:"omitempty,uuid"`
	Name        string      `json:"name" binding:"required,max=255"`
	Description null.String `json:"description" binding:"omitempty,max=255"`
	Data        []struct {
		Test  string `json:"test"`
		Test1 []struct {
			Test2 string `json:"test2"`
		} `json:"test1"`
	} `json:"data"`
}
