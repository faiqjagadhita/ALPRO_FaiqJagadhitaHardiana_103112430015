package main

import (
	"fmt"
)
func main() {
	var token string
	fmt.Scan(&token)
	for token != "12345abcde" {
		fmt.Scan(&token)
}
fmt.Println("Selamat Anda berhasil login")
}