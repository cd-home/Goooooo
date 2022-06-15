package types

// RegisterParam [Request, POST, Body]
type RegisterParam struct {
	Account  string `json:"account" binding:"required,min=4,max=50"`
	Password string `json:"password" binding:"required,min=6,max=18"`
}

// LoginParam [Request, POST, Body]
type LoginParam struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ModifyPassword [Request, POST, Body]
type ModifyPasswordParam struct {
	OriginPassword string `json:"origin_password" binding:"required"`
	NewPassword    string `json:"new_password" binding:"required"`
}
