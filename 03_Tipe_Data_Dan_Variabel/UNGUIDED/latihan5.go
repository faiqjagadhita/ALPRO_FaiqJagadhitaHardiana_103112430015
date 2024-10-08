package main

import(
	"fmt"
)

func main() {
	var r float64 //variabel yang bernama r yang berarti jari jari
	var pi float64 // varibel yang bernama pi 
	var volume float64 // variabel yang bernama volume 
	var luas float64 // variabel yang bernama luas

	fmt.Print("masukan r atau jari-jarinya : ") // input yang r yang diberikan oleh pengguna
	fmt.Scanln(&r)
	
	pi = 3.1415926535 // nilai dari variabel pi

	volume = (4.0 / 3.0) * pi * (r * r * r) // rumus volume bola
	luas = 4.0 * pi * (r * r ) // rumus luas bola

	fmt.Println("jadi hasinya memiliki volume yaitu :", volume ,"dan luas :" , luas )//  output yang di hasilkan dari penghitungan 2 rumus tsb

}