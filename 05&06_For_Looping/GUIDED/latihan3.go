package main

import(
	"fmt"
)

func main() {
	var b1 int // varabel bernama b1 beripe data integer
	var b2 int // varabel bernama b2 beripe data integer
	var hasil int = 0 // varabel bernama hasil beripe data integer

	fmt.Print("masukan bilangan 1 :") // input bill 1
	fmt.Scanln(&b1)
	fmt.Print("masukan bilangan 2 :") // input bill 2
	fmt.Scanln(&b2)

	for i := 1 ; i <= b2 ; i++ { // pengulangan dari 1 = 1 menuju i <= b2 
		hasil = hasil + b1 // rumus 
	}
	fmt.Print(hasil) // output
}