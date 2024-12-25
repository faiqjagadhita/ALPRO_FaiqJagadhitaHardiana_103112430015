package main

import "fmt"

func main() {
	var berat_kiri, berat_kanan, total_berat float64
	var oleng = false
	for total_berat < 150 {
		fmt.Printf("Masukkan berat belanjaan di kedua kantong: ")
		fmt.Scanln(&berat_kiri, &berat_kanan)
		selesih := berat_kiri - berat_kanan
		if selesih < 0 {
			selesih = -selesih 
		}
		oleng = selesih >= 9
		total_berat = berat_kiri + berat_kanan
		if total_berat < 150 {
			fmt.Println("Sepeda motor pak Andi akan oleng: ", oleng)
		}
	}
	fmt.Println("Program selesai")
}