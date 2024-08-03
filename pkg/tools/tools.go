package tools

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func GetTask(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	return bodyBytes
}

func PostTask(url string, sol []byte) {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(sol))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println("Response:", string(body))
}
