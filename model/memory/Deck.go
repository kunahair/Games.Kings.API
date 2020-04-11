package memory

type Deck struct {
	Shuffled	bool	`json:"shuffled"`
	Remaining	int		`json:"remaining"`
	Success		bool	`json:"success"`
	DeckId		string	`json:"deck_id"`
}
