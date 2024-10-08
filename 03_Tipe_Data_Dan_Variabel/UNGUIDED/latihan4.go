package main

import(
	"fmt"
)

func main() {
	var x, fx int // 2 variabel bertipe integer yang bernama x dan juga fx

	fmt.Print("masukan X nya :")// input x dari pengguna
	fmt.Scanln(&x)

	fx = (2/(x+5)+5)// rumus Fx 

	fmt.Print("jadi fx nya adalah ", fx) // hasil /output dari f(x)

}