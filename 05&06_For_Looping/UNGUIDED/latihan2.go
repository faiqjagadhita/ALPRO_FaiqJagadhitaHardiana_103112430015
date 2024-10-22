package main

import (
	"fmt"
)

func main() {
	var n int // variabel bernama n bertipe data integer
	fmt.Scanln(&n) // Membaca jumlah kerucut (n) dari input

	var r, h float64 // variabel bernama r dan h bertipe data float
	volume := make([]float64, n) // Membuat slice untuk menyimpan volume setiap kerucut

	for i := 0; i < n; i++ {
		fmt.Scanln(&r, &h) // Membaca jari-jari (r) dan tinggi (h) dari input
		volume[i] = (1.0/3.0)*3.14159*r*r*h // Menghitung volume dengan rumus (1/3) * π * r^2 * h
	}

	for _, v := range volume {
		fmt.Println(v) // Mencetak volume setiap kerucut
	}
}