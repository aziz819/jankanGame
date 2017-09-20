package battle

import (
	"math"
)

// Judgement player1がランダムで選んだ値と人間が入力した値を比べて勝負決める
func Judgement(player1 Finger, player2 Finger) int {

	result := player1.Value - player2.Value                 // player1のランダム値から人間が選んだ値を引く
	if FingerNum(math.Abs(float64(result))) == Paper-Rock { // [絶対を値を求めて == 3(パ) -1(ケン)] → で比較している
		result = -(result) // 2 - 2 = 0  == draw(アイコです)
	}

	if result < 0 { // player1の勝ちです
		return Win
	} else if result > 0 { // player2の勝ちです
		return Lose
	}

	return Draw // アイコです
}
