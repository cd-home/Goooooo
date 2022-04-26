package types

type RegisterParam struct {
	Method   string `json:"method"`
	Account  string `json:"account"`
	Password string `json:"password"`
	Code     string `json:"code"`
}

type LoginParam struct {
	Method   string `json:"method"`
	Account  string `json:"account"`
	Password string `json:"password"`
}
