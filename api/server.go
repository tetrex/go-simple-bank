package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger" // middlerware for echo-swagger
	db "github.com/tetrex/backend-masterclass-go/db/sqlc"
	_ "github.com/tetrex/backend-masterclass-go/docs" // docs is generated by Swag CLI
	"github.com/tetrex/backend-masterclass-go/util"
)

type Server struct {
	db        *db.Store
	router    *echo.Echo
	validator *validator.Validate
}

func NewServer(s *db.Store) *Server {
	r := echo.New()
	v := util.NewValidator()

	server := &Server{
		db:        s,
		router:    r,
		validator: v,
	}
	return server
}

func (server *Server) Start() {
	config, err := util.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// health check api
	server.router.GET("/", server.health)

	//v1 group
	v1 := server.router.Group("v1/")

	// swagger docs
	v1.GET("docs/*", echoSwagger.WrapHandler)

	//accounts
	v1.POST("account", server.createAccount)
	v1.GET("account/:id", server.getAccount)
	v1.POST("accounts", server.listAccounts)

	// -------------

	log.Printf("Starting server :: %d", config.ServerPort)
	if err := server.router.Start(fmt.Sprintf(":%d", config.ServerPort)); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}

// Health godoc
//
//	@summary		For health check, of server
//	@description	Gives us Server Time , To check health of server
//	@tags			health
//	@accept			json
//	@produce		json
//	@success		200	{object}	util.OkResponse
//	@failure		500	{object}	error
//	@router			/ [get]
func (s *Server) health(c echo.Context) error {
	return c.JSON(http.StatusOK, util.OkResponse{Message: "server,all good", Data: time.Now().UnixNano()})
}
