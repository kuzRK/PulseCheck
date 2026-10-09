package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	timeout := flag.Duration(
		"timeout",
		10*time.Second,
		"Maximum wait per URL",
	)

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr,
			"Usage: go run ./cmd/service [flags] <URL> [URL...]")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *timeout <= 0 {
		fmt.Fprintln(os.Stderr, "Timeout must be greater than zero")
		os.Exit(2)
	}

	addresses := flag.Args()

	if len(addresses) == 0 {
		flag.Usage()
		os.Exit(2)
	}

	client := &http.Client{
		Timeout: *timeout,
	}

	allOK := true

	for _, address := range addresses {
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
