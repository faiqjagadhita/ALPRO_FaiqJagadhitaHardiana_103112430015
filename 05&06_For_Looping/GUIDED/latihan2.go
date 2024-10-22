package main

import(
	"fmt"
)

func main(){
	var alas int // variabel yang bernama alas bertipe data integer
	var tinggi int // variabel yang bernama tinggi bertipedata integer
	var n int // variabel bernama n bertipe data integer

	fmt.Print("masukan n :") //input n dari pengguna atau nilaii pengulangan
	fmt.Scanln(&n)

	for i := 1 ; i <= n ; i++ { // nilai pengulangan i = 1 ke i <= inputan nilai penggulangan 
		fmt.Print("masukan alasnya :") // input alasnya
		fmt.Scanln(&alas)
		fmt.Print("masukan tingginya :") // input tingginya
		fmt.Scanln(&tinggi)
		
		luas := alas * tinggi / 2 // rumus luas segita
		fmt.Println("luas",luas) // output 
	}

}