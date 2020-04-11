package memory

type GameState struct {
	CurrentCard    KingsCard `json:"current_card"`
	CurrentPlayer  string    `json:"current_player"`
	NextPlayer     string    `json:"next_player"`
	RemainingKings int       `json:"remaining_kings"`
	RemainingCards int       `json:"remaining_cards"`
}
