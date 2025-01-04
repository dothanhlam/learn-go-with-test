package structs_methods_interfaces


type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}


func Area(r Rectangle) float64 {
	return r.Width * r.Height
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

