package main

import (
	"fmt"
)

func main() {
	var beratParsial int // Berat parsial dalam gram
	var biayaTotal int   // Biaya pengiriman total

	// Input berat parsial dalam gram
	fmt.Print("Masukkan berat parsel (dalam gram): ")
	fmt.Scan(&beratParsial)

	// Hitung total berat dalam kilogram dan sisa berat dalam gram
	beratKg := beratParsial / 1000
	sisaGram := beratParsial % 1000

	// Biaya dasar per kg
	biayaKg := beratKg * 10000

	// Biaya tambahan berdasarkan sisa berat
	var biayaSisa int
	if beratKg > 10 {
		// Jika total berat lebih dari 10kg, sisa berat tidak dikenakan biaya
		biayaSisa = 0
	} else if sisaGram >= 500 {
		// Jika sisa berat lebih dari atau sama dengan 500 gram
		biayaSisa = sisaGram * 5
	} else {
		// Jika sisa berat kurang dari 500 gram
		biayaSisa = sisaGram * 15
	}

	// Hitung total biaya
	biayaTotal = biayaKg + biayaSisa

	// Output hasil perhitungan
	fmt.Printf("Total berat: %d kg dan %d gram\n", beratKg, sisaGram)
	fmt.Printf("Biaya pengiriman: Rp. %d\n", biayaTotal)
}