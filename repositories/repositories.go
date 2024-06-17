package repositories

type AccountRepository struct {
}

func (e AccountRepository) GetAccounts() []Account {
	return []Account{
		{
			FirstName: "Johnny",
			LastName:  "Appleseed",
		},
	}
}

type Account struct {
	FirstName string
	LastName  string
}
