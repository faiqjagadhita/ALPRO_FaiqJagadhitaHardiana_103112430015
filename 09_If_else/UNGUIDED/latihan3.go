package main

import "fmt"
import "math"

func main() {
	var b int

	// Input bilangan
	fmt.Print("Bilangan: ")
	fmt.Scan(&b)

	// Mencari dan menampilkan faktor
	fmt.Print("Faktor: ")
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()

	// Mengecek apakah bilangan prima
	isPrima := true
	if b <= 1 {
		isPrima = false
	} else {
		// Periksa hanya sampai akar kuadrat dari b untuk efisiensi
		for i := 2; i <= int(math.Sqrt(float64(b))); i++ {
			if b%i == 0 {
				isPrima = false
				break
			}
		}
	}

	// Output apakah bilangan prima
	if isPrima {
		fmt.Println("Prima: true")
	} else {
		fmt.Println("Prima: false")
	}
}