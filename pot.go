package main

//Pot represents the pot value and rake division
type Pot struct {
	totalPot float64
	rake     float64
}

//addMoneyToPot adds money to pot
func (p *Pot) addMoneyToPot(money float64) {
	(*p).totalPot += money
}

//makeRake function makes the rake of the total pot
func (p *Pot) makeRake() {
	switch {
	case (*p).totalPot < 100:
		(*p).rake = 0
	default:
		rakeAmount := 0.0
		rakeAmount = (*p).totalPot * 11 / 100
		if rakeAmount > 1100 {
			rakeAmount = 1100.0
		}
		(*p).rake = rakeAmount

	}
}
