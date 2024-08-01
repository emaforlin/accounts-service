package interceptors

import (
	"context"
	"slices"
	"strings"

	"github.com/emaforlin/accounts-service/auth"
	"github.com/emaforlin/accounts-service/config"
	"github.com/emaforlin/accounts-service/x/entities"
	"github.com/emaforlin/accounts-service/x/usecases"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var cfg *config.Config

func JWTAuth(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	cfg = config.LoadConfig()
	jwtFactory := auth.NewJWTFactory(cfg.Jwt)
	usecase := usecases.NewAuthUsecaseImpl(jwtFactory)

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "missing metadata")
	}

	userData := usecase.Authorize(md["authorization"])
	if userData == nil {
		userData = &entities.UserData{
			Role: "Visitor",
		}
	}
	if !checkAccess(strings.ToLower(userData.Role), info.FullMethod) {
		return nil, status.Error(codes.Unauthenticated, "user need permission")
	}

	return handler(ctx, req)
}

func checkAccess(role, method string) bool {
	return slices.Index(cfg.AccessControl[role], method) != -1
}
