package main

import (
	"fmt"
)

type Stack map[string][]string

func main() {
	stack := Stack{"paul": {"james", "bruce", "clark"}}
	stack["james"] = append(stack["james"], "bond")
	stack["bruce"] = append(stack["bruce"], "banner")
	stack["bond"] = append(stack["bond"], "007")

	/*
		paul --- james --- bond
		 |
		 |----- bruce --- banner
		 |----- clark
	*/

	queue := stack["paul"]

	findNameWithBFS("henry", queue, stack)
}

func findNameWithBFS(name string, queue []string, stack Stack) {
	for {
		fmt.Printf("queue: %v \n", queue)

		if name == queue[0] {
			fmt.Print("We found the name!")
			return
		}

		fmt.Printf("nome isnt %s!\n", queue[0])
		queue = append(queue, stack[queue[0]]...)
		queue = queue[1:]
		fmt.Printf("queue after update: %v \n\n\n", queue)
	}
}
