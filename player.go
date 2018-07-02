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

//AddMoney adds money for a player
func (player *Player) AddMoney(money float64) {
	(*player).money += money
}

//RemoveMoney removes money from a player
func (player *Player) RemoveMoney(money float64) {
	(*player).money -= money
}

//placeBet function places bet which adds money to the pot and reduces money of the player

func (player *Player) placeBet(money float64, pot *Pot) string {
	if money > (*player).money {
		return "Sorry you are are short of cash!"
	}
	(*player).RemoveMoney(money)
	(*pot).addMoneyToPot(money)
	return "Successful"
}
