package readjson

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person2 struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name"`
	Age  int    //`json:"age,omitempty"`
	Sex  sex
}

func read(fileName string) []byte { // read file
	b, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return b
}

func unmarshal1(fileName []byte) []Person2 { // แปลงfileที่เป็น []byte ของjson ออกมาเป็น struct
	data := []Person2{}
	err := json.Unmarshal(fileName, &data)
	if err != nil {
		panic(err)
	}
	return data
}

func write(nameFile string, arrByte []byte) {
	err := os.WriteFile(nameFile, arrByte, 0664)
	if err != nil {
		panic(err)
	}
}

func deletePerson(persons []Person2, name string) []Person2 {
	result := []Person2{}
	for i := range persons {
		if persons[i].Name != name {
			result = append(result, persons[i])
		}
	}
	return result
}

func AddDeletFile() {

	inputBytes := read("input.json")       // data = []byte
	inputPersons := unmarshal1(inputBytes) // data = []Person2

	addBytes := read("add.json")       // data = []byte
	addPersons := unmarshal1(addBytes) // data = []Person2

	result := append(inputPersons, addPersons...)

	//result = deletePerson(result, "yong")

	fmt.Println(result)

	b, err := json.Marshal(result) // sent json
	if err != nil {
		panic(err)
	}

	write("out.json", b)

}

//=================================

func Homework1() {

	inputByte := read("input.json")
	inputPersons := unmarshal1(inputByte)

	addByte := read("add.json")
	addPersons := unmarshal1(addByte)

	for i := range addPersons {
		v := addPersons[i].Age
		v += 10
		addPersons[i].Age = v
	}

	result := append(inputPersons, addPersons...)
	_ = result
	//fmt.Println(result)

	b, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	write("output.json", b)
}

func Homework2() {

	input2Byte := read("input2.json")
	inputPersons := unmarshal1(input2Byte)
	_ = inputPersons

	malePersons := []Person2{}
	femalePersons := []Person2{}

	for _, v := range inputPersons {
		if v.Sex == "male" {
			malePersons = append(malePersons, v)
			continue
		}
		femalePersons = append(femalePersons, v)
	}

	fmt.Println(malePersons)
	fmt.Println(femalePersons)

	b, err := json.Marshal(malePersons)
	if err != nil {
		panic(err)
	}
	write("output-male.json", b)

	b, err = json.Marshal(femalePersons)
	if err != nil {
		panic(err)
	}
	write("output-female.json", b)

	r1 := read("output-male.json")
	r2 := read("output-female.json")
	unmarshal1(r1)
	unmarshal1(r2)
	fmt.Println(unmarshal1(r1))
	fmt.Println(unmarshal1(r2))

}

type sex string

const male sex = "male"
const female sex = "female"

type bySex2 map[sex][]Person2

func Homework3() {
	inputByte := read("input3.json")
	inputPersons := unmarshal1(inputByte)
	fmt.Println(inputPersons)

	//output := make(bySex2)
	output1 := map[string][]Person2{}
	for _, v := range inputPersons {
		key := string(v.Sex)
		output1[key] = append(output1[key], v)
	}

	b, err := json.Marshal(output1)
	if err != nil {
		panic(err)
	}

	write("output3.json", b)

}

func tryBySex2() {
	bs := bySex2{
		male: []Person2{
			{Name: "art", Sex: male},
			{Name: "yong", Sex: female},
		},
		female: []Person2{
			{Name: "noon", Sex: female},
		},
	}

	j, _ := json.Marshal(bs)
	fmt.Printf("%s\n", j)
}
