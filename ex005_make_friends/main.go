package main

import "fmt"


type person struct {
	name string
	friends []*person
}

func main() {
	john := person{name: "John"}
	paul := person{name: "Paul"}
	george := person{name: "George"}
	ringo := person{name: "Ringo"}
	listFriends(john)
	makeFriends(&john, &paul)
	makeFriends(&john, &ringo)
	makeFriends(&george, &ringo)
	listFriends(john)
	listFriends(george)
}



func makeFriends(ant *person, bee *person) {
	ant.friends = append(ant.friends, bee)
	bee.friends = append(bee.friends, ant)

}

func listFriends(p person) {
	if len(p.friends) == 0 {
		fmt.Printf("%v has no friends.\n", p.name)
		return
	}
	list := p.friends[0].name
	for _, friend := range p.friends[1:] {
		list += ", " + friend.name
	}
	fmt.Printf("%v has friends %v.\n", p.name, list)
}
