package battle

import "jyanken/battle"

// HumanFinger 人間の出した手を返す
func HumanFinger(n int) (finger Finger) {
	switch n {
	case battle.Rock:
		finger.Value = battle.Rock
		finger.Name = battle.Rockname
	case battle.Scissors:
		finger.Value = battle.Scissors
		finger.Name = battle.Scissorsname
	case battle.Paper:
		finger.Value = battle.Paper
		finger.Name = battle.Papername
	}
	return
}
