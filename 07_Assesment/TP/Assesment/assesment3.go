package main

import(
	"fmt"
)

func main() {
	var dinar int
	var dirham int
	var fals int 
	var qirat int 
	

	fmt.Print("masukan uang dalam satuan qirat :")
	fmt.Scanln(&qirat)

	dinar = qirat/600
	dirham = qirat %600/60
	fals = qirat %600 % 60 /6
	qirat =qirat % 600 % 60 % 6 

	fmt.Print( dinar,dirham,fals,qirat)

}

