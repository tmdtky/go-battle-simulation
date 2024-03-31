package character_test

import (
	"go-battle-simulation/character" // yourprojectpathを適切なパスに置き換えてください
	"testing"
)

func TestAttack(t *testing.T) {
	// このテストは非常に基本的なもので、ランダム要素のために厳密な結果の検証は行いません。
	hero := &character.Character{"勇者", 100, 20}
	demonLord := &character.Character{"魔王", 120, 25}

	hero.Attack(demonLord)

	if demonLord.HP == 120 {
		t.Errorf("魔王がダメージを受けると予想したが、HPは変化しなかった。")
	}
}
