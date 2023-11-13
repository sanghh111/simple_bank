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

type transferMoneyParam struct {
	FromAccountID int64 `json:"from_account_id" binding:"required"`
	ToAccountID   int64 `json:"to_account_id" binding:"required"`
	Amount        int64 `json:"ammount" binding:"required"`
}

type TransferMoneyRequest struct {
	RequestInfo  schema.RequestInfoParam `json:"requestInfo"`
	TransferInfo transferMoneyParam      `json:"transferInfo"`
}

type CreateUserRequest struct {
	RequestInfo schema.RequestInfoParam `json:"requestInfo"`
	Password    string                  `json:"password"`
	Username    string                  `json:"username"`
	Email       string                  `json:"email"`
	Fullname    string                  `json:"fullname"`
}

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
}
