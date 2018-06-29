package main

//Player represents each player in the game
type Player struct {
	money float64
	name  string
}

// CreatePlayer is used for creating a player
func CreatePlayer(m float64, n string) Player {
	player := Player{
		money: m, name: n,
	}
	return player
}
