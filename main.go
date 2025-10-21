package main

import "fmt"

func area(r float32) float32{
	const phi = 3.14
	return phi * r * r
}

func around(r float32) float32{
	const phi = 3.14
	return 2*phi*r
}

func main(){
	fmt.Println("Hasil luas lingkaran dengan r: 10 = ", area(10))
	fmt.Println("Hasil keliling lingkaran dengan r: 10 = ", around(10))
}