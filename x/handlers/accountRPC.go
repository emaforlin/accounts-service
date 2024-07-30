package handlers

import (
	"context"
	"fmt"

	pb "github.com/emaforlin/accounts-service/x/handlers/grpc/protos"
	"github.com/emaforlin/accounts-service/x/models"
	"github.com/emaforlin/accounts-service/x/usecases"
	hclog "github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type accountServerImpl struct {
	pb.UnimplementedAccountsServer
	log     hclog.Logger
	usecase usecases.AccountUsecase
}

func (h *accountServerImpl) LoginUser(ctx context.Context, in *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	jwt, err := h.usecase.Login(in)
	if err != nil {
		fmt.Printf("%v", err)
		return nil, err
	}
	meta := metadata.MD{}
	meta.Append("authorization", "Bearer "+jwt)

	if err := grpc.SetHeader(ctx, meta); err != nil {
		return nil, err
	}
	return &pb.LoginUserResponse{Token: jwt}, nil
}

func (h *accountServerImpl) GetUserId(ctx context.Context, in *pb.GetUserIdRequest) (*pb.GetUserIdResponse, error) {
	h.log.Info("Handle: Get user id")

	input := &models.GetUserId{
		Username:    in.GetUsername(),
		Email:       in.GetEmail(),
		PhoneNumber: in.GetPhoneNumber(),
	}

	h.log.Debug("input: %v", input)

	id, err := h.usecase.GetUserId(input)

	if err != nil {
		return &pb.GetUserIdResponse{
			Userid: -1,
		}, nil
	}

	return &pb.GetUserIdResponse{
		Userid: id,
	}, nil
}

func (h *accountServerImpl) AddFoodPlaceAccount(ctx context.Context, fpr *pb.AddFoodPlaceAccountRequest) (*pb.AddFoodPlaceAccountResponse, error) {
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

	id, _ := h.usecase.GetUserId(&models.GetUserId{Username: input.Username})
	return &pb.AddFoodPlaceAccountResponse{Userid: id}, nil
}

func (h *accountServerImpl) AddPersonAccount(ctx context.Context, pr *pb.AddPersonAccountRequest) (*pb.AddPersonAccountResponse, error) {
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

	id, _ := h.usecase.GetUserId(&models.GetUserId{Username: input.Username})

	return &pb.AddPersonAccountResponse{Userid: id}, nil
}

func (h *accountServerImpl) GetFoodPlaceDetails(ctx context.Context, ar *pb.GetAccountDetailsRequest) (*pb.GetFoodPlaceDetailsResponse, error) {
	h.log.Info("Handle: Get food place")

	input := &models.GetAccountData{
		Id: ar.GetUserid(),
	}

	acc, err := h.usecase.GetAccountDetails(input)
	if err != nil {
		h.log.Error("Cannot find user")
		return nil, err
	}

	fp, err := h.usecase.GetFoodPlaceDetails(input)
	if err != nil {
		h.log.Error("Cannot find user")
		return nil, err
	}

	return &pb.GetFoodPlaceDetailsResponse{
		Id:           acc.ID,
		Username:     acc.Username,
		Email:        acc.Email,
		PhoneNumber:  acc.PhoneNumber,
		Role:         acc.Role,
		BusinessName: fp.BusinessName,
		Location:     fp.Location,
		Tags:         fp.Tags,
		CreatedAt:    timestamppb.New(acc.CreatedAt),
		UpdatedAt:    timestamppb.New(acc.UpdatedAt),
	}, nil
}

func (h *accountServerImpl) GetPersonDetails(ctx context.Context, ar *pb.GetAccountDetailsRequest) (*pb.GetPersonDetailsResponse, error) {
	h.log.Info("Handle: Get User")

	input := &models.GetAccountData{
		Id: ar.GetUserid(),
	}

	acc, err := h.usecase.GetAccountDetails(input)
	if err != nil {
		h.log.Error("Cannot find user")
		return nil, err
	}

	fp, err := h.usecase.GetPersonDetails(input)
	if err != nil {
		h.log.Error("Cannot find user")
		return nil, err
	}

	return &pb.GetPersonDetailsResponse{
		Id:          acc.ID,
		Username:    acc.Username,
		Email:       acc.Email,
		PhoneNumber: acc.PhoneNumber,
		Role:        acc.Role,
		FirstName:   fp.FirstName,
		LastName:    fp.LastName,
		CreatedAt:   timestamppb.New(acc.CreatedAt),
		UpdatedAt:   timestamppb.New(acc.UpdatedAt),
	}, nil
}

func NewAccountGRPCHandler(l hclog.Logger, u usecases.AccountUsecase) pb.AccountsServer {
	return &accountServerImpl{
		log:     l,
		usecase: u,
	}
}
