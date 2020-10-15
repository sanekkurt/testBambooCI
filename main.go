package main

import (
	"fmt"
	"time"
)

func service() {
	fmt.Println("Hello from service!!")
}

func main() {
	fmt.Println("main() started")

	go service()
	time.Sleep(1 * time.Second)


	fmt.Println("main() stopped")
}