package server

import (
	"fmt"
	"net"
	"os"

	"github.com/emaforlin/accounts-service/config"
	"github.com/emaforlin/accounts-service/database"
	"github.com/emaforlin/accounts-service/x/handlers"
	protos "github.com/emaforlin/accounts-service/x/handlers/grpc"
	"github.com/emaforlin/accounts-service/x/repositories"
	"github.com/emaforlin/accounts-service/x/usecases"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

type rpcServer struct {
	log hclog.Logger
	cfg *config.Config
	gs  *grpc.Server
	db  database.Database
}

func (r *rpcServer) Start() {
	r.initializeHttpHandlers()

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", r.cfg.App.Ports["rpc"]))

	if err != nil {
		r.log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}
	r.log.Info("Listening...", r.gs.Serve(l))
}

func (r *rpcServer) initializeHttpHandlers() {
	repository := repositories.NewAccountMysqlRepositoryImpl(r.db)
	usecase := usecases.NewAccountUsecaseImpl(repository)

	ah := handlers.NewAccountGRPCHandler(r.log, usecase)

	protos.RegisterAccountsServer(r.gs, ah)
}

func NewRPCServer(l hclog.Logger, c *config.Config, db database.Database) Server {
	return &rpcServer{
		log: l,
		cfg: c,
		db:  db,
		gs:  grpc.NewServer(),
	}
}
