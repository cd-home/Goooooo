package types

type CreatePermissionParam struct {
	Ptype   string `json:"ptype" binding:"required"`
	Subject string `json:"subject" binding:"required"`
	Object  string `json:"object" binding:"required"`
	Action  string `json:"action" binding:"required"`
	Version string `json:"version" binding:"required"`
}
