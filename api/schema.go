package api

import "github.com/techschool/simplebank/api/schema"

type CreateAccountRequest struct {
	RequestInfo schema.RequestInfoParam `json:"requestInfo"`
	Balance     int64                   `json:"balance" binding:"required"`
	Owner       string                  `json:"owner" binding:"required"`
}

type reqeustListAccountParam struct {
	PageId   int64 `form:"page_id"`
	PageSize int64 `form:"page_size"`
}

type getAccountByIDParam struct {
	ID int64 `uri:"id" binding:"required"`
}
