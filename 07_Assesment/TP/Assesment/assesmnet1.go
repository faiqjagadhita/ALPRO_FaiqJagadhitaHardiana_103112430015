package main

import(
	"fmt"
)

func main() {
	var x int 
	var y int
	var nilai int

	fmt.Print("masukan nilai x :")
	fmt.Scanln(&x)
	fmt.Print("masukan nilai y :")
	fmt.Scanln(&y)

	for i := x ; i <= y ; i++ {
		nilai += i
	
	}
	fmt.Print(nilai)
	
}