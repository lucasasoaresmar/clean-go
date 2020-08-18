package api

import (
	"encoding/json"
	"net/http"

	"github.com/lucasasoaresmar/clean-go/adapters/gateways/repositories"
	"github.com/lucasasoaresmar/clean-go/adapters/gateways/token"
	"github.com/lucasasoaresmar/clean-go/entities"
	"github.com/lucasasoaresmar/clean-go/usecases"
)

// UserHandler controller
type UserHandler struct {
	userUseCases *usecases.UserUseCases
}

// Login controller
func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user entities.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	token, _ := uh.userUseCases.Login(user.Email, user.Password)
	json.NewEncoder(w).Encode(&token)
}

// NewUserHandler contructor
func NewUserHandler() *UserHandler {
	tokenService := token.Token{}
	userRepository := repositories.User{}
	userUseCases := usecases.UserUseCases{
		UserRepository: &userRepository,
		TokenService:   &tokenService,
	}

	return &UserHandler{&userUseCases}
}
