package main

import (
	"fmt"
)

func main() {
	var jam int // variabel yang bernama jam dengan tipe data intger
	var menit int // variabel yang bernama menit dengan tipe data interger
	var detik int // variabel yang bernama detik dengan tipe data interger
	var detikKeseluruannya int // variabel yang bernama derikKeseluruanya interger

	fmt.Print("masukan nilai detik keseluruanya :") // meminta inputan keseleluruan detiknya kepada pengguna
	fmt.Scanln(&detikKeseluruannya)

	jam = detikKeseluruannya / 3600 // rumus menghitungg jam dari detik
	menit = (detikKeseluruannya % 3600) / 60 // rumus menghitung menit dari detik
	detik = detikKeseluruannya % 60 // rumus menghitung sisa detik 

	fmt.Println("jadi nilai dari detik keseluruan yaitu : ",detikKeseluruannya ,"menjadi :",jam ,"jam",menit ,"menit",detik ,"detik")
	// hasil output dari program 
} 