package config

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestLoadDotEnv(t *testing.T) {
	err := godotenv.Load("./testdata/.env")
	t.Log(err)

	t.Log(os.Getenv("SECRET"))
}
