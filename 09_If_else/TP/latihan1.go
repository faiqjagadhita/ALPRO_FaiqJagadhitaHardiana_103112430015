package main

import "fmt"

func main() {
	var nilai int
	fmt.Print("Masukkan nilai mahasiswa: ")
	fmt.Scan(&nilai)

	if nilai > 90 {
		fmt.Println("Indeks: A")
	} else if nilai >= 80 {
		fmt.Println("Indeks: AB")
	} else if nilai >= 70 {
		fmt.Println("Indeks: B")
	} else {
		fmt.Println("Indeks: C")
	}
}