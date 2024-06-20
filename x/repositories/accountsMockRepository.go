package repositories

import (
	"github.com/emaforlin/accounts-service/x/entities"
	"github.com/stretchr/testify/mock"
)

// MockDatabase simula las interacciones con la base de datos.
type MockDatabase struct {
	mock.Mock
}

func (db *MockDatabase) QueryUser(dto *entities.GetUserDto) (*entities.User, error) {
	args := db.Called(dto)
	user := args.Get(0)
	if user == nil {
		return nil, args.Error(1)
	}
	return user.(*entities.User), args.Error(1)
}

func (db *MockDatabase) ExecuteDeleteUser(dto *entities.GetUserDto) error {
	args := db.Called(dto)
	return args.Error(0)
}

func (db *MockDatabase) ExecuteInsertPerson(dto *entities.InsertPersonDto) error {
	args := db.Called(dto)
	return args.Error(0)
}

// accountsRepositoryImpl representa una implementación del repositorio que usa MockDatabase.
type accountsRepositoryImpl struct {
	db *MockDatabase
}

func (repo *accountsRepositoryImpl) SelectUser(in *entities.GetUserDto) (*entities.User, error) {
	return repo.db.QueryUser(in)
}

func (repo *accountsRepositoryImpl) DeleteUser(in *entities.GetUserDto) error {
	return repo.db.ExecuteDeleteUser(in)
}

func (repo *accountsRepositoryImpl) InsertPerson(in *entities.InsertPersonDto) error {
	return repo.db.ExecuteInsertPerson(in)
}

// NewAccountMockRepositoryImpl crea una nueva instancia de accountsRepositoryImpl con MockDatabase.
func NewAccountMockRepositoryImpl(db *MockDatabase) *accountsRepositoryImpl {
	return &accountsRepositoryImpl{db: db}
}
