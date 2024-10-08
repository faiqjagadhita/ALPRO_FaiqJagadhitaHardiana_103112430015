package main

import(
	"fmt"
)

func main() {
	var tahun uint32 // variabel yang memiliki nama variabel tahun bertipedata integer
	var kabisat bool // varabel yang memilki nama kabisat bertipe data boolean karena hasilnya akan untuk menentukan benar salahnya

	fmt.Print("masukan tahunya : ")// inputan tahun dari pengguna
	fmt.Scanln(&tahun)

	kabisat = tahun % 400 == 0
	kabisat = tahun % 100 != 0 //rumus untuk menghitung tahun kabisat yang menggunakan tipedata boolean yang memiliki 2 nilai saja 
	kabisat = tahun % 4 == 0

	if kabisat {
		fmt.Println(kabisat , "adalah tahun kabisat.")// output jika benar tahun kabisat
	} else {
		fmt.Println(kabisat , "bukan tahun kabisat")// output jika bukan tahun kabisat
	}
}