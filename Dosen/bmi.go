package main

import (
	"fmt"
)

func main() {
	var bb float32 
	var tb float32 
	var BMI float32 
	var nama string 

	fmt.Print("masukan nama :")
	fmt.Scanln(&nama)
	fmt.Print("masukan berat badanya :") 
	fmt.Scanln(&bb)
	fmt.Print("masukan tinggi badanya :") 
	fmt.Scanln(&tb)

	BMI = bb / (tb * tb) 

	fmt.Println("informasi BMi:")
	fmt.Println("nama :",nama)
	fmt.Println("berat :",bb,"kg")
	fmt.Println("tinggi :",tb,"m")
	fmt.Printf("BMI :%.2f\n" , BMI) 
}