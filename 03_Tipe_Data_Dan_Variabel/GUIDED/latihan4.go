package main

import (
	"fmt"
)

func main() {
	var rupiah float32 //varibel float bernama rupiah
	var hasil float32 //variabel float bernama hasil

	fmt.Print("masukan rupiahnya :") // input rupiah dari pengguna
	fmt.Scanln(&rupiah)

	hasil = rupiah/15.000 //rumus perubahan rupiah ke dolar 

	fmt.Println("hasil adalah", hasil ,"dolar") // hasil /output dari perubahan rupiah ke dolar

}