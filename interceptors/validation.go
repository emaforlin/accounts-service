package interceptors

import (
	"context"
	"errors"

	protos "github.com/emaforlin/accounts-service/x/handlers/grpc/protos"
	"github.com/emaforlin/accounts-service/x/models"
	validator "github.com/go-playground/validator/v10"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var validate *validator.Validate

func ValidationUnary(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	validate = validator.New()

	mappedReq, err := mapReq(req)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err := validate.Struct(mappedReq); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil, status.Errorf(codes.InvalidArgument, "internal validation error: %v", err)
		}
		return nil, status.Errorf(codes.InvalidArgument, "validation error: %v", err)
	}

	return handler(ctx, req)
}

func mapReq(req any) (any, error) {
	switch r := req.(type) {
	case *protos.LoginUserRequest:
		mappedStruct := models.LoginAccount{
			Email:       r.GetEmail(),
			Username:    r.GetUsername(),
			PhoneNumber: r.GetPhoneNumber(),
			Password:    r.GetPassword(),
			Role:        r.GetRole(),
		}
		return mappedStruct, nil
	case *protos.GetUserIdRequest:
		mappedStruct := &models.GetUserId{
			Username:    r.GetUsername(),
			Email:       r.GetEmail(),
			PhoneNumber: r.GetPhoneNumber(),
		}
		return mappedStruct, nil
	case *protos.GetAccountDetailsRequest:
		mappedStruct := &models.GetAccountData{
			Id: r.GetUserid(),
		}
		return mappedStruct, nil

	case *protos.AddPersonAccountRequest:
		mappedStruct := &models.AddPersonAccountData{
			Username:    r.GetUsername(),
			FirstName:   r.GetFirstName(),
			LastName:    r.GetLastName(),
			PhoneNumber: r.GetPhoneNumber(),
			Email:       r.GetEmail(),
			Password:    r.GetPassword(),
		}
		return mappedStruct, nil

	case *protos.AddFoodPlaceAccountRequest:
		mappedStruct := &models.AddFoodPlaceAccountData{
			Username:     r.GetUsername(),
			PhoneNumber:  r.GetPhoneNumber(),
			Email:        r.GetEmail(),
			Password:     r.GetPassword(),
			BusinessName: r.GetBusinessName(),
			Location:     r.GetLocation(),
			Tags:         r.GetTags(),
		}
		return mappedStruct, nil
	default:
		return nil, errors.New("request type assertion failed")
	}
}
