package main

import(
	"fmt"
)

func main() {
	var number int
	var continueloop bool
	for continueloop = true; continueloop; {
		fmt.Print("masukan bilangan :")
		fmt.Scan(&number)
		continueloop = number <= 0

	}
	fmt.Println(number, "adalah bilangan positif" )
}