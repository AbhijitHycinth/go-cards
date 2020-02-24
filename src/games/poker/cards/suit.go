package cards

type suit []string

//NewSuit Metod to create a new deck
func NewSuit() suit {
	return suit{
		"Heart",
		"Diamonds",
		"Spade",
		"Club",
	}
}
