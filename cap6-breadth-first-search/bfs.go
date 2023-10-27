package main

import "fmt"

type Stack map[string][]string

func main() {
	stack := Stack{"js": {"node", "react"}}
	stack["java"] = append(stack["java"], "spring")

	fmt.Printf("%s\n", stack["java"][0])
}
