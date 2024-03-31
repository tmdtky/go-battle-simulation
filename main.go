package main

import (
	"fmt"
	"go-battle-simulation/character"
)

func main() {
	hero := &character.Character{Name: "勇者", HP: 100, AttackPower: 20}
	demonLord := &character.Character{Name: "魔王", HP: 120, AttackPower: 25}

	fmt.Printf("戦闘開始！\n")

	for hero.HP > 0 && demonLord.HP > 0 {
		hero.Attack(demonLord)
		if demonLord.HP <= 0 {
			fmt.Printf("魔王は倒れた。世界に平和が訪れた。")
			break
		}

		demonLord.Attack(hero)
		if hero.HP <= 0 {
			fmt.Printf("勇者は倒れた。世界は闇に包まれた。")
			break
		}
	}
}
