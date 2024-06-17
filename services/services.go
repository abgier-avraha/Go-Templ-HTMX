package services

import (
	"app/repositories"
)

type Services struct {
	Repositories Repositories
}

type Repositories struct {
	Accounts repositories.AccountRepository
}
