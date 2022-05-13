package types

type CommonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type RegisterParam struct {
	Account  string `json:"account" binding:"required,min=4,max=50"`
	Password string `json:"password" binding:"required,min=6,max=18"`
}

type LoginParam struct {
	Method   string `json:"method"`
	Account  string `json:"account"`
	Password string `json:"password"`
}
