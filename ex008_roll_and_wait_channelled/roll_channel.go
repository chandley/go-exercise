package main

import "time"
import "fmt"
import "math/rand"

type player struct {
	name  string
	total int
}

type roll struct {
	score   int
	message string
}

func main() {
	gameTimer := make(chan string)
	go func() {
		time.Sleep(30 * time.Second)
		gameTimer <- "Time is up!"
	}()

	sid := player{name: "Sid"}
	nancy := player{name: "Nancy"}

	go tally(&sid, keepRolling())
	go tally(&nancy, keepRolling())

	fmt.Println(<-gameTimer)

	winner := player{"Draw! Neither", 0}
	if sid.total < nancy.total {
		winner = nancy
	} else if nancy.total < sid.total {
		winner = sid
	}
	fmt.Printf("%v won", winner.name)
	fmt.Println()
}

func rollDie() (score int, wait int) {
	score = rand.Intn(6) + 1
	wait = 7 - score
	return
}

func keepRolling() chan roll {
	channel := make(chan roll)
	go func() {
		for {
			score, wait := rollDie()
			msg := fmt.Sprintf("rolled a %v, waiting %v sec", score, wait)
			channel <- roll{score, msg}
			time.Sleep(time.Duration(wait) * time.Second)
		}
	}()
	return channel
}

func tally(player *player, channel chan roll) {
	for {
		roll := <-channel
		player.total += roll.score
		fmt.Printf("%v (%v) rolled %v \n", player.name, player.total, roll.message)
	}

}
