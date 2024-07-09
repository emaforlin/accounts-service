package handlers

import (
	"context"

	protos "github.com/emaforlin/accounts-service/x/handlers/grpc/protos"
	"github.com/emaforlin/accounts-service/x/models"
	"github.com/emaforlin/accounts-service/x/usecases"
	"github.com/go-playground/validator/v10"
	hclog "github.com/hashicorp/go-hclog"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type accountServerImpl struct {
	protos.UnimplementedAccountsServer
	log      hclog.Logger
	usecase  usecases.AccountUsecase
	validate *validator.Validate
}

func (h *accountServerImpl) AddPersonAccount(ctx context.Context, pr *protos.AddPersonAccountRequest) (*protos.GenericResponse, error) {
	input := &models.AddPersonAccountData{
		Username:    pr.GetUsername(),
		FirstName:   pr.GetFirstName(),
		LastName:    pr.GetLastName(),
		PhoneNumber: pr.GetPhoneNumber(),
		Email:       pr.GetEmail(),
		Password:    pr.GetPassword(),
	}
	// validate fields
	if err := h.validate.Struct(input); err != nil {
		h.log.Error("invalid input data", err.Error())
		return nil, err
	}

	if err := h.usecase.AddPersonAccount(input); err != nil {
		return nil, err
	}

	h.log.Info("Handle Create person account")
	return &protos.GenericResponse{
		Message: "Account successfully created",
	}, nil
}

func (h *accountServerImpl) GetAccountDetails(ctx context.Context, ar *protos.GetAccountDetailsRequest) (*protos.GetAccountDetailsResponse, error) {
	found, err := h.usecase.GetAccountDetails(&models.GetAccountData{
		Username:    ar.GetUsername(),
		Email:       ar.GetEmail(),
		PhoneNumber: ar.GetPhoneNumber(),
	})

	if err != nil {
		h.log.Error("cannot find user", err)
		return nil, err
	}
	h.log.Info("Handle Get User")

	return &protos.GetAccountDetailsResponse{
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
		log:      l,
		usecase:  u,
		validate: validator.New(),
	}
}
