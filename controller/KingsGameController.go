package controller

import (
	"git.seevo.online/jgerlach/Games.Kings.API/model/dto"
	"git.seevo.online/jgerlach/Games.Kings.API/model/memory"
	"git.seevo.online/jgerlach/Games.Kings.API/model/viewmodel"
	"git.seevo.online/jgerlach/Games.Kings.API/service/cardapiservice"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

func GetGame(games *memory.Games) gin.HandlerFunc {
	return func(c *gin.Context) {

		roomId := c.Param("room_id")

		var game memory.GameEngine
		gameFound := false
		for i := 0; i < len(games.Games); i++ {
			g := games.Games[i]
			if g.RoomId == roomId {
				game = g
				gameFound = true
				break
			}
		}

		if !gameFound {
			JsonMessageWithStatus(c, http.StatusNotFound, "game not found")
			return
		}

		c.JSON(http.StatusCreated, game)

	}
}

func PostNewGame(games *memory.Games) gin.HandlerFunc {
	return func(c *gin.Context) {

		min := 1000
		max := 9999

		roomIdInt := rand.Intn(max-min) + min
		roomId := strconv.Itoa(roomIdInt)

		//Check Room ID is not taken
		roomIdAssigned := true
		for {
			for _, val := range games.Games {
				if val.RoomId == roomId {
					roomIdInt = rand.Intn(max-min) + min
					roomId = strconv.Itoa(roomIdInt)
					roomIdAssigned = false
				}
			}
			if roomIdAssigned {
				break
			}
			roomIdAssigned = true
		}


		cardService, err := cardapiservice.NewCardApiService()
		if err != nil {
			JsonMessageWithStatus(c, http.StatusInternalServerError, "Unable to start game")
			return
		}

		newDeckId, err := cardService.NewDeck()
		if err != nil {
			JsonMessageWithStatus(c, http.StatusBadGateway, "unable to create new deck")
			return
		}

		newGame := memory.GameEngine{
			RoomId: roomId,
			DeckId: newDeckId,
			Players: []memory.Player{},
			State:  memory.GameState{
				CurrentCard:    memory.KingsCard{},
				RemainingKings: 4,
			},
		}

		games.AddGame(newGame)

		Dto := dto.NewGameDto{RoomId: roomId}

		c.JSON(http.StatusCreated, Dto)

	}
}

func PostAddPlayer(games *memory.Games) gin.HandlerFunc {
	return func(c *gin.Context) {

		roomId := c.Param("room_id")

		var game memory.GameEngine
		gameFound := false
		for i := 0; i < len(games.Games); i++ {
			g := games.Games[i]
			if g.RoomId == roomId {
				game = g
				gameFound = true
				break
			}
		}

		if !gameFound {
			JsonMessageWithStatus(c, http.StatusNotFound, "game not found")
			return
		}

		//Bind Body to View Model
		var model viewmodel.PlayerViewModel
		if err := c.ShouldBind(&model); err != nil {
			JsonMessageWithStatus(c, http.StatusBadRequest, "Unable to bind model")
			c.Abort()
			return
		}

		//Check Player does not already exist
		for _, val := range game.Players {
			if val.Name == model.Name {
				JsonMessageWithStatus(c, http.StatusConflict, "player name already exists")
				return
			}
		}

		//Add Player to Game
		game.AddPlayer(memory.Player{
			Name: model.Name,
			Cards: []memory.KingsCard{},
		})

		if len(game.Players) == 1 {
			game.State.NextPlayer = model.Name
		}

		games.UpdateGame(game)

		c.JSON(http.StatusCreated, game)

	}
}

func PostPlayerDrawCard(games *memory.Games, rules memory.RulesTopLevel) gin.HandlerFunc {
	return func(c *gin.Context) {

		roomId := c.Param("room_id")

		var game memory.GameEngine
		gameFound := false
		for i := 0; i < len(games.Games); i++ {
			g := games.Games[i]
			if g.RoomId == roomId {
				game = g
				gameFound = true
				break
			}
		}

		if !gameFound {
			JsonMessageWithStatus(c, http.StatusNotFound, "game not found")
			return
		}

		//Bind Body to View Model
		var model viewmodel.PlayerViewModel
		if err := c.ShouldBind(&model); err != nil {
			JsonMessageWithStatus(c, http.StatusBadRequest, "Unable to bind model")
			c.Abort()
			return
		}

		//Find Player in Game
		var player memory.Player
		playerFound := false
		for _, val := range game.Players{
			if val.Name == model.Name {
				player = val
				playerFound = true
			}
		}
		if !playerFound {
			JsonMessageWithStatus(c, http.StatusNotFound, "player not found")
			return
		}

		//Check they are the next player
		if game.State.NextPlayer != player.Name {
			JsonMessageWithStatus(c, http.StatusBadRequest, "not your turn")
			return
		}

		//Draw Card
		cardService, err := cardapiservice.NewCardApiService()
		if err != nil {
			JsonMessageWithStatus(c, http.StatusInternalServerError, "Unable to start game")
			return
		}

		card, err := cardService.DrawCard(game.DeckId)
		if err != nil {
			JsonMessageWithStatus(c, http.StatusBadGateway, "unable to draw card")
			return
		}

		kingsCard := memory.InitKingsCard(card, rules)

		//If there is a King, decrement
		if strings.Contains(kingsCard.Card.Code, "K"){
			game.State.RemainingKings--
		}

		//Update Game State
		game.State.CurrentPlayer = player.Name
		game.State.CurrentCard = kingsCard

		//Update Player Cards
		player.Cards = append(player.Cards, kingsCard)

		game.UpdatePlayer(player)

		//Update Next Player
		for i := 0; i < len(game.Players); i++ {
			if game.Players[i].Name == player.Name {
				//Check if Player is the last
				if i == len(game.Players) - 1 {
					//Got to Player 1
					game.State.NextPlayer = game.Players[0].Name
				} else {
					//Otherwise to go next in queue
					game.State.NextPlayer = game.Players[i+1].Name
				}
			}
		}

		games.UpdateGame(game)

		c.JSON(http.StatusOK, game)
	}
}
