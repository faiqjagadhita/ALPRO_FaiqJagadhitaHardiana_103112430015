package main

import "fmt"

func main() {
	var number float64
	fmt.Print("Masukkan bilangan desimal: ")
	fmt.Scan(&number)

	roundedUp := float64(int(number))
	if number > roundedUp {
		roundedUp++
	}

	for number < roundedUp {
		number += 0.1
		fmt.Printf("%.1f\n", number)
	}
}