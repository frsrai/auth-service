package auth

import "time"

type User struct {
	ID        UserID
	Email     Email
	Password  HashedPassword
	Roles     []Role
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) HasRole(role Role) bool {
	for _, r := range u.Roles {
		if r == role {
			return true
		}
	}
	return false
}

func (u *User) Activate() {
	u.IsActive = true
}

func (u *User) Deactivate() {
	u.IsActive = false
}

type Claims struct {
	UserID    UserID
	Email     Email
	Roles     Role
	IssuedAt  time.Time
	ExpiresAt time.Time
	Issuer    string
}
