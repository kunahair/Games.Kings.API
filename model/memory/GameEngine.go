package memory

type GameEngine struct {
	RoomId	string		`json:"room_id"`
	DeckId	string		`json:"deck_id"`
	Players	[]Player	`json:"players"`
	State	GameState	`json:"state"`
}

func (ge *GameEngine) AddPlayer(player Player) {
	ge.Players = append(ge.Players, player)
}

func (ge *GameEngine) UpdatePlayer(player Player) {
	for i := 0; i < len(ge.Players); i++ {
		if ge.Players[i].Name == player.Name {
			ge.Players[i] = player
			break
		}
	}
}

