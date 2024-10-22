package main

import(
	"fmt"
)

func main() {
	var s int
	var keliling int
	var luas int 
	//fmt. Scanln(&s)

	s = 27
	keliling = 4 * s
	luas = s * s

	fmt.Println("keliling alun-alun adalah", keliling , "meter ,dan luas alun-alun adalah", luas , "meter pesergi")
}