package main

import (
	"fmt"
)

func main() {
	var a, b, c int

	fmt.Print("Masukkan 3 bilangan (contoh: 1 2 3): ")
	fmt.Scan(&a, &b, &c)

	max := a
	min := a

	if b > max {
		max = b
	}
	if c > max {
		max = c
	}

	if b < min {
		min = b
	}
	if c < min {
		min = c
	}

	fmt.Printf("Bilangan maksimum dan minimum dari %d, %d, %d adalah: max = %d, min = %d\n", a, b, c, max, min)
}
