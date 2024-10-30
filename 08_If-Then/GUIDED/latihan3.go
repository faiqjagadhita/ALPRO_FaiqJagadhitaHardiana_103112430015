package main

import(
	"fmt"
)

func main() {
	var bil int // variabel bernama bil bertipe data integer
	var n bool // variabel bernama n betipe data boolean 

	fmt.Print("masukan bil bulat :") // meminta pengguna menginputkan bill bulat
	fmt.Scanln(&bil)

	if bil < 0 && bil % 2 == 0 { // jika bil lebih kecil dari 0 dan bil mod 2 == 0 maka true
		n = true
	}
	fmt.Println(n) // output n/boolean
}