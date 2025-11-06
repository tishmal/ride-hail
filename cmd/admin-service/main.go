package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("admin service is running 🚖")
	// Блокируем, чтобы контейнер не завершался
	for {
		time.Sleep(10 * time.Second)
	}
	select {}
}
