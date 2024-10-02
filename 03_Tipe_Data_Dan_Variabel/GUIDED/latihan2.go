package main

import(
	"fmt"
)

func main() {
	var alas,tinggi,hasil float32 // 3 variabel bertipe float yang memiliki nama alas ,tinggi ,dan hasil

	fmt.Print("masukan alasnya :") // menerima input alas dari pengguna
	fmt.Scanln(&alas)
	fmt.Print("masukan tingginya :")// menerima input tinggi dari pengguna
	fmt.Scanln(&tinggi)

	hasil = 0.5 * alas * tinggi // rumus luas segitiga

	fmt.Println("hasil volumenya adalah ", hasil)// hasil atau output dari luas segitiga
}