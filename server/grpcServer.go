package server

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/emaforlin/accounts-service/config"
	"github.com/emaforlin/accounts-service/database"
	"github.com/emaforlin/accounts-service/interceptors"
	"github.com/emaforlin/accounts-service/x/handlers"
	protos "github.com/emaforlin/accounts-service/x/handlers/grpc/protos"
	"github.com/emaforlin/accounts-service/x/repositories"
	"github.com/emaforlin/accounts-service/x/usecases"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type rpcServer struct {
	log hclog.Logger
	cfg *config.Config
	gs  *grpc.Server
	db  database.Database
}

func (r *rpcServer) Start() {
	r.initializeHttpHandlers()

	// try to listen on the specified port
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", r.cfg.App.Port))
	if err != nil {
		r.log.Error(fmt.Sprintf("Unable to listen on %s", l.Addr().String()))
		r.log.Debug(err.Error())
		os.Exit(1)
	}

	r.log.Info(fmt.Sprintf("Listening on %s", l.Addr().String()))

	// channel to handle system signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// goroutine to server the RPC requests
	go func() {
		if err := r.gs.Serve(l); err != nil {
			r.log.Error("Error serving")
			r.log.Debug(err.Error())

		}
	}()

	// wait for interruption signal
	<-sigChan
	r.log.Info("Shutting down server...")

	// gracefully stop the grpc server
	r.gs.GracefulStop()
	r.log.Info("Server stopped gracefully")
}

func (r *rpcServer) initializeHttpHandlers() {
	// create repository
	repository := repositories.NewAccountMysqlRepositoryImpl(r.db)

	// create usecase
	usecase := usecases.NewAccountUsecaseImpl(repository)

	// create handler
	ah := handlers.NewAccountGRPCHandler(r.log, usecase)

	// setup server reflection
	reflection.Register(r.gs)

	// register the server
	protos.RegisterAccountsServer(r.gs, ah)
}

func NewRPCServer(l hclog.Logger, c *config.Config, db database.Database) Server {
	return &rpcServer{
		log: l,
		cfg: c,
		db:  db,
		gs:  grpc.NewServer(grpc.UnaryInterceptor(interceptors.Validation)),
	}
}
