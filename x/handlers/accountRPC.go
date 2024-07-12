package handlers

import (
	"context"

	protos "github.com/emaforlin/accounts-service/x/handlers/grpc/protos"
	"github.com/emaforlin/accounts-service/x/models"
	"github.com/emaforlin/accounts-service/x/usecases"
	"github.com/go-playground/validator/v10"
	hclog "github.com/hashicorp/go-hclog"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type accountServerImpl struct {
	protos.UnimplementedAccountsServer
	log      hclog.Logger
	usecase  usecases.AccountUsecase
	validate *validator.Validate
}

func (h *accountServerImpl) VerifyFoodPlaceAccount(ctx context.Context, fp *protos.VerifyFoodPlaceAccountRequest) (*protos.VerifyFoodPlaceAccountResponse, error) {
	input := models.VerifyFoodPlaceAccount{
		UserId: fp.GetUserid(),
	}
	// validate fields
	if err := h.validate.Struct(input); err != nil {
		h.log.Error("invalid input data", err.Error())
		return nil, err
	}

	h.log.Info("Handle: Verify food place account ")

	if err := h.usecase.VerifyFoodPlaceAccount(&input); err != nil {
		return nil, err
	}

	return &protos.VerifyFoodPlaceAccountResponse{}, nil
}

func (h *accountServerImpl) AddFoodPlaceAccount(ctx context.Context, fpr *protos.AddFoodPlaceAccountRequest) (*emptypb.Empty, error) {
	input := &models.AddFoodPlaceAccountData{
		Username:     fpr.GetUsername(),
		PhoneNumber:  fpr.GetPhoneNumber(),
		Email:        fpr.GetEmail(),
		Password:     fpr.GetPassword(),
		BusinessName: fpr.GetBusinessName(),
		Location:     fpr.GetLocation(),
		Tags:         fpr.GetTags(),
	}
	// validate fields
	if err := h.validate.Struct(input); err != nil {
		h.log.Error("invalid input data", err.Error())
		return nil, err
	}

	h.log.Info("Handle: Create food place account")
	if err := h.usecase.AddFoodPlaceAccount(input); err != nil {
		h.log.Error("error creating account")
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (h *accountServerImpl) AddPersonAccount(ctx context.Context, pr *protos.AddPersonAccountRequest) (*emptypb.Empty, error) {
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

	h.log.Info("Handle: Create person account")
	if err := h.usecase.AddPersonAccount(input); err != nil {
		h.log.Error("error creating account", err.Error())
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (h *accountServerImpl) GetAccountDetails(ctx context.Context, ar *protos.GetAccountDetailsRequest) (*protos.GetAccountDetailsResponse, error) {
	input := &models.GetAccountData{
		Username:    ar.GetUsername(),
		Email:       ar.GetEmail(),
		PhoneNumber: ar.GetPhoneNumber(),
	}

	// validate fields
	if err := h.validate.Struct(input); err != nil {
		h.log.Error("invalid input data", err.Error())
		return nil, err
	}

	found, err := h.usecase.GetAccountDetails(input)
	if err != nil {
		h.log.Error("cannot find user", err)
		return nil, err
	}
	h.log.Info("Handle: Get User")

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

func NewAccountGRPCHandler(l hclog.Logger, u usecases.AccountUsecase) protos.AccountsServer {
	return &accountServerImpl{
		log:      l,
		usecase:  u,
		validate: validator.New(),
	}
}
