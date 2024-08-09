package models

import pb "github.com/emaforlin/accounts-service/x/handlers/grpc/protos"

type (
	GetUserId struct {
		Username    string `json:"username" validate:""`
		Email       string `json:"email" validate:""`
		PhoneNumber string `json:"phone_number" validate:""`
	}

	AddPersonAccountData struct {
		Username    string `json:"username" validate:"required,min=4,max=30"`
		FirstName   string `json:"firstname" validate:"required,max=40"`
		LastName    string `json:"lastname" validate:"required,max=80"`
		PhoneNumber string `json:"phone_number" validate:"required,e164"`
		Email       string `json:"email" validate:"required,email"`
		Password    string `json:"password" validate:"min=8,max=64"`
	}

	GetAccountData struct {
		Id uint32 `json:"id" validate:"required"`
	}

	AddFoodPlaceAccountData struct {
		Username     string   `json:"username" validate:"required,min=4,max=20"`
		PhoneNumber  string   `json:"phone_number" validate:"required,e164"`
		Email        string   `json:"email" validate:"required,email"`
		Password     string   `json:"password" validate:"required,min=8,max=64"`
		BusinessName string   `json:"business_name" validate:"required,max=256"`
		Location     string   `json:"location" validate:"required,max=256"`
		Tags         []string `json:"tags" validate:"required"`
	}

	LoginAccount struct {
		Email       string `json:"email" validate:""`
		Username    string `json:"username" validate:""`
		PhoneNumber string `json:"phone_number" validate:""`
		Password    string `json:"password" validate:""`
		Role        string `json:"role" validate:""`
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
