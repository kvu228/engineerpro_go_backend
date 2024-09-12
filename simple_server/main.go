package main

import (
	"fmt"
	"net/http"
	"simple_server/handlers"
	"simple_server/storage"
)

func main() {
	// Load users from file on startup
	if err := storage.LoadUsers(); err != nil {
		fmt.Printf("Error loading users: %v\n", err)
		return
	}

	// Set up routes
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/profile", handlers.ProfileHandler)

	// Start the server
	fmt.Println("Server running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		return
	}
}
