package battle

type countVictoryAndDefeatAndDraw int

// Count 		ゲームの回数をカウントする
// Player1Win		プレイヤー1の勝ちをカウント
// Player1Defeat	プレイヤー1の負けをカウント
// PlayerDraw		アイコをカウント
// Player2Win		プレイヤー2の勝ちをカウント
// Player2Defeat	プレイヤー2の負けをカウント
var (
	Count         = 1                          // カウント
	Player1Win    countVictoryAndDefeatAndDraw // プレイヤー1の勝ちをカウント
	Player1Defeat countVictoryAndDefeatAndDraw // プレイヤー1の負けをカウント
	PlayerDraw    countVictoryAndDefeatAndDraw // アイコをカウント

	Player2Win    countVictoryAndDefeatAndDraw // プレイヤー2勝ちをカウント
	Player2Defeat countVictoryAndDefeatAndDraw // プレイヤー2負けをカウント
)

// FingerNum ジャンケンパ入力数字
type FingerNum int

// FingerAppearance ジャンケンパ形
type FingerAppearance string

// Rock			数値1で選択(拳)
// Scissors		数値2で選択(ジャン)
// Paper		数値3で選択(パ)
// Rockname		拳
// Scissorsname	ジャン
// Papername	パ
// Win			勝
// Lose			負
// Draw			アイコ
const (
	Rock     FingerNum = 1
	Scissors FingerNum = 2
	Paper    FingerNum = 3

	Rockname     FingerAppearance = "✊"
	Scissorsname FingerAppearance = "✌️"
	Papername    FingerAppearance = "✋"

	Win  int = 1
	Lose int = -1
	Draw int = 0
)

// Finger ...フィールドの型定義
type Finger struct {
	Value FingerNum
	Name  FingerAppearance
}
