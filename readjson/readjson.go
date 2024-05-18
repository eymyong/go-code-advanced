package readjson

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name  string
	Sex   string
	Skill map[string]int
	Cars  []Car
}

type Car struct {
	Brand string
	HP    int
}

func Readjson() {

	yong := Person{
		Name: "yong",
		Sex:  "Male",
	}
	_ = yong

	art := Person{
		Name: "art",
		Sex:  "Male",
		Skill: map[string]int{
			"football": 3,
		},
	}

	persons := []Person{yong, art}

	vios := Car{
		Brand: "Toyota",
		HP:    110,
	}

	pak := Person{
		Name: "pak",
		Sex:  "Male",
		Cars: []Car{vios},
	}
	_ = pak

	b, err := json.Marshal(persons)
	if err != nil {
		panic(err)
	}

	writeFile("pak.json", b)

	r1 := readFile("pak.json")
	//r2 := readFile("yong.json")
	fmt.Println(r1)

	// unmarshal(r1)
	// unmarshal(r2)
}
func writeFile(nameFile string, arrByte []byte) {
	err := os.WriteFile(nameFile, arrByte, 0664)
	if err != nil {
		panic(err)
	}
}

func readFile(fileName string) []byte {

	b, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return b
}

func unmarshal(b []byte) {

	//b := []byte(a)
	var p []Person

	err := json.Unmarshal(b, &p)
	if err != nil {
		panic(err)
	}

	fmt.Println(p)
}

// read file json
// Unmarshal > struct || ปริ้นให้ได้ทุก field

// personsByte, err := json.Marshal(persons)
// if err != nil {
// 	panic(err.Error())
// }

// err = os.WriteFile("persons/yong:)", personsByte, 0664)
// if err != nil {
// 	panic(err.Error())
// }
