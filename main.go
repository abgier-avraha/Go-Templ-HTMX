package main

import (
	"app/pages/index"
	"app/pages/index/indexcomponents"
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

	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		accounts := services.Repositories.Accounts.GetAccounts()

		query := r.URL.Query().Get("query")

		names := []string{}
		for _, acc := range accounts {
			fullName := fmt.Sprintf("%s %s", acc.FirstName, acc.LastName)
			if strings.Contains(strings.ToLower(fullName), strings.ToLower(query)) {
				names = append(names, fullName)
			}
		}
		if len(names) > 5 {
			names = names[0:5]
		}

		hxTarget := r.Header.Get("Hx-Target")

		// Components
		if hxTarget == "search-template-root" {
			indexcomponents.SearchTableTemplate(&indexcomponents.SearchTableModel{
				Names:        names,
				DefaultQuery: query,
			}).Render(context.Background(), w)
			return
		}

		// Page
		index.IndexTemplate(&index.IndexModel{
			Names:        names,
			DefaultQuery: query,
		}).Render(context.Background(), w)
	})

	// Start server
	log.Printf("Starting server on port http://localhost:%v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
