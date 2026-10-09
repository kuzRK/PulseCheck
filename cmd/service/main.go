package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run ./cmd/service <URL>")
		os.Exit(2)
	}

	address := os.Args[1]

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	start := time.Now()
	response, err := client.Get(address)
	elapsed := time.Since(start)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Check failed:", err)
		os.Exit(1)
	}
	defer response.Body.Close()

	fmt.Println("URL:", address)
	fmt.Println("HTTP status:", response.Status)
	fmt.Printf("Time to response headers: %d ms\n", elapsed.Milliseconds())

	if response.StatusCode >= 200 && response.StatusCode < 300 {
		fmt.Println("Result: successful HTTP response")
	} else {
		fmt.Println("Result: unexpected HTTP status")
	}
}
