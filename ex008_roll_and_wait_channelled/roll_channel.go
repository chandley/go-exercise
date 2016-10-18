package main

import "time"
import "fmt"
import "math/rand"

type player struct {
	name  string
	total int
}

func main() {
	gameTimer := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		gameTimer <- "Time is up!"
	}()

	sid := player{"Sid", 0}
	nancy := player{"Nancy", 0}

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

func rollDie() (roll int, wait int) {
	roll = rand.Intn(6) + 1
	wait = 7 - roll
	return
}

func keepRolling() chan int {
	channel := make(chan int)
	go func() {
		for {
			roll, wait := rollDie()
			fmt.Printf("rolled a %v, waiting %v sec \n", roll, wait)
			channel <- roll
			time.Sleep(time.Duration(wait) * time.Second)
		}
	}()
	return channel
}

func tally(player *player, channel chan int) {
	for {
		roll := <-channel
		player.total += roll
		fmt.Printf("%v (%v) rolled %v \n", player.name, player.total, roll)
	}

}
