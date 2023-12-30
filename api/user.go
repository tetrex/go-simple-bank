package api

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	db "github.com/tetrex/backend-masterclass-go/db/sqlc"
	"github.com/tetrex/backend-masterclass-go/util"
)

type userResponse struct {
	Username          string    `json:"username"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
}

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
//	@success		200	{object}	util.OkResponse{data=userResponse}
//	@failure		500	{object}	util.ErrorResponse
//	@router			/v1/user [post]
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
	return c.JSON(http.StatusOK, util.OkResponse{Message: "user profile created successfully", Data: newUserResponse(user)})
}

type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}

// Login godoc
//
//	@tags			v1/login
//	@summary		logs in user
//	@description	returns accessToken
//	@accept			json
//	@produce		json
//	@param			body body loginUserRequest true "loginUserRequest"
//	@success		200	{object}	util.OkResponse{data=loginUserResponse}
//	@failure		500	{object}	util.ErrorResponse
//	@router			/v1/login [post]
func (server *Server) loginUser(c echo.Context) error {
	var req loginUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, util.ErrorResponse{Error: err.Error()})
	}

	user, err := server.db.GetUser(context.Background(), req.Username)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ErrorResponse{Error: err.Error()})
	}

	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ErrorResponse{Error: err.Error()})
	}

	accessToken, err := server.tokenMaker.CreateToken(
		user.Username,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ErrorResponse{Error: err.Error()})

	}

	rsp := loginUserResponse{
		AccessToken: accessToken,
		User:        newUserResponse(user),
	}
	return c.JSON(http.StatusOK, util.OkResponse{Message: "user profile created successfully", Data: rsp})
}
