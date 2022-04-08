package config

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestNewViper(t *testing.T) {

	_mockProd := true

	if _mockProd {
		godotenv.Load("./testdata/.env")
	}

	tests := []struct {
		App  string
		Mode string
		Path []string
	}{
		{
			App:  "admin",
			Mode: "dev",
			Path: []string{"./testdata/configs/"},
		},
		{
			App:  "api",
			Mode: "dev",
			Path: []string{"./testdata/configs/"},
		},
		{
			App:  "admin",
			Mode: "prod",
		},
	}
	for _, tt := range tests {
		t.Run(tt.Mode, func(t *testing.T) {
			vp := NewViper(tt.App, tt.Mode, tt.Path...)
			t.Log(vp.GetString("APP.SECRET"))
		})
	}
}
