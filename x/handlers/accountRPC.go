package handlers

import (
	"context"

	protos "github.com/emaforlin/accounts-service/protos/gen/grpc/protos"
	"github.com/emaforlin/accounts-service/x/models"
	"github.com/emaforlin/accounts-service/x/usecases"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type accountServerImpl struct {
	protos.UnimplementedAccountsServer
	log            hclog.Logger
	accountUsecase usecases.AccountUsecase
}

func (h *accountServerImpl) GetAccountDetails(ctx context.Context, ar *protos.AccountDetailsRequest) (*protos.AccountDetailsResponse, error) {
	found, err := h.accountUsecase.GetAccountDetails(&models.GetAccountData{
		Username:    ar.GetUsername(),
		Email:       ar.GetEmail(),
		PhoneNumber: ar.GetPhoneNumber(),
	})

	if err != nil {
		h.log.Error("cannot find user", err)
		return nil, err
	}
	h.log.Info("Get user by username.")

	return &protos.AccountDetailsResponse{
		Id:          found.ID,
		Username:    found.Username,
		Email:       found.Email,
		PhoneNumber: found.PhoneNumber,
		Role:        found.Role,
		CreatedAt:   timestamppb.New(found.CreatedAt),
		UpdatedAt:   timestamppb.New(found.UpdatedAt),
	}, nil
}

func NewAccountGRPCHandler(l hclog.Logger, u usecases.AccountUsecase) *accountServerImpl {
	return &accountServerImpl{
		log:            l,
		accountUsecase: u,
	}
}
