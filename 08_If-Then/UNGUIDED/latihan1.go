package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah orang yang akan touring: ")
	fmt.Scan(&n)

	var motor int

	// Menggunakan IF untuk menentukan jumlah motor yang diperlukan
	if n%2 == 0 {
		// Jika jumlah n genap, bagi dengan 2
		motor = n / 2
	} else {
		// Jika jumlah n ganjil, bagi dengan 2 dan tambahkan 1
		motor = (n / 2) + 1
	}

	fmt.Printf("Jumlah motor yang diperlukan: %d\n", motor)
}