package main

import (
	"fmt"
)
func main (){
	var ph float64
	fmt.Print("masukan nilia ph anda :")
	fmt.Scan(&ph)

	switch {
	case ph >= 6.5 && ph <= 8.6 :
		fmt.Print("air layak minum")
	case ph >= 0 && ph < 6.5 || ph > 8.6 && ph <= 14 :
		fmt.Print("air tidak layak minum")
	case ph < 0 || ph > 14 :
		fmt.Print("Nilai pH tidak valid. Nilai pH harus antara 0 dan 14.")
	default:
		fmt.Print("masukan nilai ph harus angka ")
	}

}