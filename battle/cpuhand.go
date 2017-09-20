package battle

import (
	"jyanken/battle"
	"math/rand"
	"time"
)

// CPURandomFinger の手をランダムで選ぶ
func CPURandomFinger() (finger Finger) {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	var num = rand.Intn(3) + 1
	switch num {
	case battle.Rock:
		finger.Value = battle.Rock
		finger.Name = battle.Rockname
	case battle.Scissors:
		finger.Value = battle.Scissors
		finger.Name = battle.Scissorsname
	case battle.Draw:
		finger.Value = battle.Paper
		finger.Name = battle.Papername
	}
	return
}
