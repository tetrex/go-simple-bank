package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	db "github.com/tetrex/backend-masterclass-go/db/sqlc"
	"github.com/tetrex/backend-masterclass-go/token"
	"github.com/tetrex/backend-masterclass-go/util"
)

type TransferRequest struct {
	FromAccountId int64  `json:"from_account_id" validate:"required,min=1"`
	ToAccountId   int64  `json:"to_account_id" validate:"required,min=1"`
	Amount        int64  `json:"amount" validate:"required,gt=0"`
	Currency      string `json:"currency" validate:"required,oneof=USD EUR CAD"`
}

// Create Transfer Money godoc
//
//	@tags			v1/TransferMoney
//	@summary		Transfer's money from Acc1 to Acc2
//	@description	takes input and transfers money from -> to
//
// @Security ApiKeyAuth
//
//	@accept			json
//	@produce		json
//	@param			body body TransferRequest true "TransferRequest"
//	@success		200	{object}	util.OkResponse{data=db.TransferTxResult}
//	@failure		500	{object}	util.ErrorResponse
//	@router			/v1/transfer [post]
func (s *Server) createTransfer(c echo.Context) error {
	req := new(TransferRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, util.ErrorResponse{Error: err.Error()})
	}

	if err := s.validator.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, util.ErrorResponse{Error: err.Error()})
	}

	fromAccount, err := s.validAccount(req.FromAccountId, req.Currency)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ErrorResponse{Error: err.Error()})
	}

	authPayload := c.Get(authorizationPayloadKey).(*token.Payload)
	if fromAccount.Owner != authPayload.Username {
		err := errors.New("from account doesn't belong to the authenticated user")
		return c.JSON(http.StatusBadRequest, util.ErrorResponse{Error: err.Error()})
	}

	if _, err := s.validAccount(req.ToAccountId, req.Currency); err != nil {
		return c.JSON(http.StatusBadRequest, util.ErrorResponse{Error: err.Error()})
	}

	args := db.TransferTxParams{
		FromAccountId: req.FromAccountId,
		ToAccountId:   req.ToAccountId,
		Amount:        req.Amount,
	}
	result, err := s.db.TransferTx(context.Background(), args)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ErrorResponse{Error: err.Error()})
	}
	return c.JSON(http.StatusOK, util.OkResponse{Message: "transfer successful", Data: result})
}

func (s *Server) validAccount(accountID int64, currency string) (db.Account, error) {
	account, err := s.db.GetAccount(context.Background(), accountID)
	if err != nil {
		return account, err
	}

	if account.Currency != currency {
		err := fmt.Errorf("account [%d] currency mismatch: %s vs %s", account.ID, account.Currency, currency)
		return account, err
	}

	return account, err
}
