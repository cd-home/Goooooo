package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	for _, e := range os.Environ() {
		fmt.Println(e)
	}
	fmt.Println(os.Getenv("SECRET"))
}
