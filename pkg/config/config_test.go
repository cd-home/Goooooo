package config

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestNewViper(t *testing.T) {
	godotenv.Load("./testdata/.env")
	tests := []struct {
		Mode string
		Path []string
	}{
		{
			Mode: "dev",
		},
		{
			Mode: "prod",
		},
	}
	for _, tt := range tests {
		t.Run(tt.Mode, func(t *testing.T) {
			vp := NewViper(tt.Mode)
			if tt.Mode == "dev" {
				t.Log(vp.GetString("APP.SECRET"))
			} else {
				t.Log(vp.GetString("SECRET"))
			}
		})
	}
}
