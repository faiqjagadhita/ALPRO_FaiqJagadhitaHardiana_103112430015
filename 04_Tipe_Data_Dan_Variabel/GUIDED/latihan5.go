package main

import (
	"fmt"
)

func main() {
	var bb float32 // variabel yang bernama bb dengan tipedata float
	var tb float32 // variabel yang bernama tb dengan tipedat float
	var BMI float32 // variabel yang bernama BMI dengan tipedata float

	fmt.Print("masukan berat badanya :") // meminta pengguna menginputkan berat badanya
	fmt.Scanln(&bb)
	fmt.Print("masukan tinggi badanya :") // meminta pengguna menginputkan tinggi badan 
	fmt.Scanln(&tb)

	BMI = bb / (tb * tb) // rumus BMI 

	fmt.Printf("jadi BMInya adalah :%.2f\n" , BMI) // hasil output dari rumus BMI
}