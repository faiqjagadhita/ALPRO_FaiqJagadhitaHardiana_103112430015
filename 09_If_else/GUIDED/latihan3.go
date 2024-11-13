package main

import (
	"fmt"
)

func main() {
	var angka int

	// Meminta pengguna memasukkan bilangan 4 digit
	fmt.Print("Masukkan bilangan 4 digit : ")
	fmt.Scan(&angka)

	// Memeriksa apakah bilangan memenuhi syarat 4 digit
	if angka < 1000 || angka > 9999 {
		fmt.Println("Bilangan harus 4 digit dan berada antara 1000 hingga 9999.")
		return
	}	

	// Mengambil setiap digit dari bilangan
	digit1 := angka / 1000            // ribuan
	digit2 := (angka / 100) % 10      // ratusan
	digit3 := (angka / 10) % 10       // puluhan
	digit4 := angka % 10              // satuan
	

	// Mengecek pola urutan digit
	if digit1 < digit2 && digit2 < digit3 && digit3 < digit4 {
		fmt.Println("Digit terurut membesar")
	} else if digit1 > digit2 && digit2 > digit3 && digit3 > digit4 {
		fmt.Println("Digit terurut mengecil")
	} else {
		fmt.Println("Digit tidak terurut")
	}
}
