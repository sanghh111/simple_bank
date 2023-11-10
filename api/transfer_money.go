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

func (server *Server) transferMoney(ctx *gin.Context) {
	var req TransferMoneyRequest
	var err error
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err, "", time.Now()))
		return
	}

	if err := req.TransferInfo.validate(server.store, ctx); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err, "", time.Now()))
		return
	}

	requestInfoValue, err := req.RequestInfo.ValidateRequestInfo()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err, "", time.Now()))
		return
	}

	arg := db.TransferTxParams{
		FromAccountID: req.TransferInfo.FromAccountID,
		ToAccountID:   req.TransferInfo.ToAccountID,
		Amount:        req.TransferInfo.Amount,
	}

	transferResult, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err, requestInfoValue.RequestId, requestInfoValue.RequestTime))
		return
	}

	response := schema.GetResponse(transferResult, requestInfoValue.RequestId, requestInfoValue.RequestTime)
	ctx.JSON(http.StatusOK, response)
}

// validate transferMoneyParam check from_accout and to_account exist
func (tfMoney *transferMoneyParam) validate(sotre db.Store, ctx *gin.Context) error {
	var fromAccount db.Account
	var err error
	if tfMoney.FromAccountID == tfMoney.ToAccountID {
		return errors.New(uti.TransferSameAccount)
	}
	if tfMoney.FromAccountID < tfMoney.ToAccountID {
		fromAccount, err = sotre.GetAccount(ctx, tfMoney.FromAccountID)
		if err != nil {
			if err == sql.ErrNoRows {
				return errors.New(uti.FromAccountExisted)
			} else {
				return err
			}
		}

		_, err = sotre.GetAccount(ctx, tfMoney.ToAccountID)
		if err != nil {
			if err == sql.ErrNoRows {
				return errors.New(uti.ToAccountExisted)
			} else {
				return err
			}
		}
	} else {
		_, err = sotre.GetAccount(ctx, tfMoney.ToAccountID)
		if err != nil {
			if err == sql.ErrNoRows {
				return errors.New(uti.ToAccountExisted)
			} else {
				return err
			}
		}

		fromAccount, err = sotre.GetAccount(ctx, tfMoney.FromAccountID)
		if err != nil {
			if err == sql.ErrNoRows {
				return errors.New(uti.FromAccountExisted)
			} else {
				return err
			}
		}
	}

	if fromAccount.Balance < tfMoney.Amount {
		return errors.New(uti.FromAccountNotEnoughMoney)
	}
	return nil
}
