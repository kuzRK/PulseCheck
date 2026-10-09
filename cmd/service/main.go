package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run ./cmd/service <URL> [URL...]")
		os.Exit(2)
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	allOK := true

	for _, address := range os.Args[1:] {
		if !checkURL(client, address) {
			allOK = false
		}
		fmt.Println()
	}

	if !allOK {
		os.Exit(1)
	}
}

func checkURL(client *http.Client, address string) bool {
	fmt.Println("URL:", address)

	start := time.Now()
	response, err := client.Get(address)
	elapsed := time.Since(start)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Check failed:", err)
		return false
	}
	defer response.Body.Close()

	fmt.Println("HTTP status:", response.Status)
	fmt.Printf("Time to response headers: %d ms\n", elapsed.Milliseconds())

	if response.StatusCode >= 200 && response.StatusCode < 300 {
		fmt.Println("Result: successful HTTP response")
		return true
	}

	fmt.Println("Result: unexpected HTTP status")
	return false
}
