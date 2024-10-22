package main

import "fmt"

func main() {
	// Angka target ditentukan langsung sebagai 42 untuk kesederhanaan
	target := 40
	var tebakan int

	fmt.Println("Tebak angka antara 1 hingga 100. Anda punya 5 kali kesempatan.")

	// Loop untuk 5 kesempatan
	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d: Masukkan tebakan Anda: ", i)
		fmt.Scan(&tebakan)

		// Cek apakah tebakan benar
		if tebakan == target {
			fmt.Println("Selamat! Tebakan Anda benar.")
			return
		} else if tebakan < target {
			fmt.Println("Terlalu kecil!")
		} else {
			fmt.Println("Terlalu besar!")
		}
	}

	// Jika tidak berhasil dalam 5 kali percobaan
	fmt.Println("Maaf, Anda kehabisan kesempatan. Angka yang benar adalah:", target)
}