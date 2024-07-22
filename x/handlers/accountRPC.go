package handlers

import (
	"context"

	"github.com/emaforlin/accounts-service/x/handlers/grpc/protos"
	"github.com/emaforlin/accounts-service/x/models"
	"github.com/emaforlin/accounts-service/x/usecases"
	hclog "github.com/hashicorp/go-hclog"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type accountServerImpl struct {
	protos.UnimplementedAccountsServer
	log     hclog.Logger
	usecase usecases.AccountUsecase
}

func (h *accountServerImpl) GetUserId(ctx context.Context, in *protos.GetUserIdRequest) (*protos.GetUserIdResponse, error) {
	h.log.Info("Handle: Get user id")

	input := &models.GetUserId{
		Username:    in.GetUsername(),
		Email:       in.GetEmail(),
		PhoneNumber: in.GetPhoneNumber(),
	}

	h.log.Debug("input: %v", input)

	id, err := h.usecase.GetUserId(input)

	if err != nil {
		return &protos.GetUserIdResponse{
			Userid: -10,
		}, nil
	}

	return &protos.GetUserIdResponse{
		Userid: id,
	}, nil
}

func (h *accountServerImpl) VerifyFoodPlaceAccount(ctx context.Context, fp *protos.VerifyFoodPlaceAccountRequest) (*protos.VerifyFoodPlaceAccountResponse, error) {
	h.log.Info("Handle: Verify food place account ")

	input := &models.VerifyFoodPlaceAccount{
		UserId: fp.GetUserid(),
	}

	if err := h.usecase.VerifyFoodPlaceAccount(input); err != nil {
		h.log.Error("Account could not be verified")
		return nil, err
	}

	return &protos.VerifyFoodPlaceAccountResponse{}, nil
}

func (h *accountServerImpl) AddFoodPlaceAccount(ctx context.Context, fpr *protos.AddFoodPlaceAccountRequest) (*protos.AddFoodPlaceAccountResponse, error) {
	h.log.Info("Handle: Create food place account")

	input := &models.AddFoodPlaceAccountData{
		Username:     fpr.GetUsername(),
		PhoneNumber:  fpr.GetPhoneNumber(),
		Email:        fpr.GetEmail(),
		Password:     fpr.GetPassword(),
		BusinessName: fpr.GetBusinessName(),
		Location:     fpr.GetLocation(),
		Tags:         fpr.GetTags(),
	}

	if err := h.usecase.AddFoodPlaceAccount(input); err != nil {
		h.log.Error("Error creating account")
		return nil, err
	}
	return &protos.AddFoodPlaceAccountResponse{}, nil
}

func (h *accountServerImpl) AddPersonAccount(ctx context.Context, pr *protos.AddPersonAccountRequest) (*protos.AddPersonAccountResponse, error) {
	h.log.Info("Handle: Create person account")

	input := &models.AddPersonAccountData{
		Username:    pr.GetUsername(),
		FirstName:   pr.GetFirstName(),
		LastName:    pr.GetLastName(),
		PhoneNumber: pr.GetPhoneNumber(),
		Email:       pr.GetEmail(),
		Password:    pr.GetPassword(),
	}

	if err := h.usecase.AddPersonAccount(input); err != nil {
		h.log.Error("Error creating account")
		return nil, err
	}

	return &protos.AddPersonAccountResponse{}, nil
}

func (h *accountServerImpl) GetAccountDetails(ctx context.Context, ar *protos.GetAccountDetailsRequest) (*protos.GetAccountDetailsResponse, error) {
	h.log.Info("Handle: Get User")

	input := &models.GetAccountData{
		Id: ar.GetUserid(),
	}

	found, err := h.usecase.GetAccountDetails(input)
	if err != nil {
		h.log.Error("Cannot find user")
		return nil, err
	}

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
		log:     l,
		usecase: u,
	}
}
