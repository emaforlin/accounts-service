package server

import (
	"fmt"

	"github.com/emaforlin/accounts-service/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type echoServer struct {
	app *echo.Echo
	db  *gorm.DB
	cfg *config.Config
}

func (s *echoServer) Start() {
	s.initializeHttpHandlers()

	s.app.Use(middleware.Logger())

	serverURL := fmt.Sprintf(":%d", s.cfg.App.Port)
	s.app.Logger.Fatal(s.app.Start(serverURL))
}

func (s *echoServer) initializeHttpHandlers() {
	// Initialize repositories, usescases, handlers here...
	// httpHandler := handlers.NewUserHttpHandler(usecase)

	// 	Routers
	_ = s.app.Group("/accounts/" + s.cfg.App.ApiVersion)
}

func NewEchoServer(db *gorm.DB, cfg *config.Config) Server {
	return &echoServer{
		app: echo.New(),
		db:  db,
		cfg: cfg,
	}
}
