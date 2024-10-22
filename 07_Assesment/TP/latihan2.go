package main

import "fmt"

func main() {
	var jumlahBarang int // variabel bernama jumlahbarang bertipe data integer
	fmt.Print("Masukkan jumlah barang yang dibeli: ") // input pengguna jumlahbarang
	fmt.Scanln(&jumlahBarang)

	var totalPoin int = 0 // variabel bernama total poin yang bernilai 0 

	totalPoin = jumlahBarang * 10 // rumus menghitung poin normal

	// Menghitung poin tambahan
	if jumlahBarang > 5 { // penggulanga jika baranga lebih dari 5 maka akan mendapatkan poin tambahan sebesar 5 poin setiap barang
		totalPoin += (jumlahBarang - 5) * 5
	}

	fmt.Printf("Total poin yang didapatkan: %d poin\n", totalPoin) //output
}