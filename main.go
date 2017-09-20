package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"jankenGame/battle"

	"github.com/fatih/color"
)

func main() {

	reader := bufio.NewReader(os.Stdin) // 標準入力読み込み

	for {
		fmt.Println("------------[✌️ ✊ ✋ GAME]-------------")
		fmt.Println("1~3までひとつ選択してゲームを始めましょう！")
		fmt.Println("① CPU\tVS\tCPU\n② CPU\tVS\tHUMAN\n③ HUMAN\tVS\tHUMAN")
		fmt.Println("ゲームを終了したい場合は9を入力してください。")
		fmt.Println("-------------------------------------")

		data, _, err := reader.ReadLine()
		if err != nil { // 読み込み不具合が発生したら
			log.Println("エラーが発生しました。プログラムを終了します。")
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		input, err := strconv.Atoi(string(data)) // 文字列 → 数値変換

		if err != nil { // 入力されたのを数値に変換して、数字かどうかを確認しています。
			log.Println("数字のみを入力してください。\nゲームを終了したい場合は9を入力してください。")
			continue
		}

		if input == 9 { // 入力されたのは数値9であればゲーム終了します。
			fmt.Println("ゲームを終了しました。")
			break
		}

		switch input {
		case 1:
			battle.CPUDuelCPU()
		case 2:
			battle.HumanDuelCPU()
		case 3:
			battle.HumanDuelHuman()
		default:
			color.Red("------------[注意]-------------")
			color.Red("1~3までの数字を入力してください。\nゲームを終了したい場合は9を入力してください。")
			color.Red("------------------------------\n\n")
			continue
		}
	}
}
