package api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	db "github.com/tetrex/backend-masterclass-go/db/sqlc"
	"github.com/tetrex/backend-masterclass-go/util"
)

type CreateAccountRequest struct {
	Owner    string `json:"owner" binding:"required,owner"`
	Currency string `json:"currency" binding:"required,currency"`
}

// Create Account godoc
//
//	@tags			v1/Account
//	@summary		Creates account Of user
//	@description	takes input of Owner,Currency , and creates account
//	@accept			json
//	@produce		json
//	@param			body body CreateAccountRequest true "CreateAccountRequest"
//	@success		200	{object}	util.OkResponse
//	@failure		500	{object}	util.ErrorResponse
//	@router			/v1/account/ [post]
func (s *Server) createAccount(c echo.Context) error {
	req := new(CreateAccountRequest)
	if err := c.Bind(req); err != nil {

		return c.JSON(http.StatusBadRequest, util.ErrorResponse{Error: err.Error()})
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, util.ErrorResponse{Error: err.Error()})
	}

	args := db.CreateAccountParams{
		Owner:    req.Owner,
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
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (s *Server) getAccount(c echo.Context) error {

	return nil
}

type ListAccountRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (s *Server) listAccounts(c echo.Context) error {

	return nil
}
