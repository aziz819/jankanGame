package battle

import m "math"

// CPUがランダムで選んだ値と人間が入力した値を比べて勝負決める
func IsAiWin(ai Finger, user Finger) int {

	result := ai.Value - user.Value // CPUのランダム値から人間が選んだ値を引く

	if int(m.Abs(float64(result))) == Paper-Rock { // [絶対を値を求めて == 3(パ) -1(ケン)] → で比較している
		result = -(result) // 2 - 2 = 0  == draw(アイコです)
	}

	if result < 0 { // CPUの勝ちです
		return Win
	} else if result > 0 { // CPUの負けです
		return Lose
	}

	return Draw // アイコです
}
