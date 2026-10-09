package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Запуск: go run main.go https://example.com")
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
		fmt.Fprintln(os.Stderr, "Ошибка проверки:", err)
		os.Exit(1)
	}
	defer response.Body.Close()

	fmt.Println("Адрес:", address)
	fmt.Println("HTTP-статус:", response.Status)
	fmt.Printf("Время до получения заголовков: %d мс\n", elapsed.Milliseconds())

	if response.StatusCode >= 200 && response.StatusCode < 300 {
		fmt.Println("Результат: успешный HTTP-ответ")
	} else {
		fmt.Println("Результат: сервер ответил, но статус требует проверки")
	}
}
