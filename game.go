package main

//Game represents the current game
type Game struct {
	flop        Deck
	turn        Card
	river       Card
	Burn1       Card
	Burn2       Card
	Burn3       Card
	currentDeck Deck
	players     []Player
	match       map[string]Hand
}

//Hand represents individual hand
type Hand struct {
	firstCard  Card
	secondCard Card
}

//CreateGame is responsible for all the pre meditations for a one game
func CreateGame(players []Player, deck Deck) Game {
	game := Game{}
	burn1 := Card{}
	burn2 := Card{}
	burn3 := Card{}
	deck.shuffle()
	match := Distribute(players, &deck, len(players))
	burn1 = deck.Burn()
	flop := deck.createFlop()
	burn2 = deck.Burn()
	turn := deck.createTurnOrRiver()
	burn3 = deck.Burn()
	river := deck.createTurnOrRiver()
	game.flop = flop
	game.match = match
	game.turn = turn
	game.river = river
	game.currentDeck = deck
	game.Burn1 = burn1
	game.Burn2 = burn2
	game.Burn3 = burn3
	return game
}

//Distribute is responsible for distributing and creating a map for players and their respective set of cards
func Distribute(players []Player, d *Deck, playerCount int) map[string]Hand {
	var firstcard Card
	var secondCard Card
	var counterFirstHalf int
	var counterSecondHalf int
	counterFirstHalf = 0
	counterSecondHalf = 0
	//var HANDS map[string]Hand
	HANDS := make(map[string]Hand)
	var playerName string
	for i := 0; i < playerCount*2; i++ {

		if i < playerCount {
			playerName = players[counterFirstHalf].name
			*d, firstcard = (*d).deal()
			counterFirstHalf++

			HANDS[playerName] = Hand{firstCard: firstcard}
			//	HANDS[playerName].firstCard = firstcard
		} else {
			playerName = players[counterSecondHalf].name
			*d, secondCard = (*d).deal()
			counterSecondHalf++
			//fmt.Println(HANDS[playerName].firstCard)
			HANDS[playerName] = Hand{firstCard: HANDS[playerName].firstCard, secondCard: secondCard}
			//HANDS[playerName].secondCard = secondCard
		}

	}
	return HANDS

}

//This function is responsible for returning a flop
func (d *Deck) createFlop() Deck {
	var flop Deck
	var card Card
	for i := 0; i < 3; i++ {
		*d, card = (*d).deal()
		flop = append(flop, card)
	}
	return flop
}

// To create turn and river
func (d *Deck) createTurnOrRiver() Card {
	var card Card
	*d, card = (*d).deal()
	return card
}

// Burn while distribution
func (d *Deck) Burn() Card {
	var card Card
	*d, card = (*d).deal()
	return card
}
