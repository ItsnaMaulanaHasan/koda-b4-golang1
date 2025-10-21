package main

import "fmt"

func area(r float32) float32{
	const phi = 3.14
	return phi * r * r
}

func circumference(r float32) float32{
	const phi = 3.14
	return 2*phi*r
}

func main(){
	var r float32
	fmt.Print("Enter the radius of the circle: ")
	fmt.Scan(&r)
	fmt.Println("The result of the area of a circle with a radius of", r, "=", area(r))
	fmt.Println("The result of the circumference of a circle with a radius of", r, "=", circumference(r))
}