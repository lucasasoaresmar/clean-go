package usecases

import (
	"github.com/lucasasoaresmar/clean-go/entities"
)

// mockedUserRepository ...
type mockedUserRepository struct {
	expectedUser  entities.User
	expectedError error
}

// GetByEmailAndPassword method mocked
func (mur *mockedUserRepository) GetByEmailAndPassword(email string, password string) (user *entities.User, err error) {
	user = &mur.expectedUser
	err = mur.expectedError
	if err != nil {
		return
	}

	return
}

// mockedTokenService ...
type mockedTokenService struct {
	expectedToken Token
}

// Create method mocked
func (mts *mockedTokenService) Create(id string, email string, roles []string) Token {
	return mts.expectedToken
}
