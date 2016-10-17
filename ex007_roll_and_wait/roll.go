package main

import "time"
import "fmt"
import "math/rand"

func main() {
	go player("Sid")
	go player("Nancy")
	time.Sleep( 30 * time.Second)
	fmt.Println("Time up!")
}


func player(name string) {
	total := 0
	for  {
		total = rollDie(name,total)
	}
}

func rollDie(name string, total int) int {
	roll := rand.Intn(6) + 1
	total += roll
	wait := 7 - roll
	fmt.Printf("%v (%v) rolled a %v, waiting %v sec", name, total, roll, wait)
	fmt.Println()
	time.Sleep(time.Duration(wait) * time.Second )

	return total
}


