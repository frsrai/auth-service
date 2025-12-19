package jwt

import (
	domain "auth-service/internal/domain/auth"

	"github.com/golang-jwt/jwt/v5"
)

type Provider struct {
	secret string
	issuer string
}

func New(secret string, issuer string) *Provider {
	return &Provider{secret: secret, issuer: issuer}
}

func (p *Provider) Generate(c domain.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   string(c.UserID),
		"roles": c.Roles,
		"email": c.Email,
		"iat":   c.IssuedAt.Unix(),
		"exp":   c.ExpiresAt.Unix(),
		"iss":   p.issuer,
	})

	return token.SignedString([]byte(p.secret))
}
