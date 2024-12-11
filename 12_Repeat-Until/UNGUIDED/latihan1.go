package main

import "fmt"

func main() {
	var number int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&number)

	if number <= 0 {
		fmt.Println("Masukkan bilangan bulat positif.")
		return
	}

	digitCount := 0
	for number > 0 {
		number /= 10 
		digitCount++ 
	}

	fmt.Printf("Banyaknya digit adalah %d\n", digitCount)
}