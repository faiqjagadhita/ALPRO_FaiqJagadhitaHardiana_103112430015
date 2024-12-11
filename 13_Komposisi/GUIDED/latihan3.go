package main

import "fmt"

func main() {
	var n int

	fmt.Print("Masukkan sebuah bilangan bulat positif: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Harap masukkan bilangan bulat positif.")
		return
	}

	for i := 1; i <= n; i++ {
		if n%i == 0 { 
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}
