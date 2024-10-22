package main

import (
	"fmt"
)

func main() {
	// Meminta pengguna memasukkan bilangan bulat positif
	var bilpositif int
	var total int

	fmt.Println("Masukkan bilangan bulat positif n: ")
	fmt.Scan(&bilpositif) // Membaca input dari pengguna

	// Jika n lebih besar dari 0, lakukan penjumlahan
	if bilpositif > 0 {
		total = 0 // Variabel untuk menyimpan total penjumlahan

		// Perulangan untuk menjumlahkan angka dari 1 hingga n
		for i := 1; i <= bilpositif; i++ {
			total = total + i // Menambahkan angka i ke total
		}

		// Menampilkan hasil penjumlahan
		fmt.Println("Jumlah deret dari 1 hingga", bilpositif, "adalah:", total)
	} else {
		// Jika input tidak valid, tampilkan pesan kesalahan
		fmt.Println("Masukkan bilangan bulat positif yang valid!")
	}
}