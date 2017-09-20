package battle

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/fatih/color"
)

// CPUDuelCPU CPUとCPUのジャンケン
func CPUDuelCPU() {

	var cpu1, cpu2 Finger // CPUプレイヤーと人間プレイヤーの型定義
	fmt.Println("-------[CPU[1] VS CPU[2]]-------")
	fmt.Println("\t   《遊び方》\n一回ずつ数字[1] + enterキーを押したら二つのCPUがランダムで手を出します。\nゲーム終了したいなら9を入力してください。\n----------[1回目]-----------")

	reader := bufio.NewReader(os.Stdin) // 標準入力読み込み

	for { // 入力されたのは数字かどうかを確認しています。もし数字9であればゲーム終了します。
		data, _, err := reader.ReadLine() // 入力値を読み込んでいます。

		if err != nil { // 読み込み不具合が発生したら
			log.Println("エラーが発生しました。プログラムを終了します。")
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		input, err := strconv.Atoi(string(data)) // 文字列 → 数値変換（パース）

		if err != nil { // 入力されたのを数値に変換して、数字かどうかを確認しています。
			log.Printf("数字以外の入力は無効です。数字[1]を入力してください。じゃんけんをやめたい場合は9を入力してください。\n1->✊\n2->✌️\n3->✋\n9->終了\n------------[%d回目]-------------\n", Count)
			continue
		}

		if input == 9 { // 入力されたのは数値9であればゲーム終了します。
			fmt.Println("CPU[1] VS CPU[2] を終了しました。")
			fmt.Println("------------[最終結果]-------------")
			color.Yellow("CPU[1]の勝ち : %d\tあいこ : %d\tCPU[2]の勝ち : %d\nCPU[1]の負け : %d\t\t\tCPU[2]の負け : %d\n", Player1Win, PlayerDraw, Player2Win, Player1Defeat, Player2Defeat)
			fmt.Print("--------------[終了]---------------\n\n")
			Player1Win = 0
			Player1Defeat = 0
			Player2Win = 0
			Player2Defeat = 0
			PlayerDraw = 0
			Count = 0

			break
		}

		if input > 0 && input < 4 {

			cpu1 = CPURandomFinger() // CPUランダムでジャンケンパを選ぶ
			fmt.Print("CPU[1]: ", cpu1.Name, "   VS   ")

			cpu2 = CPURandomFinger() // 人間に入力された数値で選ぶ
			fmt.Print(cpu2.Name, " :CPU[2]\n")

			gameResult := Judgement(cpu1, cpu2) // CPUの手と人間の手を比べる

			Count++

			if gameResult == Win {
				fmt.Printf("《CPU[1]の勝ちです!》\n\nもう1回じゃんけんさせましょう。数字[1]を入力してください。\nじゃんけんをやめたい場合は9を入力してください。\n1->✊\n2->✌️\n3->✋\n9->終了\n------------[%d回目]-------------\n", Count)
				Player1Win++
				Player2Defeat++
			} else if gameResult == Lose {
				fmt.Printf("《CPU[2]の勝ちです!》\n\nもう1回じゃんけんさせましょう。数字[1]を入力してください。\nじゃんけんをやめたい場合は9を入力してください。\n1->✊\n2->✌️\n3->✋\n9->終了\n------------[%d回目]-------------\n", Count)
				Player2Win++
				Player1Defeat++
			} else {
				fmt.Printf("《あいこです!》\n\nもう1回じゃんけんさせましょう。数字[1]を入力してください。\nじゃんけんをやめたい場合は9を入力してください。\n1->✊\n2->✌️\n3->✋\n9->終了\n------------[%d回目]-------------\n", Count)
				PlayerDraw++
			}
		} else {
			fmt.Printf("正しい入力ではありません。数字[1]を入力してください。\nじゃんけんをやめたい場合は9を入力してください。\n1->✊\n2->✌️\n3->✋\n9->終了\n------------[%d回目]-------------\n", Count)
		}

	}
}
