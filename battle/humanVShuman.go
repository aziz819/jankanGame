package battle

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"syscall"

	"github.com/fatih/color"
	"golang.org/x/crypto/ssh/terminal"
)

// HumanDuelHuman 人間とCPUのジャンケン
func HumanDuelHuman() {

	var humanPlayer1, humanPlayer2 Finger // CPUプレイヤーと人間プレイヤーの型定義
	fmt.Println("-------[Human[1] VS HUMAN[2]]-------")
	fmt.Printf("\t《遊び方》\nHUMAN[1]とHUMAN[2]の手は二人とも入力が終わってから見せます。\n1から3の数字を入力してください。\nゲーム終了したいなら9を入力してください。\n%d->✊\n%d->✌️\n%d->✋\n9->終了\n----------[1回目]-----------\n", Rock, Scissors, Paper)

	for {

		// 結果が出る前にどちらの手も見せないようにパスワード入力、受け取り式を使用しました。
		color.Blue("HumanPlayer[1]の入力：")
		data1, err := terminal.ReadPassword(int(syscall.Stdin))

		color.Magenta("HumanPlayer[2]の入力：")
		data2, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			log.Println("エラーが発生しました。プログラムを終了します。")
			fmt.Sprintln(syscall.Stderr, err)
			os.Exit(1)
		}

		input1, err := strconv.Atoi(string(data1)) // 文字列 → 数値変換（パース）
		input2, err := strconv.Atoi(string(data2)) // 文字列 → 数値変換（パース）

		if err != nil { // 入力されたのを数値に変換して、数字かどうかを確認しています。
			log.Printf("数字以外の入力は無効です。1~3の数字を入力してください。じゃんけんをやめたい場合は9を入力してください。\n1->✊\n2->✌️\n3->✋\n9->終了\n------------[%d回目]-------------\n", Count)
			continue
		}

		if input1 == 9 || input2 == 9 { // 入力されたのは数値9であればゲーム終了します。
			fmt.Println("HUMAN VS HUMAN を終了しました。")
			fmt.Println("------------[最終結果]-------------")
			color.Yellow("HUMAN[1]の勝ち : %d\tあいこ : %d\tHUMAN[2]の勝ち : %d\nHUMAN[1]の負け : %d\t\t\tHUMAN[2]の負け : %d\n", Player1Win, PlayerDraw, Player2Win, Player1Defeat, Player2Defeat)
			fmt.Print("--------------[終了]---------------\n\n")
			Player1Win = 0
			Player1Defeat = 0
			Player2Win = 0
			Player2Defeat = 0
			PlayerDraw = 0
			Count = 0

			break
		}

		if (input1 > 0 && input1 < 4) && (input2 > 0 && input2 < 4) {

			humanPlayer1 = HumanFinger(input1)
			humanPlayer2 = HumanFinger(input2)

			fmt.Print("HUMAN[1]: ", humanPlayer1.Name, "   VS   ")

			fmt.Print(humanPlayer2.Name, " :HUMAN[2]\n")

			gameResult := Judgement(humanPlayer1, humanPlayer2)

			Count++

			if gameResult == Win {
				fmt.Printf("《HUMAN[1]の勝ちです!》\n\nもう1回じゃんけんしましょう。1から3の数字を入力してください。\nじゃんけんをやめたい場合は9を入力してください。\n1->✊\n2->✌️\n3->✋\n9->終了\n------------[%d回目]-------------\n", Count)
				Player1Win++
				Player2Defeat++
			} else if gameResult == Lose {
				fmt.Printf("《HUMAN[2]の勝ちです!》\n\nもう1回じゃんけんしましょう。1から3の数字を入力してください。\nじゃんけんをやめたい場合は9を入力してください。\n1->✊\n2->✌️\n3->✋\n9->終了\n------------[%d回目]-------------\n", Count)
				Player2Win++
				Player1Defeat++
			} else {
				fmt.Printf("《あいこです!》\n\nもう1回じゃんけんしましょう。1から3の数字を入力してください。\nじゃんけんをやめたい場合は9を入力してください。\n1->✊\n2->✌️\n3->✋\n9->終了\n------------[%d回目]-------------\n", Count)
				PlayerDraw++
			}
		} else {
			color.Red("正しい入力ではありません。1から3の数字を入力してください。\nじゃんけんをやめたい場合は9を入力してください。\n1->✊\n2->✌️\n3->✋\n9->終了\n------------[%d回目]-------------\n", Count)
		}

	}
}
