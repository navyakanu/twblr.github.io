package interfaces

type Shape interface {
	Area() int
}

type Square struct {
	side int
}

type Rectangle struct {
	length  int
	breadth int
}

type Hybrid struct {
	square    Square
	rectangle Rectangle
}

func (r Rectangle) Area() int {
	return r.length * r.breadth
}

func (s Square) Area() int {
	return s.side * s.side
}

func (h Hybrid) Area() int {

	return (h.square.Area()) + (h.rectangle.Area())

}
