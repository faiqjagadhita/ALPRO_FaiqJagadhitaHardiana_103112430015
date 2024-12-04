package main

import (
	"fmt"
)
func main() {
	var N, s1, s2, j, temp int
	fmt.Scan(&N)
	s1 = 0
	s2 = 1
	j = 0
	for j < N {
		fmt.Print(s1," ")
		temp = s1 + s2
		s1 = s2
		s2 = temp
		j = j + 1
	}
}