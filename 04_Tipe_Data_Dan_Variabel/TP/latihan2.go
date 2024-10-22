// gaji bulanan
package main

import "fmt"

func main() {
    var jks1,jks2,jks3,jks4 float64
    var upahPerJam float64
    var jl1,jl2,jl3,jl4 float64

    fmt.Print("Masukkan jumlah jam kerja dalam seminggu: ")
    fmt.Scanln(&jks1, &jks2, &jks3, &jks4)

    fmt.Print("Masukkan upah per jam: Rp ")
    fmt.Scanln(&upahPerJam)

    if jks1  > 40 {
        jl1 = jks1 - 40
    } 
	if jks2  > 40 {
        jl2 = jks2 - 40
    } 
	if jks3  > 40 {
        jl3 = jks3 - 40
    } 
	if jks4  > 40 {
        jl4 = jks4 - 40
    }


    gajiNormal := 40 * upahPerJam
    gajiLembur1 := jl1 * 1.5 * upahPerJam
	gajiLembur2 := jl2 * 1.5 * upahPerJam
	gajiLembur3 := jl3 * 1.5 * upahPerJam
	gajiLembur4 := jl4 * 1.5 * upahPerJam

    totalGaji1 := gajiNormal + gajiLembur1
	totalGaji2 := gajiNormal + gajiLembur2
	totalGaji3 := gajiNormal + gajiLembur3
	totalGaji4 := gajiNormal + gajiLembur4

    totalGajiBulanan := totalGaji1 + totalGaji2 + totalGaji3 + totalGaji4

    fmt.Printf("Total gaji bulanan: Rp %.2f\n", totalGajiBulanan)
}
