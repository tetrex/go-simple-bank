package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	db "github.com/tetrex/backend-masterclass-go/db/sqlc"
	"github.com/tetrex/backend-masterclass-go/util"
)

type Server struct {
	db        *db.Store
	router    *echo.Echo
	validator *validator.Validate
}

func NewServer(store *db.Store) *Server {
	router := echo.New()
	validator := util.NewValidator()

	return &Server{
		db:        store,
		router:    router,
		validator: validator,
	}
}

func (server *Server) Start() {

	//v1 group
	v1 := server.router.Group("/v1")

	//account groups
	account := v1.Group("account")

	//accounts
	account.POST("/", server.createAccount)
	account.GET("/:id", server.getAccount)
	account.GET("/", server.listAccounts)

	// -------------
	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8080),
		Handler: server.router,
	}

	log.Println("Starting server " + s.Addr)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}
