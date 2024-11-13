package main

import (
	"fmt"
)

func main() {
	var bil int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bil)

	if bil < 0 && bil%2 == 0 {
		fmt.Println("genap negatif")
	} else {
		fmt.Println("bukan")
	}
}