package main

import (
	"fmt"
)

func main() {
	var total float64
	for { 
		var barang string
		fmt.Print("Masukkan nama barang (atau ketik 'selesai' untuk keluar): ")
		fmt.Scanln(&barang)

		if barang == "selesai" {
			break
		}

		var harga float64
		fmt.Print("Masukkan harga barang: ")
		fmt.Scanln(&harga)

		total += harga
	}

	fmt.Printf("Total belanja: %.2f\n", total)
	fmt.Println("Transaksi selesai. Terima kasih!")
}