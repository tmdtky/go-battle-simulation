package character

import (
	"fmt"
	"math/rand"
	"time"
)

type Character struct {
	Name        string
	HP          int
	AttackPower int
}

func (c *Character) Attack(target *Character) {
	rand.Seed(time.Now().UnixNano())
	chance := rand.Intn(100)

	if chance < 10 {
		fmt.Printf("%s は攻撃をミスした！\n", c.Name)
		return
	}

	damage := c.AttackPower
	if chance < 30 {
		damage = int(float64(damage) * 1.5)
		fmt.Printf("クリティカルヒット! ")
	}

	target.HP -= damage
	fmt.Printf("%s は %s に対して攻撃した。 %d のダメージ！\n", c.Name, target.Name, damage)
}
