package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type player struct {
	name string
	pos  int
}

var playerCnt int

var players []player

var snakes = map[int]int{
	22: 17,
	33: 4,
	47: 13,
	69: 25,
	77: 42,
	84: 37,
	93: 58,
	99: 67,
}

var ladders = map[int]int{
	3:  15,
	8:  26,
	12: 48,
	21: 43,
	37: 59,
	61: 81,
	73: 93,
}

var max_pos int = 100
var gameStatus bool = false

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func setPlayerName() {
	for i := 1; i <= playerCnt; i++ {
		p := player{name: fmt.Sprintf("Player %d", i), pos: 0}
		players = append(players, p)
	}
}

func playGame() {
	for !gameStatus {
		for key, v := range players {
			fmt.Printf("%s roll dice by pressing 5 : ", v.name)
			var pressedNumber int
			fmt.Scanf("%d", &pressedNumber)
			for pressedNumber != 5 {
				fmt.Println()
				fmt.Printf("Try again to roll your dice by pressing 5 : ")
				fmt.Scanf("%d", &pressedNumber)
			}
			number := random(1, 6)
			fmt.Println("Rolled : ", number)
			temp := v.pos + number
			if temp == max_pos {
				gameStatus = true
				fmt.Printf("%s won the game", v.name)
				break
			}
			if temp > max_pos {
				fmt.Println("Try again in next turn")
				fmt.Println()
				continue
			}
			if revPos, found := snakes[temp]; found {
				fmt.Println("Snake Bite")
				players[key].pos = revPos
			} else if revPos, found := ladders[temp]; found {
				fmt.Println("Got Ladder")
				players[key].pos = revPos
			} else {
				players[key].pos += number
			}
			fmt.Printf("%s moved to position : %d", v.name, players[key].pos)
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println()
		}
	}
}

func main() {

	fmt.Println("---------------Snake & Ladder----------------")
	fmt.Println("Snakes : ")
	for k, v := range snakes {
		fmt.Println("Head : ", k, "Tail : ", v)
	}
	fmt.Println("Ladders : ")
	for k, v := range ladders {
		fmt.Println("From : ", k, "To : ", v)
	}
	fmt.Printf("Enter the number of player: ")

	min, max := 2, 4
	fmt.Scanf("%d", &playerCnt)

	if playerCnt < min || playerCnt > max {
		fmt.Println("Min Player : ", min, "Max Player : ", max)
		os.Exit(-1)
	}

	setPlayerName()
	playGame()
	fmt.Println()
}
