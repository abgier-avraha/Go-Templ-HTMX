package main

import (
	"app/pages/index"
	"app/repositories"
	"app/services"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Port
	port := 8081

	// Services
	services := &services.Services{
		Repositories: services.Repositories{
			Accounts: repositories.AccountRepository{},
		},
	}

	// Static files
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))

	// Routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		index.Handler(w, r, services)
	})

	// Start server
	log.Printf("Starting server on port http://localhost:%v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
