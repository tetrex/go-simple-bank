package api

import (
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	db "github.com/tetrex/backend-masterclass-go/db/sqlc"
	"github.com/tetrex/backend-masterclass-go/token"
	"github.com/tetrex/backend-masterclass-go/util"
)

type CreateAccountRequest struct {
	Currency string `json:"currency" validate:"required,currency"`
}

// Create Account godoc
//
//	@tags			v1/Account
//	@summary		Creates account Of user
//	@description	takes input of Owner,Currency , and creates account
//
// @Security ApiKeyAuth
//
//	@accept			json
//	@produce		json
//	@param			body body CreateAccountRequest true "CreateAccountRequest"
//	@success		200	{object}	util.OkResponse{data=db.Account}
//	@failure		500	{object}	util.ErrorResponse
//	@router			/v1/account [post]
func (s *Server) createAccount(c echo.Context) error {
	req := new(CreateAccountRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, util.ErrorResponse{Error: err.Error()})
	}

	if err := s.validator.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, util.ErrorResponse{Error: err.Error()})
	}

	authPayload := c.Get(authorizationPayloadKey).(*token.Payload)

	args := db.CreateAccountParams{
		Owner:    authPayload.Username,
		Currency: req.Currency,
		Balance:  0,
	}
	account, err := s.db.CreateAccount(context.Background(), args)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ErrorResponse{Error: err.Error()})
	}
	return c.JSON(http.StatusOK, util.OkResponse{Message: "account created successfully", Data: account})
}

type GetAccountRequest struct {
	ID int64 `param:"id" validate:"required,min=1"`
}

// Create Account godoc
//
//	@tags			v1/Account
//	@summary		Gets User Account
//	@description	takes id of user and returns user account
//
// @Security ApiKeyAuth
//
//	@accept			json
//	@produce		json
//	@param			id	path		int	true	"User ID"
//	@success		200	{object}	util.OkResponse{data=db.Account}
//	@failure		500	{object}	util.ErrorResponse
//	@router			/v1/account/{id} [get]
func (s *Server) getAccount(c echo.Context) error {
	req := new(GetAccountRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, util.ErrorResponse{Error: err.Error()})
	}

	if err := s.validator.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, util.ErrorResponse{Error: err.Error()})
	}

	account, err := s.db.GetAccount(context.Background(), req.ID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ErrorResponse{Error: err.Error()})
	}

	authPayload := c.Get(authorizationPayloadKey).(*token.Payload)
	if account.Owner != authPayload.Username {
		err := errors.New("account doesn't belong to the authenticated user")
		return c.JSON(http.StatusUnauthorized, util.ErrorResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, util.OkResponse{Message: "account", Data: account})
}

type ListAccountRequest struct {
	PageID   int32 `json:"page_id" validate:"required,min=1"`
	PageSize int32 `json:"page_size" validate:"required,min=5,max=10"`
}

// Create Account godoc
//
//	@tags			v1/Account
//	@summary		Gets List Of User Account
//	@description	takes pages and pagesize
//
// @Security ApiKeyAuth
//
//	@accept			json
//	@produce		json
//	@param			body body ListAccountRequest true "ListAccountRequest"
//	@success		200	{object}	util.OkResponse{data=[]db.Account}
//	@failure		500	{object}	util.ErrorResponse
//	@router			/v1/accounts [post]
func (s *Server) listAccounts(c echo.Context) error {
	req := new(ListAccountRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, util.ErrorResponse{Error: err.Error()})
	}

	if err := s.validator.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, util.ErrorResponse{Error: err.Error()})
	}
	authPayload := c.Get(authorizationPayloadKey).(*token.Payload)

	args := db.ListAccountsParams{
		Owner:  authPayload.Username,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	account, err := s.db.ListAccounts(context.Background(), args)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ErrorResponse{Error: err.Error()})
	}

	return c.JSON(http.StatusOK, util.OkResponse{Message: "accounts", Data: account})
}
