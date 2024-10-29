package main

import "fmt"

func main() {
    var nilai int

    fmt.Print("Masukkan nilai ujian: ")
    fmt.Scanln(&nilai)

    if nilai >= 70 {
        fmt.Println("Lulus")
    } else {
        fmt.Println("Tidak Lulus")
    }
}
