package main

import (
	"fmt"
)

func main() {
	var usia int
	var kk bool

	fmt.Print("Masukan usia anda: ")
	fmt.Scan(&usia)

	fmt.Print("Apakah anda mempunyai KK? (true/false): ")
	fmt.Scan(&kk)

	if usia >= 17 && kk == true{
		fmt.Print("Bisa membuat KTP")
	} else{
		if usia <17 {
			fmt.Print("Belum bisa membuat KTP")
		}
		if kk == false{
			fmt.Print("belum bisa mebuat KTP")
		}
	}
}