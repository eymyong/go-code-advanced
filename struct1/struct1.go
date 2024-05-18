package struct1

import (
	"fmt"
	"math"
)

type person struct {
	name string
	sex  string
	age  int
}

func Cherkcode() {

	var personList = []person{
		person{
			name: "yong",
			sex:  "male",
			age:  25,
		},
		person{
			name: "pak",
			sex:  "female",
			age:  20,
		},
		person{
			name: "art",
			sex:  "female",
			age:  35,
		},
	}

	age18(personList)
	fuck(personList[0], personList[1])

}

func age18(personList []person) []person {
	personNew := []person{}
	for i := range personList {
		v := personList[i]
		vv := personList[i].age
		if vv > 18 {
			personNew = append(personNew, v)
		}
	}
	return personNew
}

func fuck(p1 person, p2 person) {
	//ถ้า p1, p2 อายุห่างกันเกิน 4 ปี, ไม่เย็ด
	if p1.age < 18 || p2.age < 18 {
		fmt.Println("no fuck")
		return
	}

	if math.Abs(float64(p1.age-p2.age)) >= 4 {
		fmt.Println("no fuck")
		return
	}
	// ถ้า p1, p2 ญ คู่เลย ให้เย็ดกันโดยร้องว่า "Ah yes <ชื่อ p2> please fist my ass"
	if p1.sex == p2.sex {
		if p1.sex == "female" {
			fmt.Println("Ah yes", p2.name, "please fist my ass")
			return
		}
		// ถ้า p1, p2 ช คู่เลย ให้เย็ดกันโดยร้องว่า "<ชื่อ p2>, you gave the best blowjob"
		if p1.sex == "male" {
			fmt.Println("Ah yes", p2.name, "please fist my ass")
			return
		}

	}
	// ถ้า p1, p2 ต่างเพศ ให้เย็ดกันโดยร้องว่า "<ชื่อคนที่เป็น ผช> ahhh this is the best sex in <ชื่อ ผญ>'s life!"
	if p1.sex != p2.sex {
		fmt.Println(p1.name, " ahhh this is the best sex in ", p2.name, "'s life!")
		return
	}

}
