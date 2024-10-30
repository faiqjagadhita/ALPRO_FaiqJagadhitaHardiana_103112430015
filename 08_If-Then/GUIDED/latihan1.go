package main

import(
	"fmt"
)

func main() {
	var bil int // variabel bernama bil bertipe data integer

	fmt.Print("masukan bilangan bulat :")// meminta pengguna menginputkan bil bulat 
	fmt.Scanln(&bil)

	if bil < 0 { // jika bil kurang dari 0
		bil = -bil // maka bilangan dikalikan -
	
	}
	fmt.Println(bil) // jika lebih dari 0
}