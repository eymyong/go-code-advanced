package struct1

import (
	"fmt"
	"time"
)

type Person struct {
	Name      string
	Sex       Sex
	BirthYear int
	Children  []*Person
}

type Sex int

// ค่าคงที่
const (
	// ไอโอต้า
	Male Sex = iota + 1
	Female
	LGTV
)

const (
	_Male   Sex = 1
	_Female Sex = 2
	_LGTV   Sex = 3
)

func (s Sex) String() string {
	if s == Male {
		return "Male"
	}
	if s == Female {
		return "Female"
	}

	if s == LGTV {
		return "LGTV+"
	}

	return "ไม่รู้เพศ"

}

func (a Person) Age() int {

	return time.Now().Year() - a.BirthYear
}

func (i Person) IsAdult() bool {
	if i.Age() >= 20 {
		return true
	}
	return false
}

func (konyed *Person) Yed(tokyed *Person, s string) *Person {
	//baby :=[]Person{}

	if konyed.Sex != tokyed.Sex {
		baby := Person{
			Name:      s,
			Sex:       konyed.Sex,
			BirthYear: time.Now().Year(),
		}
		fmt.Println("child of", konyed.Name, "yed", tokyed.Name, "::", baby)

		konyed.Children = append(konyed.Children, &baby)
		tokyed.Children = append(tokyed.Children, &baby)
		return &baby
	}

	return nil
}

func printNotNil(p *Person) {
	if p != nil {
		fmt.Println("printNotNil", p)
	}
}

func checkSex(expect Sex, actual *Person) bool {
	if actual != nil {

		if expect == actual.Sex {
			return true
		}
	}

	return false
}

func Struct() {

	yong := Person{
		Name:      "yong",
		Sex:       Male,
		BirthYear: 1997,
		Children:  nil,
	}
	noon := Person{
		Name:      "noon",
		Sex:       Female,
		BirthYear: 1997,
		Children:  nil,
	}

	pak := Person{
		Name:      "pak",
		Sex:       Male,
		BirthYear: 1997,
		Children:  nil,
	}

	summer := yong.Yed(&noon, "summer")
	black := noon.Yed(&yong, "black")

	// Adopt (รับเลี้ยง) เป็น function ที่รับเลี้ยงเด็กโดยต้องส่งคนเข้าไป
	pak.Adopt(summer)
	pak.Adopt(black)

	fmt.Println("pak:", pak)
	fmt.Println(*pak.Children[0]) // &{summer Male 2024 nil} (ต้องมี summer เป็น children ของ pak)
	fmt.Println("====")

	// สร้าง function mapOddEven รับ array ที่ return map[int]string
	// หาก key เป็นเลขคู่ value จะเป็น even หาก key เป็น เลขคี่ value จะเป็น odd
	// ใช้ for loop ได้
	numbers := []int{1, 90, 75}
	evenNumbers := mapOddEven(numbers)
	fmt.Println(evenNumbers[1])  // odd
	fmt.Println(evenNumbers[90]) // even
	fmt.Println(evenNumbers[75]) // odd
}

func mapOddEven(arr []int) map[int]string {
	m := make(map[int]string)
	for i := range arr {
		v := arr[i]
		if v%2 != 0 {
			m[v] = "odd"
		} else {
			m[v] = "even"
		}

		// if v%2 != 0 {
		// 	m[v] = "odd"
		// 	continue
		// }

		// m[v] = "even"
	}
	return m
}

// Adopt (รับเลี้ยง)
func (p *Person) Adopt(kid *Person) {
	p.Children = append(p.Children, kid)
	fmt.Println(p.Name, "รับ", kid.Name, "มาเลี้ยง")
}

// func main() {

// 	numbers := []int{1,3,4,3,3,2,1}

// 	uniqueNumbers := unique(numbers)

// 	fmt.Println(uniqueNumbers) // [1,2,3,4] (ห้ามมีตัวซ้ำ)

//   }

//   func unique(arr []int) []int{
// 	return nil
//   }
