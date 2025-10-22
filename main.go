package main

import (
	"fmt"
	circle "go-practice/circle"
)

func main(){
	var r float32
	fmt.Print("Enter the radius of the circle: ")
	fmt.Scan(&r)

	circle := circle.Circle{
		Radius: r,
	}
	fmt.Println("The result of the area of a circle with a radius of", r, "=", circle.Area())
	fmt.Println("The result of the circumference of a circle with a radius of", r, "=", circle.Circumference())
}