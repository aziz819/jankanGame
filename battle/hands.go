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

// Rock 		数値1で選択(拳)
// Scissors 	数値2で選択(ジャン)
// Paper		数値3で選択(パ)
// Rockname		拳
// Scissorsname ジャン
// Papername	パ
// Win			勝
// Lose			負
// Draw			アイコ
const (
	Rock     int = 1
	Scissors int = 2
	Paper    int = 3

	Rockname     string = "✊"
	Scissorsname string = "✌️"
	Papername    string = "✋"

	Win  int = 1
	Lose int = -1
	Draw int = 0
)

// Finger ...フィールドの型定義
type Finger struct {
	Value int
	Name  string
}
