package main

import (
	"fmt"
)

func main() {
	const ktBenar = "123" 
	const maxPercobaan = 3          

	var kataSandi string
	var jumlahPercobaan int

	for jumlahPercobaan < maxPercobaan {
		fmt.Print("Masukkan kata sandi Anda: ")
		fmt.Scan(&kataSandi)

		if kataSandi == ktBenar {
			fmt.Println("Login berhasil!")
			return
		} else {
			jumlahPercobaan++
			fmt.Printf("Kata sandi salah. Percobaan tersisa: %d\n", maxPercobaan-jumlahPercobaan)
		}
	}

	fmt.Println("Login ditolak")
}
