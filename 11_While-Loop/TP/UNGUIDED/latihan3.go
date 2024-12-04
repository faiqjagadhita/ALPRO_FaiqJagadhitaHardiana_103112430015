package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukkan dua bilangan (x y): ")
	fmt.Scanln(&x, &y)

	result := 0

	for x >= y {
		x -= y
		result++
	}

	fmt.Println("Hasil:", result)
}