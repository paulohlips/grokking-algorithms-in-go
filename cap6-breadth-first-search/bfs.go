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

	for {
		fmt.Printf("fila: %v \n", queue)
		if "007" == queue[0] {
			fmt.Print("nome encontrado!")
			return
		}
		fmt.Printf("nome não é %s!\n", queue[0])
		queue = append(queue, stack[queue[0]]...)
		queue = queue[1:]
		fmt.Printf("nova fila: %v \n\n\n", queue)
	}
}
