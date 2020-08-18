package usecases

import (
	"github.com/lucasasoaresmar/clean-go/entities"
)

// UserRepository to be implemented in the database domain
type UserRepository interface {
	GetByEmailAndPassword(email string, password string) (user *entities.User, err error)
}

// UserUseCases ...
type UserUseCases struct {
	UserRepository UserRepository
	TokenService   TokenService
}

// Login returns a token or an error
func (u *UserUseCases) Login(email string, password string) (token Token, err error) {
	user, err := u.UserRepository.GetByEmailAndPassword(email, password)
	if err != nil {
		return
	}

	token = u.TokenService.Create(user.ID, user.Email, user.Roles)
	return
}
