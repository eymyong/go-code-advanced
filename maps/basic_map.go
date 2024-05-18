package maps

import "fmt"

// ถ้ามี `map` ชื่อว่า mapInt := map[String]int  ==> หมายความว่า ตัวแปร `mapInt` เป็น map ซึ่งมี `key`เป็น `type string` และ `value` เป็น `type int`
// mapInt = ["one": 1,"two": 2,"three": 3]  ==> `mapInt` มีข้อมูลอยู่แล้ว
// mapString =[ 1:"one", 2:"two", 3:"three"]

// *******************************************************************************************//
// * การเขียนค่า * assing  หรือ เพิ่มข้อมูลลงใน map  ==> (ทำได้โดยการเขียนแบบนี้  mapInt["four"] = 4 , mapString[5] = "five")
//        ( mapInt["four"] = 4 )

// * การอ่านค่า * การเข้าถึง value หรือการเรียกค่าของ map ออกมาดู ==>  (ทำได้โดยการเขียนแบบนี้ mapInt["four"] ==> value ที่ออกมา = 4)
// * การอ่านค่าโดยสร้างตัวแปรขึ้นมารับ (เพื่อที่จำนำข้อมูลไปใช้ต่อ)*  ==> (valueM := mapInt["four"] ) หมายความว่า value ของM["four"] ถูกเก็บไว้ใน `valueM`
//        ( valueM4 := mapInt["four"] ) ซึ่งหมายความว่าจริงๆแล้ว ( mapInt["four"] และ valueM4 ก็คือการเข้าถึง value เหมือนกันแค่เขียนกันคนละแบบ และนำไปใช้ได้ต่างกัน )

//  Ex1)  สร้าง map ใหม่ ที่ ( key เป็น value ของ mapInt ) และ ( value มาจาก value ของ mapString )
//	 mapSum := map[int][]String{}

//  mapSum[mapInt["one"]] = mapString[1]   ==> เรียกผ่าน key = เรียกผ่าน key
//   mapSum[value1] = valueOne              ==> เรียกตัวแปร value โดยตรง = เรียกตัวแปร value โดยตรง
//   mapSum[mapInt["one"]] = valueONe       ==> เรียกผ่าน key = เรียกตัวแปร value โดยตรง
//	 mapSum[value1] = mapString[1]          ==> เรียกตัวแปร value โดยตรง =  เรียกผ่าน key

// * การใช้ _,ok *   ==> `ok` เป็น `type bool` ซึ่งเอาไว้เช็คว่า map ที่เรากำลังกระทำอยู่ มี key นั้นๆหรือไม่

//============
// func EX1(ex1 []int) map[int]int{} >> รับ[]int ตอน for loop ควรสร้าง for _,v := range ex1 {}

type student struct {
	id    string
	class int
}

//Add ของลงใน map
func AddMapString() map[int]string {
	mapString := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	// เพิ่มของเข้าไป 3 ตัว
	mapString[4] = "four"
	mapString[5] = "five"
	mapString[6] = "six"
	fmt.Println("mapString = ", mapString)

	// อ่านของออกจาก map ทั้งหมด
	// valueMapString1 := mapString[1]
	// valueMapString2 := mapString[2]
	// valueMapString3 := mapString[3]
	// valueMapString4 := mapString[4]
	// valueMapString5 := mapString[5]
	// valueMapString6 := mapString[6]
	// fmt.Println(valueMapString1)
	// fmt.Println(valueMapString2)
	// fmt.Println(valueMapString3)
	// fmt.Println(valueMapString4)
	// fmt.Println(valueMapString5)
	// fmt.Println(valueMapString6)

	return mapString
}

//Add ของลงใน map
func AddMapArrInt() map[string][]int {
	mapArrInt := map[string][]int{
		"A": {1, 2, 3},
		"B": {4, 5, 6},
		"C": {7, 8, 9},
	}
	_ = mapArrInt
	// เพิ่มของเข้าไป 3 ตัว
	mapArrInt["D"] = []int{2, 4, 6}
	mapArrInt["E"] = []int{1, 3, 5}
	mapArrInt["F"] = []int{3, 5, 7}

	// อ่านของออกจาก map ทั้งหมด
	// valueInKeyA := mapArrInt["A"]
	// valueInKeyB := mapArrInt["B"]
	// valueInKeyC := mapArrInt["C"]
	// valueInKeyD := mapArrInt["D"]
	// valueInKeyE := mapArrInt["E"]
	// valueInKeyF := mapArrInt["F"]
	// fmt.Println(valueInKeyA)
	// fmt.Println(valueInKeyB)
	// fmt.Println(valueInKeyC)
	// fmt.Println(valueInKeyD)
	// fmt.Println(valueInKeyE)
	// fmt.Println(valueInKeyF)

	return mapArrInt
}
func Foo1() {
	dataMapString := AddMapString()
	dataMapArrInt := AddMapArrInt()
	_ = dataMapString
	_ = dataMapArrInt

}

// //สร้าง map ใหม่ ที่ key เป็น value ของ mapString และ value มาจาก value ของ mapArrInt
// func AddMapStringAndInt(dataMapString map[int]string, dataMapInt map[string][]int) map[string][]int {
// 	sumMap := make(map[string][]int)
// 	// var valurString string
// 	// var valueArrInt []int
// 	for i := range dataMapString {
// 		valurString := dataMapString[i]
// 		sumMap[valurString] = []int{}
// 	}
// 	fmt.Println("====", sumMap)

// 	for i := range dataMapInt {
// 		//fmt.Println(v)
// 		value := dataMapInt[i]
// 		sumMap[i] = value
// 		fmt.Println(value)
// 	}
// 	fmt.Println("====", sumMap)

// 	return sumMap

// }
func AddmapReadmap() {

	mapString := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}

	// เพิ่มของเข้าไป 3 ตัว
	mapString[4] = "four"
	mapString[5] = "five"
	mapString[6] = "six"
	fmt.Println("mapString = ", mapString)

	// อ่านของออกจาก map ทั้งหมด
	valueMapString1 := mapString[1]
	valueMapString2 := mapString[2]
	valueMapString3 := mapString[3]
	valueMapString4 := mapString[4]
	valueMapString5 := mapString[5]
	valueMapString6 := mapString[6]

	mapArrInt := map[string][]int{
		"A": {1, 2, 3},
		"B": {4, 5, 6},
		"C": {7, 8, 9},
	}
	_ = mapArrInt
	// เพิ่มของเข้าไป 3 ตัว
	mapArrInt["D"] = []int{2, 4, 6}
	mapArrInt["E"] = []int{1, 3, 5}
	mapArrInt["F"] = []int{3, 5, 7}
	fmt.Println("mapArrInt = ", mapArrInt)

	// อ่านของออกจาก map ทั้งหมด

	valueMATa := mapArrInt["A"]
	valueMAIb := mapArrInt["B"]
	valueMAIc := mapArrInt["C"]
	valueMAId := mapArrInt["D"]
	valueMAIe := mapArrInt["E"]
	valueMAIf := mapArrInt["F"]

	//สร้าง map ใหม่ ที่ key เป็น value ของ mapString และ value มาจาก value ของ mapArrInt
	mapSum := map[string][]int{}
	_ = mapSum

	mapSum[mapString[1]] = mapArrInt["A"]
	mapSum[valueMapString1] = valueMATa

	mapSum[valueMapString2] = mapArrInt["B"]
	mapSum[mapString[2]] = valueMAIb

	mapSum[valueMapString3] = valueMAIc
	mapSum[mapString[3]] = mapArrInt["C"]

	mapSum[valueMapString4] = mapArrInt["D"]
	mapSum[mapString[4]] = valueMAId

	mapSum[mapString[5]] = mapArrInt["E"]
	mapSum[valueMapString5] = valueMAIe

	mapSum[mapString[6]] = valueMAIf
	mapSum[valueMapString6] = mapArrInt["F"]

	m := make(map[string]student)
	_ = m
	// เพิ่มของเข้าไป 3 ตัว
	m["yong"] = student{id: "YY", class: 1}
	m["pak"] = student{id: "PP", class: 2}
	m["art"] = student{id: "AA", class: 3}

	fmt.Println("map m = ", m)

	// อ่านของออกจาก map ทั้งหมด
	valueMYong := m["yong"]
	valueMPak := m["pak"]
	valueMArt := m["art"]

	fmt.Println("valueMYong =", valueMYong)
	fmt.Println("valueMYong = ", valueMPak)
	fmt.Println("valueMYong = ", valueMArt.class)

	ms := make(map[string][]student)
	_ = ms
	// เพิ่มของเข้าไป 3 ตัว
	ms["maleList"] = []student{{id: "mond", class: 1}, {id: "pong", class: 2}, {id: "puen", class: 3}}
	ms["femaleList"] = []student{{id: "cee", class: 2}, {id: "noon", class: 1}, {id: "pla", class: 3}}
	ms["LGBTQList"] = []student{{id: "aa", class: 2}, {id: "bb", class: 1}, {id: "cc", class: 3}}

	// อ่านของออกจาก map ทั้งหมด
	valueMalelist := ms["maleList"]
	valueFemalelist := ms["femaleList"]
	valueLGBTQList := ms["LGBTQList"]

	fmt.Println("valueMalelist", valueMalelist)
	fmt.Println("valueFemalelist", valueFemalelist)
	fmt.Println("valueLGBTQList", valueLGBTQList)

	mapSumM := make(map[student][]student) // สร้าง mapSumM โดยที่ key ของ mapSumM (มาจาก key m)   ||value ของ mapSumM (จะต้องมาจาก m)
	_ = mapSumM

}

//======================================================================================================

var dictEng = map[int]string{
	2: "two",
	1: "one",
	3: "three",
}

var dictThai = map[int]string{
	3: "สาม",
	1: "หนึ่ง",
	2: "สอง",
}

var dictEngGer = map[string]string{
	"one":   "eine",
	"two":   "zwei",
	"three": "drei",
}

var arr = []int{1, 2, 3}

func dict(arr []int) {
	for i := range arr {
		valueInt := arr[i]

		en := dictEng[valueInt]
		th := dictThai[valueInt]
		ger := dictEngGer[en]

		fmt.Println("en", en, "th", th, "ger", ger)
	}
}

//=========================================================================================

// type fruitColors map[string]string

// func main() {

// 	f := fruitColors{
// 		"mango":  "yellow",
// 		"orange": "orange",
// 		"banana": "yellow",
// 		"tangmo": "red",
// 	}
// 	fmt.Println(f)
// 	fmt.Println(f.getColor("banana"))

// }

// func (f fruitColors) addNewFruit(fruit string, color string) { // add fruit
// 	f[fruit] = color
// }

// func (f fruitColors) getColor(fruitName string) string { // cherk color โดยรับ fruit ที่ใส่เข้าไป
// 	color := f[fruitName]
// 	return color
// }

// func (f fruitColors) groupNumberFruits() map[string]int { //แยกสี โดยแต่ละสีต้องมีผลไม้ที่เป็นสีนั้นๆ
// 	resule := map[string]int{}
// 	for fruit := range f {
// 		color := f[fruit]
// 		resule[color]++
// 	}
// 	return resule
// }

// func (f fruitColors) groupFruits() map[string][]string { //แยกสี โดยแต่ละสีต้องมีผลไม้ที่เป็นสีนั้นๆ
// 	result := map[string][]string{}
// 	for fruit := range f {
// 		color := f[fruit]
// 		result[color] = append(result[color], fruit)
// 	}
// 	return result
// }

// func deleteFruit() {

// }
