package main

import(
	"fmt"
)

func main() {
	var input string
	var repetitions int

	fmt.Print("masukan kata dan angka :")
	fmt.Scanln(&input, &repetitions)
	counter := 0 
	for done := false; !done; {
		fmt.Println(input)
		counter++
		done = (counter >= repetitions)
	}

}