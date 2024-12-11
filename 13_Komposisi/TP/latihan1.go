package main

import "fmt"

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var input int

	fmt.Print("Masukkan sebuah bilangan bulat: ")
	fmt.Scan(&input)

	fmt.Printf("Bilangan prima dari 1 hingga %d adalah:\n", input)
	for i := 1; i <= input; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}