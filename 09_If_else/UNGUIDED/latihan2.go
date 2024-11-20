package main

import "fmt"

func main() {
	var nam float64
	var nmk string

	// Input nilai akhir mata kuliah
	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scan(&nam)

	// Menentukan nilai mata kuliah berdasarkan nilai akhir
	if nam > 80 {
		nmk = "A"
	} else if nam > 72.5 {
		nmk = "AB"
	} else if nam > 65 {
		nmk = "B"
	} else if nam > 57.5 {
		nmk = "BC"
	} else if nam > 50 {
		nmk = "C"
	} else if nam > 40 {
		nmk = "D"
	} else {
		nmk = "E"
	}

	// Menampilkan hasil
	fmt.Println("Nilai mata kuliah: ", nmk)
}