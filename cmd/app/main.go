package main

import (
	"fmt"
	"url-shortener/internal/server"
)

func main() {
	fmt.Println("Starting server...")
	server.NewServer().Start()
}
