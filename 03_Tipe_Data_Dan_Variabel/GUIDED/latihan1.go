package main

import(
	"fmt"
)

func main() {
	var s int 						// variabel bertipe integer yang bernama s atau sisi 
	var volume int 					// variabel bertipe integer yang bernama volume

	fmt.Print("masukan sisinya:") 	//menerima input sisi dari pengguna
	fmt.Scanln(&s)				  	//menerima input sisi dari pengguna

	volume = s * s * s 				//rumus volume kubus

	fmt.Println("volume kubus adalah ", volume) // menampilakan hasil output yaitu volume kubus
}
