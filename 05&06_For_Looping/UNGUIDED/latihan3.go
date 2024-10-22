package main

import "fmt"

func main() {
	var basis, pangkat int // variabel bernama basis dan pangkat bertipe data integer
	fmt.Print("Masukkan basis: ")
	fmt.Scanln(&basis) // Meminta input dari pengguna untuk basis
	fmt.Print("Masukkan pangkat: ")
	fmt.Scanln(&pangkat) // Meminta input dari pengguna untuk pangkat

	hasil := 1 // Inisialisasi hasil dengan 1
	for i := 0; i < pangkat; i++ { // Looping sebanyak pangkat kali
		hasil *= basis // Mengalikan hasil dengan basis pada setiap iterasi
	}
	fmt.Println("Hasil pemangkatan:", hasil) // Mencetak hasil pemangkatan
}