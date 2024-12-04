package main

import "fmt"

func main() {
	// Angka rahasia
	secretNumber := 7
	var tebak int

	for {
		// Meminta input dari pengguna
		fmt.Print("Tebak angka (1-10): ")
		fmt.Scan(&tebak)

		// Mengecek apakah tebakan benar
		if tebak == secretNumber {
			fmt.Println("Selamat, tebakan Anda benar!")
			break
		} else {
			fmt.Println("Tebakan Anda salah, coba lagi.")
		}
	}
}