package types

type CreatePermissionParam struct {
	P []string `json:"p"`
	G []string `json:"g"`
}

type CreatePermissionsParam struct {
	P [][]string `json:"p"`
	G [][]string `json:"g"`
}
