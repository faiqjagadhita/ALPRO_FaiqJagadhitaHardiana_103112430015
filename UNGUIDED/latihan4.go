package main

import "fmt"

func main() {
	var i, banyak_bunga int
	var bunga, pita string
	i = 1
	banyak_bunga = 0
	for bunga != "SELESAI" {
		fmt.Printf("Bunga %d: ", i)
		fmt.Scan(&bunga)
		if bunga != "SELESAI" {
			pita += bunga + " - "
			i++
			banyak_bunga++
		}
	}
	fmt.Println("Pita: ", pita)
	fmt.Print("Bunga: ", banyak_bunga)
}