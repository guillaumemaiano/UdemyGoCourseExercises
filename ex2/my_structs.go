package main

type square struct {
	sideLength float64;
}


type triangle struct {
	height float64;
	base float64;
}

//func (triangle) getArea (base float64, height float64) float64 {
func (tr triangle) getArea () float64 {
	return tr.base * tr.height / 2;
}

//func (square) getArea (side float64 ) float64 {
func (sq square) getArea () float64 {
	return sq.sideLength * sq.sideLength;
}

type  Surface interface {
	getArea() float64;
}
