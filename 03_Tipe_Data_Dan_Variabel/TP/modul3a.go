package main

import (
	"fmt"
)

func main() {
	var angka1 float64
	var angka2 float64
	var hasil float64
	var operator string

	fmt.Print("masukan angka pertama : ")
	fmt.Scanln(&angka1)

	fmt.Print("masukan operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Print("masukan angka kedua : ")
	fmt.Scanln(&angka2)

	switch operator{
	
		case "+":
			hasil = angka1 + angka2
		case "-":
			hasil = angka1 - angka2
		case "*":
			hasil = angka1 * angka2
		case "/":
	}

	fmt.Println("hasil dari ", angka1, operator, angka2,"=" ,hasil)

}