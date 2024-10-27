package main

import "fmt"

func main() {
	var n int
	var hasil bool

	fmt.Print("Masukkan bilangan tiga digit: ")
	fmt.Scan(&n)

	digit1 := n / 100       
	digit2 := (n / 10) % 10 
	digit3 := n % 10        

	hasil = (digit1 < digit2 && digit2 < digit3) || (digit1 > digit2 && digit2 > digit3)

	fmt.Println(hasil)
}