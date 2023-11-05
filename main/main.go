package main

import (
	"encoding/json"
	"fmt"
	"os"

	httplib "github.com/tsemach/go-http"
)

func First[T, U any](val T, _ U) T {
	return val
}

func main() {
	fmt.Println("main called, os.Getwd():", First(os.Getwd()))

	var http = httplib.NewHTTP[any, any]("https://localhost:8080/health")
	fmt.Println("new http:", http)

	err := http.SetCerts("certs/ca.crt", "certs/client.crt", "certs/client.key")
	if err != nil {
		fmt.Println("[ERROR] on set certificates, err:", err)
		return
	}

	resp, err := http.Get(nil)
	defer resp.Body.Close()

	var rs any
	err = json.NewDecoder(resp.Body).Decode(&rs)
	fmt.Println(rs)
}
