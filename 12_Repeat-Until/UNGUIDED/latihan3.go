package main

import "fmt"

func main() {
	var target int
	fmt.Print("Masukkan target donasi: ")
	fmt.Scan(&target)

	var totalDonasi, jumlahDonatur int

	for {
		var donasi int
		fmt.Print("Masukkan donasi dari donatur: ")
		fmt.Scan(&donasi)

		totalDonasi += donasi
		jumlahDonatur++

		fmt.Printf("Total donasi: %d, Jumlah donatur: %d\n", totalDonasi, jumlahDonatur)

		if totalDonasi >= target {
			break
		}
	}

	fmt.Println("Target donasi tercapai!")
	fmt.Printf("Total donasi akhir: %d\n", totalDonasi)
	fmt.Printf("Jumlah donatur: %d\n", jumlahDonatur)
}