package main

import "fmt"

func main() {
	var n int // variabel bernama n bertipe data integer
	fmt.Print("Masukkan bilangan bulat non negatif: ") // input bil bulat non negatif 
	fmt.Scanln(&n)

	faktorial := 1
	// Loop untuk menghitung faktorial
	for i := 1; i <= n; i++ {
		faktorial *= i
	}

	fmt.Printf("Hasil faktorial dari %d adalah %d\n", n, faktorial)  // output
}