package memory

import "sync"

type Games struct {
	Games	[]GameEngine
	sync.Mutex
}

func (g *Games) AddGame(game GameEngine) {
	g.Lock()
	defer g.Unlock()

	g.Games = append(g.Games, game)
}

func (g *Games) UpdateGame(engine GameEngine) {
	g.Lock()
	defer g.Unlock()

	for i := 0; i < len(g.Games); i++ {
		ga := g.Games[i]
		if ga.RoomId == engine.RoomId {
			g.Games[i] = engine
		}
	}


}
