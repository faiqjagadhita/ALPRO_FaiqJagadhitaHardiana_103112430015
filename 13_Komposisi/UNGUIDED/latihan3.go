package main

import (
	"fmt"
)

func main() {
	var w1, w2, w3, w4 string
	var Urut = true
	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d : ", i)
		fmt.Scanln(&w1, &w2, &w3, &w4)
		if w1 != "merah" || w2 != "kuning" || w3 != "hijau" || w4 != "ungu" {
			Urut = false
		}
	}

	if Urut {
		fmt.Println("BERHASIL: True")
	} else {
		fmt.Println("BERHASIL: False")
	}
}