package api

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/techschool/simplebank/api/schema"
	"github.com/techschool/simplebank/uti"

	db "github.com/techschool/simplebank/db/sqlc"
)

func (server *Server) createAccount(ctx *gin.Context) {
	requestTime := time.Now()
	user, ok := schema.GetUserByBeareToken(ctx, server.jwtMarker, server.store)
	if !ok {
		ctx.JSON(http.StatusForbidden, errResponse(errors.New(uti.Forbidden), "", requestTime))
		return
	}
	var req CreateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err, "", requestTime))
		return
	}

	requestInfoValue, err := req.RequestInfo.ValidateRequestInfo()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err, "", requestTime))
		return
	}

	arg := db.CreateAccountParams{
		Owner:    user,
		Balance:  req.Balance,
		Currency: "VND",
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err, requestInfoValue.RequestId, requestTime))
		return
	}

	response := schema.GetResponse(account, requestInfoValue.RequestId, requestTime)
	ctx.JSON(http.StatusOK, response)
}

func (server *Server) getAccountById(ctx *gin.Context) {
	requestTime := time.Now()
	requestId := uuid.New().String()
	user, ok := schema.GetUserByBeareToken(ctx, server.jwtMarker, server.store)
	if !ok {
		ctx.JSON(http.StatusForbidden, errResponse(errors.New(uti.Forbidden), "", requestTime))
		return
	}
	var req getAccountByIDParam
	err := ctx.ShouldBindUri(&req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err, requestId, requestTime))
		return
	}

	account, err := server.store.GetAccountByOwner(ctx, db.GetAccountByOwnerParams{
		ID:    req.ID,
		Owner: user,
	})

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
	user, ok := schema.GetUserByBeareToken(ctx, server.jwtMarker, server.store)
	if !ok {
		ctx.JSON(http.StatusForbidden, errResponse(errors.New(uti.Forbidden), "", requestTime))
		return
	}
	var queryParam reqeustListAccountParam
	_ = ctx.ShouldBindQuery(&queryParam)
	if queryParam.PageId <= 0 {
		queryParam.PageId = 1
	}

	if queryParam.PageSize <= 0 {
		queryParam.PageSize = 10
	}

	accounts, err := server.store.ListAccount(ctx, db.ListAccountParams{
		Owner:  user,
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
