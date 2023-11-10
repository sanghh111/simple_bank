package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/techschool/simplebank/api/schema"

	db "github.com/techschool/simplebank/db/sqlc"
)

func (server *Server) createAccount(ctx *gin.Context) {
	var req CreateAccountRequest
	var err error
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err, "", time.Now()))
		return
	}

	requestInfoValue, err := req.RequestInfo.ValidateRequestInfo()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err, "", time.Now()))
		return
	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Balance:  req.Balance,
		Currency: "VND",
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err, requestInfoValue.RequestId, requestInfoValue.RequestTime))
		return
	}

	response := schema.GetResponse(account, requestInfoValue.RequestId, requestInfoValue.RequestTime)
	ctx.JSON(http.StatusOK, response)
}

func (server *Server) getAccountById(ctx *gin.Context) {
	requestTime := time.Now()
	requestId := uuid.New().String()
	var req getAccountByIDParam
	err := ctx.ShouldBindUri(&req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err, requestId, requestTime))
		return
	}

	account, err := server.store.GetAccount(ctx, req.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err, requestId, requestTime))
			return
		} else {
			ctx.JSON(http.StatusBadRequest, errResponse(err, requestId, requestTime))
			return
		}
	}

	response := schema.GetResponse(account, requestId, requestTime)
	ctx.JSON(http.StatusOK, response)
}

func (server *Server) getListAccount(ctx *gin.Context) {
	requestTime := time.Now()
	requestId := uuid.New().String()
	var queryParam reqeustListAccountParam
	_ = ctx.ShouldBindQuery(&queryParam)
	if queryParam.PageId <= 0 {
		queryParam.PageId = 1
	}

	if queryParam.PageSize <= 0 {
		queryParam.PageSize = 10
	}

	accounts, err := server.store.ListAccount(ctx, db.ListAccountParams{
		Offset: (queryParam.PageSize * (queryParam.PageId - 1)),
		Limit:  queryParam.PageSize,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err, requestId, requestTime))
		return
	}

	response := schema.GetResponse(accounts, requestId, requestTime)

	ctx.JSON(http.StatusOK, response)
}
