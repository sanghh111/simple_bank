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
	token, ok := schema.ParseBearerAuth(ctx)
	if !ok {
		ctx.JSON(http.StatusForbidden, errResponse(errors.New(uti.Forbidden), "", requestTime))
		return
	}
	payload, err := server.jwtMarker.VerifyToken(token)
	if err != nil {
		ctx.JSON(http.StatusForbidden, errResponse(err, "", requestTime))
		return
	}
	user, _ := server.store.GetUser(ctx, payload.Username)
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
		Owner:    user.Username,
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
	token, ok := schema.ParseBearerAuth(ctx)
	if !ok {
		ctx.JSON(http.StatusForbidden, errResponse(errors.New(uti.Forbidden), requestId, requestTime))
		return
	}
	payload, err := server.jwtMarker.VerifyToken(token)
	if err != nil {
		ctx.JSON(http.StatusForbidden, errResponse(err, requestId, requestTime))
		return
	}
	user, _ := server.store.GetUser(ctx, payload.Username)
	var req getAccountByIDParam
	err = ctx.ShouldBindUri(&req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err, requestId, requestTime))
		return
	}

	account, err := server.store.GetAccountByOwner(ctx, db.GetAccountByOwnerParams{
		ID:    req.ID,
		Owner: user.Username,
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
	token, ok := schema.ParseBearerAuth(ctx)
	if !ok {
		ctx.JSON(http.StatusForbidden, errResponse(errors.New(uti.Forbidden), requestId, requestTime))
		return
	}
	payload, err := server.jwtMarker.VerifyToken(token)
	if err != nil {
		ctx.JSON(http.StatusForbidden, errResponse(err, requestId, requestTime))
		return
	}
	user, _ := server.store.GetUser(ctx, payload.Username)
	var queryParam reqeustListAccountParam
	_ = ctx.ShouldBindQuery(&queryParam)
	if queryParam.PageId <= 0 {
		queryParam.PageId = 1
	}

	if queryParam.PageSize <= 0 {
		queryParam.PageSize = 10
	}

	accounts, err := server.store.ListAccount(ctx, db.ListAccountParams{
		Owner:  user.Username,
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
