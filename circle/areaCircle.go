package circle

func (fn Circle) Area() float32{
	const phi = 3.14
	return phi * fn.Radius * fn.Radius
}