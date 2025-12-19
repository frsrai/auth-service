package auth

type Permission string

const (
	PermUserCreate Permission = "user:create"
	PermUserUpdate Permission = "user:update"
	PermUserDelete Permission = "user:delete"
	PermUserRead   Permission = "user:read"
)
