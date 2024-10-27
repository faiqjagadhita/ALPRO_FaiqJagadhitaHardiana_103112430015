package main

import "fmt"

func main() {
    var n int
    var str string

    fmt.Print("Masukkan jumlah pengulangan: ")
    fmt.Scanln(&n)

    fmt.Print("Masukkan string yang akan dicetak: ")
    fmt.Scanln(&str)

    for i := 0; i < n; i++ {
        fmt.Println(str)
    }
}
