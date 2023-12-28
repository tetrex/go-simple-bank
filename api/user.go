package api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	db "github.com/tetrex/backend-masterclass-go/db/sqlc"
	"github.com/tetrex/backend-masterclass-go/util"
)

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

// Create User godoc
//
//	@tags			v1/User
//	@summary		Creates user profile
//	@description	returns user newly created user profile
//	@accept			json
//	@produce		json
//	@param			body body CreateUserRequest true "CreateUserRequest"
//	@success		200	{object}	util.OkResponse
//	@failure		500	{object}	util.ErrorResponse
//	@router			/v1/user/ [post]
func (s *Server) createUser(c echo.Context) error {
	req := new(CreateUserRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, util.ErrorResponse{Error: err.Error()})
	}

	if err := s.validator.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, util.ErrorResponse{Error: err.Error()})
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ErrorResponse{Error: err.Error()})
	}

	args := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
	}

	user, err := s.db.CreateUser(context.Background(), args)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ErrorResponse{Error: err.Error()})
	}
	return c.JSON(http.StatusOK, util.OkResponse{Message: "user profile created successfully", Data: user})
}
