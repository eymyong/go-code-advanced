package methot

import "fmt"

type Person struct {
	name     string
	sex      string
	age      int
	children []Person
}

type circle struct {
	r float64
}

func MethotBasic() {
	pong := Person{name: "pong", age: 20}
	ped := Person{name: "ped", age: 28}

	c1 := circle{r: 4}
	c2 := circle{r: 6}

	haveSex(ped, pong)
	ped.haveSex(pong)

	fmt.Println(pong.isOlder(ped)) //false
	fmt.Println(ped.isOlder(pong)) // true

	fmt.Println(isOlder(ped, pong)) // false
	fmt.Println(isOlder(pong, ped)) // true
	fmt.Println(pong.isOlder(ped))  // false

	fmt.Println(c1.isBigger(c2))

}

func (c circle) area() float64 {
	return 3.14 * c.r * c.r
}

// c1.isBigger(c2)
func (c1 circle) isBigger(c2 circle) bool {
	return c1.area() > c2.area()
}

func greet(p1, p2 Person) {
	fmt.Println("hello", p2.name, "my name is", p1.name)
}

func (p Person) greet(other Person) {
	fmt.Println("hello", other.name, "my name is", p.name)
}

func haveSex(p1, p2 Person) {
	fmt.Println(p1.name, "is fucking", p2.name)
}

func (p Person) haveSex(other Person) {
	fmt.Println(p.name, "is fucking", other.name)
}

func isOlder(p1, p2 Person) bool {
	return p1.age > p2.age
}

func (p Person) isOlder(other Person) bool {
	return p.age > other.age
}

// ===================================

func Methot1() {
	art := Person{name: "art", age: 27}

	fmt.Printf("main: before: %+v\n", art)
	birthday(art)
	fmt.Printf("main: after: %+v\n", art)

	i := 10 // int
	addOne(&i)
	fmt.Println(i)
}

func birthday(p Person) {
	fmt.Printf("birthday: before: %+v\n", p)
	p.age++
	fmt.Printf("birthday: after: %+v\n", p)
}

func addOne(i *int) {
	fmt.Println("addOne: before", i)
	*i++
	fmt.Println("addOne: after", i)
}
