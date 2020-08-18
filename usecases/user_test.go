package usecases

import (
	"errors"
	"testing"

	"github.com/lucasasoaresmar/clean-go/entities"
	. "github.com/smartystreets/goconvey/convey"
)

func TestLogin(t *testing.T) {
	expectedToken := Token("A Token")
	expectedError := errors.New("Some Error")
	expectedUser := entities.User{
		ID:       "1",
		Name:     "John",
		Email:    "john@teste.com.br",
		Password: "pass",
		Roles:    []string{"admin", "user"},
	}

	mus := mockedTokenService{expectedToken}
	mur := mockedUserRepository{expectedUser, nil}
	userUseCases := UserUseCases{&mur, &mus}

	Convey("A user made login", t, func() {
		returnedToken, _ := userUseCases.Login(expectedUser.Email, expectedUser.Password)

		Convey("A token should be returned", func() {
			So(returnedToken, ShouldEqual, expectedToken)
		})

		Convey("An error should be returned in case of database failure", func() {
			mur.expectedError = expectedError
			_, returnedError := userUseCases.Login("john@teste.com.br", "pass")
			So(returnedError, ShouldEqual, expectedError)
		})
	})
}
