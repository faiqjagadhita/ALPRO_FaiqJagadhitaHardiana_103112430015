package main

import(
	"fmt"
)

func main() {
	var bilawal int // variabel bernama bilawal bertipedata integer
	var bilakhir int // variabel bernama bilakhir bertipedata integer
	var i int // variabel bernama i bertipedata integer

	fmt.Print("masukan bilangan awal :") // input user bilawal
	fmt.Scanln(&bilawal)
	fmt.Print("masukan bilangan akhir :") // input user bilakhir 
	fmt.Scanln(&bilakhir)

	for i = bilawal; i <= bilakhir; i++ { // pengulangan i = bilawal = 2  sampai i kurang dari sama dengan bilakhir 7 
		fmt.Print(i) // output hasil dari pengulangan i
	}

}