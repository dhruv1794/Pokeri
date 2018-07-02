package main

import (
	"fmt"
)

func main() {
	cards := newDeck()
	cards.shuffle()
	//cards.deal()
	//	cards.print()
	// player := CreatePlayer(25, "Dhruv")
	// fmt.Println(player)
	var players []Player
	players = append(players, CreatePlayer(25, "Dhruv"))
	players = append(players, CreatePlayer(25, "Pv"))
	players = append(players, CreatePlayer(25, "Pato"))
	players = append(players, CreatePlayer(25, "Ishita"))
	game := CreateGame(players, cards)
	fmt.Println("Flop : ", game.flop)
	fmt.Println("Match : ", game.match)
	fmt.Println("Turn : ", game.turn)
	fmt.Println("River : ", game.river)
	fmt.Println("Burn 1: ", game.Burn1)
	fmt.Println("Burn 2 : ", game.Burn2)
	fmt.Println("Burn 3: ", game.Burn3)
	fmt.Println("Current Deck", game.currentDeck, "Len", len(game.currentDeck))
}
