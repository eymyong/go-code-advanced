package inf

import "fmt"

type shape interface {
	area() float64
	area32() float32
}

type circle struct {
	r float64
}

func (c circle) area() float64 {
	return 3.14 * c.r * c.r
}

func (c circle) area32() float32 {
	return float32(c.area())
}

type rectangle struct {
	w float64
	h float64
}

func (r rectangle) area() float64 {
	return r.w * r.h
}

func (r rectangle) area32() float32 {
	return float32(r.area())
}

type triangle struct {
	base float64
	h    float64
}

func (t *triangle) area() float64 {
	return 0.5 * t.base * t.h
}

func (t *triangle) area32() float32 {
	return float32(t.area())
}

func Foo() {
	c := circle{r: 10}
	r := rectangle{w: 10, h: 20}
	t := triangle{base: 10, h: 5}

	fmt.Println("c area", c.area())
	fmt.Println("r area", r.area())
	fmt.Println("t area", t.area())

	fmt.Println("c > r", isBiggerCool(&t, r))
	fmt.Println("r > c", isBiggerCool(r, c))
}

func isBiggerCool(l shape, r shape) bool {
	return l.area() > r.area()
}

func isBigger(c circle, r rectangle) bool {
	return c.area() > r.area()
}
