package main

import (
	"fmt"
)

func main() {
	var bb float32 // variabel yang bernama bb dengan tipedata float
	var tb float32 // variabel yang bernama tb dengan tipedat float
	var BMI float32 // variabel yang bernama BMI dengan tipedata float

	fmt.Print("masukan BMI nya :") // meminta pengguna menginputkan berat BMInya
	fmt.Scanln(&BMI)
	fmt.Print("masukan tinggi badanya :") // meminta pengguna menginputkan tinggi badan 
	fmt.Scanln(&tb)

	bb = BMI * (tb * tb) + 1 // rumus bb

	fmt.Println("jadi BMInya adalah :" , int(bb)) // hasil output dari rumus bb
}