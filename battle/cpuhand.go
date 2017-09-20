package battle

import (
	"math/rand"
	"time"
)

// CPURandomFinger の手をランダムで選ぶ
func CPURandomFinger() (finger Finger) {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	var num = rand.Intn(3) + 1
	switch num {
	case 1:
		finger.Value = Rock
		finger.Name = Rockname
	case 2:
		finger.Value = Scissors
		finger.Name = Scissorsname
	case 3:
		finger.Value = Paper
		finger.Name = Papername
	}
	return
}
