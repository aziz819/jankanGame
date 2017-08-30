package battle

var (
	Count     int = 1 // カウント
	Youwin    int     // 勝ちをカウント
	Youdefeat int     // 負けをカウント
	Youdraw   int     // アイコをカウント
)

const (
	Rock     int = 1 // ✊
	Scissors int = 2 // ✌️
	Paper    int = 3 // ✋

	Rockname     string = "✊"
	Scissorsname string = "✌️"
	Papername    string = "✋"

	Win  int = 1
	Lose int = -1
	Draw int = 0
)

// フィールドの型定義
type Finger struct {
	Value int
	Name  string
}
