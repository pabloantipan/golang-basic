package structs

type Car struct {
	Brand string
	Year  int
}

func Structing() interface{} {
	myCar := Car{
		Brand: "Porsche",
		Year:  2022,
	}

	return myCar
}
