package main

import(
	"fmt"
)

func main() {
	var fahrenheit float64
	var kelvin float64
	
	fmt.Print("masukan suhu dalam derajat Fahrenheit: ")
	fmt.Scanln(&fahrenheit)

	kelvin = (fahrenheit - 32) * 5 / 9 + 273.15

	fmt.Println("hasil ", fahrenheit ,"fahrenheit menjadi ",kelvin, "kelvin")
}