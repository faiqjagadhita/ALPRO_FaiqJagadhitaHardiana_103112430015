package main

import "fmt"

func main() {
	var n int
	var prima bool
	prima = true
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)
	for j := 2; j*j <= n; j++ {
		if n%j == 0 {
			prima = false
		}
	}
	if prima {
		fmt.Print("Bilangan Prima")
	} else {
		fmt.Print("Bukan Bilangan Prima")
	}
}