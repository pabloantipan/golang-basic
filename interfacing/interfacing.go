package interfacing

type Figure2D interface {
	Area() float64
}

type Square struct {
	Base float64
}

type Rectangle struct {
	Base   float64
	Height float64
}

func (square Square) Area() float64 {
	return square.Base * square.Base
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Base * rectangle.Height
}

func Calculate(figure Figure2D) float64 {
	return figure.Area()
}

func ShowInterfaceList() []interface{} {
	myInterface := []interface{}{"Hey!", 12, 4.96}
	return myInterface
}
