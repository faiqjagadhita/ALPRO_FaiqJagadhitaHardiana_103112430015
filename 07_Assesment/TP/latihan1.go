package main

import "fmt"

func main() {
	var n int // variabel bernama n bertipe data integer
	fmt.Print("Masukkan nilai N: ") // Meminta pengguna memasukkan nilai N
	fmt.Scanln(&n)                
	fmt.Print("Hasil kuadrat dari 1 sampai ", n, ": ") // Menampilkan teks sebelum mencetak hasil kuadrat

	for i := 1; i <= n; i++ { // Looping dari 1 hingga N
		fmt.Print(i*i, " ") // Mencetak kuadrat dari i, yaitu i * i, diikuti spasi
	}
	fmt.Println() //  output
}