package server

import (
	"fmt"
	"time"

	"github.com/emaforlin/accounts-service/config"
	"github.com/emaforlin/accounts-service/database"
	"github.com/emaforlin/accounts-service/x/handlers"
	"github.com/emaforlin/accounts-service/x/repositories"
	"github.com/emaforlin/accounts-service/x/usecases"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type echoServer struct {
	app *echo.Echo
	db  database.Database
	cfg *config.Config
}

func (s *echoServer) Start() {
	s.initializeHttpHandlers()

	s.app.Use(middleware.Logger())
	s.app.Use(middleware.Recover())
	s.app.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{Timeout: 5 * time.Second}))

	serverURL := fmt.Sprintf(":%d", s.cfg.App.Ports["web"])
	s.app.Logger.Fatal(s.app.Start(serverURL))
}

func (s *echoServer) initializeHttpHandlers() {
	// Initialize repositories, usescases, handlers here...
	repository := repositories.NewAccountMysqlRepositoryImpl(s.db)
	usecase := usecases.NewAccountUsecaseImpl(repository)
	httpHandler := handlers.NewAccountHttpHandler(usecase)

	// 	Routers
	s.app.GET(s.cfg.App.ApiVersion+"/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	router := s.app.Group(s.cfg.App.ApiVersion + "/accounts")
	router.POST("/signup", httpHandler.SignupPerson)
}

func NewEchoServer(cfg *config.Config, db database.Database) Server {
	return &echoServer{
		app: echo.New(),
		db:  db,
		cfg: cfg,
	}
}
