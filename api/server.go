package api

import (
	"github.com/labstack/echo/v4"
	db "github.com/tetrex/backend-masterclass-go/db/sqlc"
)

type Server struct {
	db     *db.Store
	router *echo.Echo
}

func NewServer(store *db.Store) *Server {
	router := echo.New()
	return &Server{
		db:     store,
		router: router,
	}
}

type ServerConfig struct {
	port string
}

func (server *Server) Start(config *ServerConfig) error {
	// v1 group of apis
	v1 := server.router.Group("/v1")

	//account groups
	account := v1.Group("account")

	account.GET("/")
}
