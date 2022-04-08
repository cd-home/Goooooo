package config

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestLoadDotEnv(t *testing.T) {
	godotenv.Load("./testdata/.env")
	t.Log(os.Getenv("APP_NAME"))
}
