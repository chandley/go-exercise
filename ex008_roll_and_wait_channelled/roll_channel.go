package main

import "time"
import "fmt"
import "math/rand"
import "sort"
import "strconv"

type player struct {
	name  string
	total int
}

type roll struct {
	score   int
	message string
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	ticker := time.NewTicker(time.Second)
	go func() {
		for _ = range ticker.C {
			fmt.Printf(".")
		}
	}()

	gameTimer := make(chan string)
	go func() {
		time.Sleep(15 * time.Second)
		gameTimer <- "Time is up!"
	}()

	players := []player{}

	for i := 0; i < 1000000; i++ {
		players = append(players, player{name: "player " + strconv.Itoa(i)})
	}



	//players := []player{player{name: "Sid"}, player{name: "Nancy"}, player{name: "Gavin"}, player{name: "Tracey"}}
	for index := range players {
		go tally(&players[index], keepRolling())
	}

	fmt.Println(<-gameTimer)

	sort.Sort(ByTotal(players))
	showFrom := len(players) - 100
	for _, finisher := range players[showFrom:] {
		fmt.Printf("%v scored total %v\n", finisher.name, finisher.total)
	}
	sort.Sort(sort.Reverse(ByTotal(players)))
	if players[0].total == players[1].total {
		fmt.Printf("It was a draw")
	} else {
		winner := players[0]
		fmt.Printf("%v won with score %v", winner.name, winner.total)
	}

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
	//	fmt.Printf("%v (%v) rolled %v \n", player.name, player.total, roll.message)
	}
}

type ByTotal []player

func (p ByTotal) Len() int {
	return len(p)
}
func (p ByTotal) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p ByTotal) Less(i, j int) bool {
	return p[i].total < p[j].total
}
