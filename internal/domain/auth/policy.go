package auth

var RolePermission = map[Role][]Permission{
	RoleAdmin: {
		PermUserCreate,
		PermUserUpdate,
		PermUserDelete,
		PermUserRead,
	},
	RoleUser: {
		PermUserRead,
	},
}
