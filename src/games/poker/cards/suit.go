package cards

type suit []string

func NewSuit() suit {
	return suit{
		"Heart",
		"Diamonds",
		"Spade",
		"Club",
	}
}
