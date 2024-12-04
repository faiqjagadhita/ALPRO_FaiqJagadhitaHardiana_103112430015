package main

import "fmt"

func main() {
	const usernameBenar = "Admin"
	const passwordBenar = "Admin"

	var username, password string
	gagalLogin := 0

	for {
		fmt.Print("Masukkan username: ")
		fmt.Scanln(&username)
		fmt.Print("Masukkan password: ")
		fmt.Scanln(&password)

		if username == usernameBenar && password == passwordBenar {
			fmt.Printf("\n%d percobaan gagal login\n", gagalLogin)
			break
		} else {
			fmt.Println("Username atau password salah. Coba lagi.")
			gagalLogin++
		}
	}
}