package main

import(
	"fmt"
)

func main() {
	var totalBelanja1 int // variabel yang bernama totalbelanja1 yang bertipe data integer
	var diskon int // variabel yang bernama diskon yang bertipe data integer
	var hasilDiskon int // variabel yang bernama hasildiskon yang beripe data integer
	var totalBelanja2 int // variabel yang bernama totalbelanja2 yang bertipe data interger

	fmt.Print("masukan total belanja awal : ") // meminta pengguna menginputkan total belanja awal
	fmt.Scanln(&totalBelanja1)
	fmt.Print("masukan jumlah diskon : ") // meminta pengguna menginputkan jumlah % diskon 
	fmt.Scanln(&diskon)

	hasilDiskon = (totalBelanja1 * diskon) / 100 // rumus diskon yang di dapat
	totalBelanja2 = totalBelanja1 - hasilDiskon // rumus total belanja yang sudah di potong diskon

	fmt.Print("jadi total belanja setelah terpotong diskon adalah :",totalBelanja2)
	//output yang di hasilkan dari totalbelanja2
}
