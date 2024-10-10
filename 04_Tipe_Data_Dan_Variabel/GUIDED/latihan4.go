package main

import (
	"fmt"
)

func main() {
	var bilangan int // variabel yang bernama bilangan dengan tipedata integer
	var d1 int // variabel yang bernama d1 dengan tipedata integer
	var d2 int // variabel yang bernama d2 dengan tipedata integerr
	var d3 int // variabel yang bernama d3 dengan tipedata integerr

	fmt.Print("masukan bilangan tiga digit :") // meminta pengguna menginputkan bilangan  3 digit
	fmt.Scanln(&bilangan)

	d1 = bilangan / 100 // rumus menghitung digit 1
	d2 = (bilangan % 100) / 10 // rumus menghitung d2
	d3 = bilangan % 10 // rumus menghitung d3

	fmt.Println(d1 <= d2 && d2 <= d3) // hasilnya karena ada operator boolean yaitu && jadi output yang di kediluarkan jadi variabel langsung membaca itu boolean tanpa harus adanya deklarasi boolean di awal
}