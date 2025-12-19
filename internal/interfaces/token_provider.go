package interfaces

import "auth-service/internal/domain/auth"

type TokenProvider interface {
	Generate(claims auth.Claims) (string, error)
	Validate(token string) (*auth.Claims, error)
}
