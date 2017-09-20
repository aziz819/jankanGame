package battle

import (
	"math/rand"
	"time"
)

// CPURandomFinger の手をランダムで選ぶ
func CPURandomFinger() (finger Finger) {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	var num = rand.Intn(3) + 1
	switch FingerNum(num) {
	case Rock:
		finger.Value = Rock
		finger.Name = Rockname
	case Scissors:
		finger.Value = Scissors
		finger.Name = Scissorsname
	case Paper:
		finger.Value = Paper
		finger.Name = Papername
	}
	return
}
