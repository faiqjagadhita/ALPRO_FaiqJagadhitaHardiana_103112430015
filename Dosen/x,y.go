package main

import(
	"fmt"
)

func main() {
	var x float32
	var y float32
	var hasil float32

	fmt.Print("masukan x : ")
	fmt.Scanln(&x)
	fmt.Print("masukan y : ")
	fmt.Scanln(&y)

	hasil = (1 / ((3 * (x*x)) + 10)) + 10 * y + 7
	fmt.Println("hasil F(x,y) adalah ", hasil)
}