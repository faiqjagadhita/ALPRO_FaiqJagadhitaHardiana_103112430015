package main

import "fmt"

func main() {
	var angka, tbg int
	fmt.Print("Masukkan bilangan bulat positif n: ")
	fmt.Scan(&angka)
	for i := 1; i <= angka; i++ {
		if i%2 != 0 {
			tbg++
		}
	}
	fmt.Printf("Terdapat %d bilangan ganjil", tbg)
}