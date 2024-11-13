package main

import "fmt"

func main() {
	var huruf string
	fmt.Print("Masukkan sebuah huruf: ")
	fmt.Scan(&huruf)

	if huruf == "A" || huruf == "I" || huruf == "U" || huruf == "E" || huruf == "O" ||
		huruf == "a" || huruf == "i" || huruf == "u" || huruf == "e" || huruf == "o" {
		fmt.Println("Huruf Vokal")
	} else {
		fmt.Println("Huruf Konsonan")
	}
}