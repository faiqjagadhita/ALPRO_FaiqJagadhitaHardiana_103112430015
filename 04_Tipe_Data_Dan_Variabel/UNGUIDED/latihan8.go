package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2, x3, y3 float64

	// Input koordinat titik A, B, dan C
	fmt.Println("Masukkan koordinat titik A, B, dan C (x1 y1 x2 y2 x3 y3):")
	fmt.Scan(&x1, &y1, &x2, &y2, &x3, &y3)


	// Fungsi untuk menghitung jarak antar dua titik
	distance := func(x1, y1, x2, y2 float64) float64 {
		return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	}

	// Hitung panjang sisi-sisi segitiga
	AB := distance(x1, y1, x2, y2) // Jarak antara A dan B
	BC := distance(x2, y2, x3, y3) // Jarak antara B dan C
	AC := distance(x1, y1, x3, y3) // Jarak antara A dan C

	// Tentukan sisi terpanjang
	longest := AB
	if BC > longest {
		longest = BC
	}
	if AC > longest {
		longest = AC
	}
	

	// Tampilkan hasil dengan 2 angka di belakang koma
	fmt.Printf("%.2f\n", longest)
}