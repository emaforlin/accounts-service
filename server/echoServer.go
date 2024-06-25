package server

import (
	"fmt"
	"time"

	"github.com/emaforlin/accounts-service/config"
	"github.com/emaforlin/accounts-service/database"
	"github.com/emaforlin/accounts-service/x/handlers"
	"github.com/emaforlin/accounts-service/x/repositories"
	"github.com/emaforlin/accounts-service/x/usecases"
	"github.com/go-playground/validator/v10"
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

	s.app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${method} ${uri} ${status}\n",
	}))
	s.app.Use(middleware.Recover())
	s.app.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{Timeout: 5 * time.Second}))

	s.app.Validator = &handlers.CustomValidator{V: validator.New()}

	serverURL := fmt.Sprintf(":%d", s.cfg.App.Ports["web"])
	s.app.Logger.Fatal(s.app.Start(serverURL))
}

func (s *echoServer) initializeHttpHandlers() {
	// Initialize repositories, usescases, handlers here...
	repository := repositories.NewAccountMysqlRepositoryImpl(s.db)
	usecase := usecases.NewAccountUsecaseImpl(repository)
	httpHandler := handlers.NewAccountHttpHandler(usecase)

	// 	Routers
	router := s.app.Group(s.cfg.App.ApiVersion + "/accounts")
	router.GET("/health", func(c echo.Context) error {
		s.app.Logger.Info("Handle /health")
		return c.String(200, "OK")
	})
	router.POST("/signup", httpHandler.SignupPerson)
}

func NewEchoServer(cfg *config.Config, db database.Database) Server {
	return &echoServer{
		app: echo.New(),
		db:  db,
		cfg: cfg,
	}
}
