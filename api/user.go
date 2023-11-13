package api

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/techschool/simplebank/api/schema"
	db "github.com/techschool/simplebank/db/sqlc"
	"github.com/techschool/simplebank/uti"
)

func (server *Server) createUser(ctx *gin.Context) {
	var req CreateUserRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err, "", time.Now()))
		return
	}
	requestValue, err := req.RequestInfo.ValidateRequestInfo()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err, "", time.Now()))
		return
	}

	err = (&req).valid(server.store, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err, requestValue.RequestId, requestValue.RequestTime))
		return
	}

	createUserParams := db.CreateUserParams{
		Username:     req.Username,
		HashPassword: req.Password,
		FullName:     req.Fullname,
		Email:        req.Email,
	}
	user, err := server.store.CreateUser(ctx, createUserParams)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err, requestValue.RequestId, requestValue.RequestTime))
		return
	}
	userResponse := UserResponse{
		Username: user.Username,
		Email:    user.Email,
		Fullname: user.FullName,
	}

	response := schema.GetResponse(userResponse, requestValue.RequestId, requestValue.RequestTime)
	ctx.JSON(http.StatusOK, response)
}

func (req *CreateUserRequest) valid(store db.Store, ctx *gin.Context) error {
	if req.Username == "" {
		return errors.New(uti.UsernameEmpty)
	}
	if req.Password == "" {
		return errors.New(uti.PasswordEmpty)
	}
	if req.Fullname == "" {
		return errors.New(uti.FullNameEmpty)
	}
	if req.Email == "" {
		return errors.New(uti.EmailEmpty)
	}
	_, err := store.GetUser(ctx, req.Username)
	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}
	} else {
		return errors.New(uti.UserNameExisted)
	}

	_, err = store.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}
	} else {
		return errors.New(uti.EmailExisted)
	}
	req.Password, err = uti.HashPassword(req.Password)
	if err != nil {
		return err
	}
	return nil
}

func (server *Server) login(ctx *gin.Context) {
	now := time.Now()
	username, password, ok := ctx.Request.BasicAuth()
	if !ok {
		ctx.JSON(http.StatusForbidden, errResponse(
			errors.New("Forbidden"),
			"",
			now))
		return
	}

	user, err := server.store.GetUser(ctx, username)
	if err != nil {
		ctx.JSON(http.StatusForbidden, errResponse(
			errors.New("Forbidden"),
			"",
			now))
		return
	}
	ok = uti.CheckPassword(password, user.HashPassword)
	if !ok {
		ctx.JSON(http.StatusForbidden, errResponse(
			errors.New("Forbidden"),
			"",
			now))
		return
	}
	userResponse := UserResponse{
		Username: user.Username,
		Email:    user.Email,
		Fullname: user.FullName,
	}
	response := schema.GetResponse(userResponse, "", now)
	ctx.JSON(http.StatusOK, response)
}
