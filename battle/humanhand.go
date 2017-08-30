package battle

// 人間の出した手を返す
func CreateFinger(n int) (finger Finger) {
	switch n {
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
