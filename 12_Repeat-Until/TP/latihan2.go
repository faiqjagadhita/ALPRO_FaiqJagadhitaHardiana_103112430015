package main

import "fmt"

func main() {
	var input string
	for {
		// Meminta input pengguna
		fmt.Print("Masukkan kata: ")
		fmt.Scanln(&input)

		// Periksa apakah input adalah "telkom"
		if input == "telkom" {
			fmt.Println("Program selesai.")
			break
		}

		// Menampilkan input pengguna
		fmt.Println("Anda mengetik:", input)
	}
}