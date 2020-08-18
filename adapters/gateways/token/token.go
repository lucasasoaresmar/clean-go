package token

import (
	"fmt"

	"github.com/lucasasoaresmar/clean-go/usecases"
)

// Token service
type Token struct{}

// Create returns a valid token
func (t *Token) Create(id string, email string, roles []string) usecases.Token {
	tokenString := fmt.Sprintf("this is a token with: id (%s), email(%s) and roles(%s)", id, email, roles)
	return usecases.Token(tokenString)
}
