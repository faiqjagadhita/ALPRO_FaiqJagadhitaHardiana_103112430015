package main

import "fmt"

func main() {
	var harga, total int

	for {
		// Meminta input harga barang
		fmt.Print("Masukkan harga barang (ketik 0 untuk selesai): ")
		fmt.Scanln(&harga)

		// Periksa apakah input adalah 0
		if harga == 0 {
			break
		}
		// Tambahkan harga ke total
		total += harga
	}
	// Tampilkan total belanja
	fmt.Printf("Total belanja Anda: %d\n", total)
}