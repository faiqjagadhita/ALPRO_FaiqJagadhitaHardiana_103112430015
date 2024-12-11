package main

import "fmt"

func isPerfectNumber(num int) bool {
	if num <= 0 {
		return false
	}
	sum := 0
	for i := 1; i < num; i++ {
		if num%i == 0 {
			sum += i
		}
	}
	return sum == num
}

func main() {
	var input int
	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scan(&input)

	if isPerfectNumber(input) {
		fmt.Printf("Ya (karena faktor dari %d adalah bilangan yang jumlahnya sama dengan %d)\n", input, input)
	} else {
		fmt.Printf("Tidak (bukan bilangan sempurna)\n")
	}
}