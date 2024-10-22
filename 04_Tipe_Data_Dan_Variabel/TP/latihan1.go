package main

import(
	"fmt"
)

func main() {
	var r float64 
	var pi float64 
	var keliling float64  
	var luas float64 

	fmt.Print("masukan r atau jari-jarinya : ") 
	fmt.Scanln(&r)
	
	pi = 3.14 

	keliling = 2 * pi * r  
	luas = pi * (r * r ) 

	fmt.Println("jadi hasinya memiliki keliling yaitu :", keliling ,"dan luas :" , luas )

}