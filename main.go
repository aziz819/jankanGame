package main

import (
	b "bufio"
	f "fmt"
	o "os"
	s "strconv"

	"github.com/user/janken/battle"
)

func main() {
	// CPUプレイヤーと人間プレイヤーの型定義
	var aiFinger, userFinger battle.Finger

	reader := b.NewReader(o.Stdin) // 標準入力読み込み

	f.Printf("CPUとジャンケンしましょう!\n\t《遊び方》\n1から3の数字を入力してください。\nゲーム終了したいなら9を入力してください。\n%d->✊\n%d->✌️\n%d->✋\n9->終了\n------------[1回目]-------------\n", battle.Rock, battle.Scissors, battle.Paper)
	for { // 入力されたのは数字かどうかを確認しています。もし数字9であればゲーム終了します。
		data, _, err := reader.ReadLine() // 入力値を読み込んでいます。

		if err != nil { // 読み込み不具合が発生したら
			f.Printf("エラーが発生しました。プログラムを終了します。")
			break
		}

		input, err := s.Atoi(string(data)) // 文字列 → 数値変換（パース）

		if err != nil { // 入力されたのを数値に変換して、数字かどうかを確認しています。
			f.Printf("正しい入力ではありません。1から3の数字を入力してください。じゃんけんをやめたい場合は9を入力してください。\n1->✊\n2->✌️\n3->✋\n9->終了\n------------[%d回目]-------------\n", battle.Count)
			continue
		}

		if input == 9 { // 入力されたのは数値9であればゲーム終了します。
			f.Println("ゲームを終了しました。\n\n------------[最終結果]-------------\n")
			f.Printf("あなたの勝ち : %d\tあなたの負け : %d\tあいこ : %d\n", battle.Youwin, battle.Youdefeat, battle.Youdraw)
			break
		}

		switch input {
		case battle.Rock, battle.Scissors, battle.Paper: // 人間が入力した 1 ~ 3までの数値で判断
			aiFinger = battle.RandFinger() // CPUランダムでジャンケンパを選ぶ
			f.Print("CPU: ", aiFinger.Name, "   VS   ")

			userFinger = battle.CreateFinger(input) // 人間に入力された数値で選ぶ
			f.Print(userFinger.Name, " :あなた\n")

			aiWin := battle.IsAiWin(aiFinger, userFinger) // CPUの手と人間の手を比べる

			battle.Count += 1

			if aiWin == battle.Win {
				f.Printf("《コンピュータの勝ちです》\n\nもう1回じゃんけんしましょう。1から3の数字を入力してください。\nじゃんけんをやめたい場合は9を入力してください。\n1->✊\n2->✌️\n3->✋\n9->終了\n------------[%d回目]-------------\n", battle.Count)
				battle.Youdefeat += 1
			} else if aiWin == battle.Lose {
				f.Printf("《コンピュータの負けです》\n\nもう1回じゃんけんしましょう。1から3の数字を入力してください。\nじゃんけんをやめたい場合は9を入力してください。\n1->✊\n2->✌️\n3->✋\n9->終了\n------------[%d回目]-------------\n", battle.Count)
				battle.Youwin += 1
			} else {
				f.Printf("《あいこです》\n\nもう1回じゃんけんしましょう。1から3の数字を入力してください。\nじゃんけんをやめたい場合は9を入力してください。\n1->✊\n2->✌️\n3->✋\n9->終了\n------------[%d回目]-------------\n", battle.Count)
				battle.Youdraw += 1
			}
		default:
			f.Printf("正しい入力ではありません。1から3の数字を入力してください。\nじゃんけんをやめたい場合は9を入力してください。\n1->✊\n2->✌️\n3->✋\n9->終了\n------------[%d回目]-------------\n", battle.Count)
		}

	}

}
