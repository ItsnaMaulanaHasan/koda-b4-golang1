package Circle

func (fn Circle) Circumference() float32{
	const phi = 3.14
	return 2 * phi * fn.Radius
}