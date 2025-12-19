package auth

import (
	"time"

	domain "auth-service/internal/domain/auth"
	"auth-service/internal/interfaces"
)

type TokenService struct {
	provider interfaces.TokenProvider
	ttl      time.Duration
	issuer   string
}

func NewTokenService(provider interfaces.TokenProvider, ttl time.Duration, issuer string) *TokenService {
	return &TokenService{provider: provider, ttl: ttl, issuer: issuer}
}

func (s *TokenService) Issue(user domain.User) (string, error) {
	now := time.Now()

	claims := domain.Claims{
		UserID:    user.ID,
		Email:     user.Email,
		Roles:     user.Roles[0],
		IssuedAt:  now,
		ExpiresAt: now.Add(s.ttl),
		Issuer:    s.issuer,
	}

	return s.provider.Generate(claims)
}

func (s *TokenService) Validate(token string) (*domain.Claims, error) {
	return s.provider.Validate(token)
}
