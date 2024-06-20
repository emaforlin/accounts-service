package repositories

import (
	"testing"

	"github.com/emaforlin/accounts-service/x/entities"
	"github.com/stretchr/testify/assert"
)

func TestSelectUser(t *testing.T) {
	mockDB := new(MockDatabase)
	repo := NewAccountMockRepositoryImpl(mockDB)

	input := &entities.GetUserDto{Username: "johndoe"}
	expectedUser := &entities.User{Username: "johndoe"}

	mockDB.On("QueryUser", input).Return(expectedUser, nil)

	user, err := repo.SelectUser(input)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)

	mockDB.AssertExpectations(t)
}

func TestDeleteUser(t *testing.T) {
	mockDB := new(MockDatabase)
	repo := NewAccountMockRepositoryImpl(mockDB)

	input := &entities.GetUserDto{ID: 1}
	mockDB.On("ExecuteDeleteUser", input).Return(nil)

	err := repo.DeleteUser(input)
	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestInsertPerson(t *testing.T) {
	mockDB := new(MockDatabase)
	repo := NewAccountMockRepositoryImpl(mockDB)

	input := &entities.InsertPersonDto{
		FirstName: "John",
		LastName:  "Doe",
		User: entities.InsertUserDto{
			Role:        "Customer",
			Username:    "johndoe",
			Email:       "johndoe@example.com",
			PhoneNumber: "1234567890",
			Password:    "password",
		},
	}

	mockDB.On("ExecuteInsertPerson", input).Return(nil)

	err := repo.InsertPerson(input)
	assert.NoError(t, err)

	mockDB.AssertExpectations(t)
}
