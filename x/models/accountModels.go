package models

import pb "github.com/emaforlin/accounts-service/x/handlers/grpc/protos"

type (
	GetUserId struct {
		Username    string `validate:""`
		Email       string `validate:""`
		PhoneNumber string `validate:""`
	}

	AddPersonAccountData struct {
		Username    string `validate:"required,min=4,max=30"`
		FirstName   string `validate:"required,max=40"`
		LastName    string `validate:"required,max=80"`
		PhoneNumber string `validate:"required,e164"`
		Email       string `validate:"required,email"`
		Password    string `validate:"min=8,max=64"`
	}

	GetAccountData struct {
		Id uint32 `validate:"required"`
	}

	AddFoodPlaceAccountData struct {
		Username     string   `validate:"required,min=4,max=20"`
		PhoneNumber  string   `validate:"required,e164"`
		Email        string   `validate:"required,email"`
		Password     string   `validate:"required,min=8,max=64"`
		BusinessName string   `validate:"required,max=256"`
		Location     string   `validate:"required,max=256"`
		Tags         []string `validate:"required"`
	}

	LoginAccount struct {
		Email       string `validate:""`
		Username    string `validate:""`
		PhoneNumber string `validate:""`
		Password    string `validate:""`
		Role        string `validate:""`
	}
)

func (m *GetUserId) ToRPCStruct() *pb.GetUserIdRequest {
	if m.Email != "" {
		return &pb.GetUserIdRequest{
			Identifiers: &pb.GetUserIdRequest_Email{Email: m.Email},
		}
	} else if m.Username != "" {
		return &pb.GetUserIdRequest{
			Identifiers: &pb.GetUserIdRequest_Username{Username: m.Username},
		}
	}
	return &pb.GetUserIdRequest{
		Identifiers: &pb.GetUserIdRequest_PhoneNumber{PhoneNumber: m.PhoneNumber},
	}
}

func (m *AddPersonAccountData) ToRPCStruct() *pb.AddPersonAccountRequest {
	return &pb.AddPersonAccountRequest{
		Username:    m.Username,
		PhoneNumber: m.PhoneNumber,
		Email:       m.Email,
		Password:    m.Password,
		FirstName:   m.FirstName,
		LastName:    m.LastName,
	}
}

func (m *GetAccountData) ToRPCStruct() *pb.GetAccountDetailsRequest {
	return &pb.GetAccountDetailsRequest{Userid: m.Id}
}

func (m *AddFoodPlaceAccountData) ToRPCStruct() *pb.AddFoodPlaceAccountRequest {
	return &pb.AddFoodPlaceAccountRequest{
		Username:     m.Username,
		PhoneNumber:  m.PhoneNumber,
		Email:        m.Email,
		Password:     m.Password,
		BusinessName: m.BusinessName,
		Location:     m.Location,
		Tags:         m.Tags,
	}
}

func (m *LoginAccount) ToRPCStruct() *pb.LoginUserRequest {
	if m.Email != "" {
		return &pb.LoginUserRequest{
			Identifiers: &pb.LoginUserRequest_Email{Email: m.Email},
			Password:    m.Password,
			Role:        m.Role,
		}
	} else if m.Username != "" {
		return &pb.LoginUserRequest{
			Identifiers: &pb.LoginUserRequest_Username{Username: m.Username},
			Password:    m.Password,
			Role:        m.Role,
		}
	}
	return &pb.LoginUserRequest{
		Identifiers: &pb.LoginUserRequest_PhoneNumber{PhoneNumber: m.PhoneNumber},
		Password:    m.Password,
		Role:        m.Role,
	}
}
