package main

import (
	"app/pages"
	"app/pages/components"
	"app/repositories"
	"app/services"
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	port := 8081

	services := &services.Services{
		Repositories: services.Repositories{
			Accounts: repositories.AccountRepository{},
		},
	}

	http.HandleFunc("/components/search-table", func(w http.ResponseWriter, r *http.Request) {
		accounts := services.Repositories.Accounts.GetAccounts()

		query := r.URL.Query().Get("query")

		names := []string{}
		for _, acc := range accounts {
			fullName := fmt.Sprintf("%s %s", acc.FirstName, acc.LastName)
			if strings.Contains(strings.ToLower(fullName), strings.ToLower(query)) {
				names = append(names, fullName)
			}
		}

		components.SearchTableTemplate(&components.SearchTableModel{
			Names: names,
		}).Render(context.Background(), w)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		accounts := services.Repositories.Accounts.GetAccounts()

		names := []string{}
		for _, acc := range accounts {
			names = append(names, fmt.Sprintf("%s %s", acc.FirstName, acc.LastName))
		}

		pages.IndexTemplate(&pages.IndexModel{
			Names: names,
		}).Render(context.Background(), w)
	})

	// Start server
	log.Printf("Starting server on port http://localhost:%v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
