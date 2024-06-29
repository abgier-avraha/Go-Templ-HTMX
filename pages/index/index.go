package index

import (
	"app/components"
	"app/services"
	"context"
	"fmt"
	"net/http"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request, services *services.Services) {
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

	hxComponent := r.Header.Get("Hx-Target")

	// Render component or page
	switch hxComponent {
	case "search-template-root":
		components.SearchTableTemplate(&components.SearchTableModel{
			Names:        names,
			DefaultQuery: query,
		}).Render(context.Background(), w)
	default:
		IndexTemplate(&IndexModel{
			Names:        names,
			DefaultQuery: query,
		}).Render(context.Background(), w)
	}
}
