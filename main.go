package main

import (
	"fmt"
	"study/http_server"
)

func main() {
	fmt.Println("Starting server...")
	err := http_server.StartHTTPServer()
	if err != nil {
		fmt.Println("Something went wrong with server starting", err)
	}

	fmt.Println("Server otrabotal succeed")
}
