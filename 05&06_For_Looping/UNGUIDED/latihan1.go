package main

import "fmt"

func main() {
	var n int // variabel bernama n bertipedata integerr
	fmt.Print("Masukkan bilangan bulat positif: ") // Meminta pengguna untuk memasukkan bilangan bulat positif
	fmt.Scanln(&n) // Membaca input dari pengguna dan menyimpannya ke variabel 'n'

	sum := 0 // Inisialisasi variabel 'sum' dengan nilai 0 untuk menyimpan hasil penjumlahan
	for i := 1; i <= n; i++ { // Looping dari 1 hingga 'n'
		sum += i // Menambahkan nilai 'i' ke variabel 'sum' pada setiap iterasi loop
	}

	fmt.Printf("Hasil penjumlahan dari 1 sampai %d adalah: %d\n", n, sum) // Mencetak hasil penjumlahan
}