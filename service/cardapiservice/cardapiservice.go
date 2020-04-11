package cardapiservice

import (
	"encoding/json"
	"errors"
	"git.seevo.online/jgerlach/Games.Kings.API/model/memory"
	"net/http"
)

type ICardApiService interface {
	NewDeck() (deckId string, err error)
	DrawCard(deckId string) (memory.Card, error)
}

type cardApiService struct {
}

func NewCardApiService() (cs ICardApiService, err error) {

	cardService := cardApiService{}

	return cardService, nil
}

func (cs cardApiService) NewDeck() (deckId string, err error) {
	//Make request to get new DeckID
	resp, err := http.Get("https://deckofcardsapi.com/api/deck/new/shuffle")
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var deck memory.Deck
	err = decoder.Decode(&deck)
	if err != nil {
		return "", err
	}

	if deck.Success == false {
		return "", errors.New("unable to create new deck")
	}

	return deck.DeckId, nil
}

func (cs cardApiService) DrawCard(deckId string) (memory.Card, error) {

	var card memory.Card

	resp, err := http.Get("https://deckofcardsapi.com/api/deck/" + deckId + "/draw/?count=1")
	if err != nil {
		return card, err
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var drawCard memory.DrawCard
	err = decoder.Decode(&drawCard)
	if err != nil {
		return card, err
	}

	if !drawCard.Success {
		return card, errors.New("unable to draw card")
	}

	card = drawCard.Cards[0]
	card.Image = "http://kings.specialfriends.chat/static/img/" + card.Code + ".png"

	return card, nil
}
