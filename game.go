package main

//Game represents the current game
type Game struct {
	flop        Deck
	turn        Deck
	river       Deck
	currentDeck Deck
	players     []Player
	match       map[string]Deck
}

//CreateGame is responsible for all the pre meditations for a one game
func CreateGame(players []Player, deck Deck) Game {
	game := Game{}
	deck.shuffle()

	return game
}
