package battle

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 人間の手をテストする
func TestCreateFinger(t *testing.T) {
	var result Finger
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := rand.Intn(3) + 1
	result = CreateFinger(num)

	if result.Value == Rock && result.Name == Rockname {
		fmt.Printf("オッケー[%d]==[1] and [%s] == [✊]", result.Value, result.Name)
	} else if result.Value == Scissors && result.Name == Scissorsname {
		fmt.Printf("オッケー[%d]==[2] and [%s] == [✌️]", result.Value, result.Name)
	} else if result.Value == Paper && result.Name == Papername {
		fmt.Printf("オッケー[%d]==[3] and [%s] == [✋]", result.Value, result.Name)
	} else {
		t.Errorf("予想外の結果：[%d] and [%s]", result.Value, result.Name)

	}

}

// CPUの手をテストする
func TestRandFinger(t *testing.T) {
	var result Finger
	result = RandFinger()

	if result.Value == Rock && result.Name == Rockname {
		fmt.Printf("オッケー[%d]==[1] and [%s] == [✊]", result.Value, result.Name)
	} else if result.Value == Scissors && result.Name == Scissorsname {
		fmt.Printf("オッケー[%d]==[2] and [%s] == [✌️]", result.Value, result.Name)
	} else if result.Value == Paper && result.Name == Papername {
		fmt.Printf("オッケー[%d]==[3] and [%s] == [✋]", result.Value, result.Name)
	} else {
		t.Errorf("予想外の結果：[%d] and [%s]", result.Value, result.Name)

	}
}

func TestIsAiWin(t *testing.T) {
	var cpu, user Finger
	cpu = RandFinger()

	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := rand.Intn(3) + 1
	user = CreateFinger(num)

	result := IsAiWin(cpu, user)

	if result == 1 {
		fmt.Printf("オッケー[%d] == Win[1]", result)
	} else if result == -1 {
		fmt.Printf("オッケー[%d] == Lose[-1]", result)
	} else if result == 0 {
		fmt.Printf("オッケー[%d] == Draw[0]", result)
	} else {
		t.Errorf("予想外の結果：[%d]", result)

	}
}
