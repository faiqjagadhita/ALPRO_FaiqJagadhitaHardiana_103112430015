package main

import(
	"fmt"
)

func main() {
	var celsius float64 // variabel ini bernama celsius yang menggunakan tipedata float32
	var fahrenheit float64 // variabel ini bernama fahrenheit yang menggunakan tipedata float32 
	var reamur float64 // variabel ini bernama reamur yang menggunakan tipedata float32
	var kelvin float64 // variabel ini bernama kelvin yang menggunakan tipedata float32

	fmt.Print("masukan temperatur celsiusnya : ") // input temperatur dari pengguna 
	fmt.Scanln(&celsius)

	fahrenheit = (9.0 / 5.0) * celsius + 32 // rumus fahrenheit
	reamur = celsius * (4.0 / 5.0) // rumus reamur 
	kelvin = (fahrenheit + 459.67) * (5.0 / 9.0) // rumus kelvin

	fmt.Println("temperatur celsius : ", celsius)// output celsius
	fmt.Println("suhu fahrenheit : ", fahrenheit)// output fahrenheit
	fmt.Println("suhu reamur : ", reamur)// output reamur
	fmt.Println("suhu kelvin : ", kelvin)// output kelvin
}