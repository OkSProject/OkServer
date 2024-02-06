package main

import (
	"fmt"
	"net/http"
)

func main() {

	// Serves site(s) from directory "/".
	dir := "./www/html/"
	file_server := http.FileServer(http.Dir(dir))
	http.Handle("/", file_server)

	// Sets server to listen on Port 8080.
	port := 8080

	// Starts server.
	fmt.Printf("OkServer! Running on: %d...\n", port)

	// Error checking
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
