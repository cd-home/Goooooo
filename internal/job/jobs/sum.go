package job

import (
	"log"
	"net/http"
)

func Sum(args ...int64) (int64, error) {
	resp, err := http.Get("http://localhost:8080/api/v1/user/login")
	log.Println(resp)
	log.Println(err)
	sum := int64(0)
	for _, arg := range args {
		sum += arg
	}
	return sum, nil
}
