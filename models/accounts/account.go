package accounts

type Account struct {
	PartyNum    int
	Num         int
	Status      string
	Identifiers []Identifier
}

type Identifier struct {
	Type   string
	ID     string
	Status string
}
