package main

import(
	"fmt"
)

func main() {
	var n int 

	fmt.Print("masukan nilai n :")
	fmt.Scanln(&n)
	fmt.Print("jadi hasil kuadratnya adalah : ",n,":")

	for i := 1 ; i <= n ; i++ {
		fmt.Print(i*i," ")
	}
}