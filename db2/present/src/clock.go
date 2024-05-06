package main

import (
	"fmt"
	"time"
)

//START OMIT
func main() {
	go hello5()
	go hello10()

	// var s string
	// fmt.Scanf("%s", s)
	time.Sleep(30 * time.Second)
}

func hello5() {
	for {
		fmt.Printf("%v: Viva Perón!\n", time.Now().Format("15:04:05"))
		time.Sleep(5 * time.Second)
	}
}

func hello10() {
	for {
		fmt.Printf("%v: Aguante Cristina!\n", time.Now().Format("15:04:05"))
		time.Sleep(10 * time.Second)
	}
}
