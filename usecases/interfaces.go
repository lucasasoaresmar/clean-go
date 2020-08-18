package usecases

// Token type
type Token string

// TokenService interface
type TokenService interface {
	Create(id string, email string, roles []string) Token
}
