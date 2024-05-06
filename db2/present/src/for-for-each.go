package main

import (
	"fmt"
)

func main() {
	a := [4]int{626, 629, 625, 624}

	for k, v := range a {
		fmt.Printf("%d %d\n", k, v)
	}
}
