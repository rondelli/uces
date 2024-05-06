package main

import "fmt"

//START OMIT
func Add(i, j int) int {
	return i + j
}

func AddAndSub(i, j int) (int, int) {
	return i + j, i - j
}

func AddSubAndDiv(i int, j int, k float64) (int, int, float64) {
	return i + j, i - j, float64(i) / k
}

func main() {
	a, b := AddAndSub(3, 5)
	fmt.Printf("%d %d\n", a, b)

	c, _ := AddAndSub(8, 2)
	fmt.Printf("%d\n", c)

	_, _, x := AddSubAndDiv(1, 3, 2)
	fmt.Printf("%f\n", x)
}
