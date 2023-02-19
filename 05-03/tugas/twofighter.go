package main

import "fmt"

type Fighter struct {
	Name            string
	Health          int
	DamagePerAttack int
}

func DeclareWinner(fighter1 Fighter, fighter2 Fighter, firstAttacker string) string {
	for {
		if firstAttacker == fighter1.Name {
			fighter2.Health -= fighter1.DamagePerAttack
			firstAttacker = fighter2.Name
		} else {
			fighter1.Health -= fighter2.DamagePerAttack
			firstAttacker = fighter1.Name
		}

		if fighter1.Health <= 0 {
			return fighter2.Name
		}
		if fighter2.Health <= 0 {
			return fighter1.Name
		}
	}
}

func main() {

	fighter1 := Fighter{"Lew", 10, 2}
	fighter2 := Fighter{"Harry", 5, 4}

	fmt.Println(DeclareWinner(fighter1, fighter2, "Lew"))
}
