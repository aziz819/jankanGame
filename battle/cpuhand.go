package battle

import (
	mr "math/rand"
	t "time"
)

func RandFinger() (finger Finger) {
	rand := mr.New(mr.NewSource(t.Now().UnixNano()))
	switch rand.Intn(3) { // 0 ~ 3乱数
	case 0:
		finger.Value = Rock
		finger.Name = Rockname
	case 1:
		finger.Value = Scissors
		finger.Name = Scissorsname
	case 2:
		finger.Value = Paper
		finger.Name = Papername
	}
	return
}
