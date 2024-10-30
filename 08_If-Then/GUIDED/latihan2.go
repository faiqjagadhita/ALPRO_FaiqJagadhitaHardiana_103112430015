package main

import (
	"fmt"
)

func main() {
	var bil int // variabel bernama bil yang bertipe data integer

	fmt.Print("masukan bilangan bulat :") // meminta pengguna menginputkan bill bulat
	fmt.Scanln(&bil)

	if bil > 1 { // jika bil lebih dari 1 
		fmt.Println("positif") // maka bill tsb bernilai positif
		return 
	}

	fmt.Println("negatif") // jika tidak maka akan negatif 
}
