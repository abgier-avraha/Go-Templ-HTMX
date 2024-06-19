package repositories

type AccountRepository struct {
}

func (e AccountRepository) GetAccounts() []Account {
	return []Account{
		{
			FirstName: "Johnny",
			LastName:  "Appleseed",
		},
		{
			FirstName: "Martha",
			LastName:  "Washington",
		},
		{
			FirstName: "Marcellus", LastName: "Nolan"},
		{
			FirstName: "Coleen", LastName: "James"},
		{
			FirstName: "Zachary", LastName: "Gillespie"},
		{
			FirstName: "Saundra", LastName: "Joseph"},
		{
			FirstName: "Amalia", LastName: "Marsh"},
		{
			FirstName: "Wes", LastName: "Dougherty"},
		{
			FirstName: "Sallie", LastName: "Blackburn"},
		{
			FirstName: "Abram", LastName: "Rangel"},
		{
			FirstName: "Rhonda", LastName: "Zhang"},
		{
			FirstName: "Rosario", LastName: "Willis"},
		{
			FirstName: "Melody", LastName: "Cross"},
		{
			FirstName: "Clint", LastName: "Love"},
		{
			FirstName: "Bonnie", LastName: "Palmer"},
		{
			FirstName: "Dwain", LastName: "Bullock"},
		{
			FirstName: "Dusty", LastName: "Flowers"},
		{
			FirstName: "Porter", LastName: "Brady"},
		{
			FirstName: "Lloyd", LastName: "Bowen"},
		{
			FirstName: "Sharlene", LastName: "Mccann"},
		{
			FirstName: "Tanya", LastName: "Casey"},
		{
			FirstName: "Bernadine", LastName: "Watkins"},
	}
}

type Account struct {
	FirstName string
	LastName  string
}
