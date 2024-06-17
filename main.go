package main

import (
	"app/pages"
	"app/repositories"
	"app/services"
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8081

	services := &services.Services{
		Repositories: services.Repositories{
			Accounts: repositories.AccountRepository{},
		},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pages.IndexTemplate(&pages.IndexModel{
			Services: services,
		}).Render(context.Background(), w)
	})

	// Start server
	log.Printf("Starting server on port http://localhost:%v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
