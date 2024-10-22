package main

import (
	"fmt"
)

func main() {
	// Menampilkan bilangan genap dari 1 hingga 50
	fmt.Println("Bilangan genap dari 1 hingga 50 adalah:")

	// Perulangan dari 1 hingga 50
	for i := 1; i <= 50; i++ {
		// Memeriksa apakah i adalah bilangan genap
		if i%2 == 0 {
			fmt.Println(i) // Jika genap, cetak angka
		}
	}
}