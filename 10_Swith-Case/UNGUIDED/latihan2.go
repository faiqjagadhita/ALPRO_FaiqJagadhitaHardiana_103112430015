package main

import "fmt"

func main() {
	var kendaraan string
	var durasi int
	var tarif int

	fmt.Print("Masukkan jenis kendaraan (Motor/Mobil/Truk): ")
	fmt.Scan(&kendaraan)
	fmt.Print("Masukkan durasi parkir (dalam jam): ")
	fmt.Scan(&durasi)

	switch {
	case kendaraan == "Motor":
		tarif = 2000 * durasi
		fmt.Print("Rp", tarif)
	case kendaraan == "Mobil":
		tarif = 5000 * durasi
		fmt.Print("Rp", tarif)
	case kendaraan == "Truk":
		tarif = 8000 * durasi
		fmt.Print("Rp", tarif)
	default:
		fmt.Println("Jenis kendaraan atau durasi tidak valid")
	}
}